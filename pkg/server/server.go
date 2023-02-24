package server

import (
	"context"
	"net/http"
	"time"

	"github.com/dgraph-io/ristretto"
	ristcachedv1 "github.com/tachunwu/ristcached/pkg/proto/ristcached/v1"
	"github.com/tachunwu/ristcached/pkg/proto/ristcached/v1/ristcachedv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/bufbuild/connect-go"
)

type RistcachedServer struct {
	cache *ristretto.Cache
}

func NewRistcachedServer() *RistcachedServer {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     2 << 30, // maximum cost of cache (2GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}

	return &RistcachedServer{
		cache: cache,
	}
}

func (s *RistcachedServer) Start() {
	mux := http.NewServeMux()
	path, handler := ristcachedv1connect.NewRistcachedServiceHandler(s)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:11212",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

func (s *RistcachedServer) Get(ctx context.Context, req *connect.Request[ristcachedv1.GetRequest]) (*connect.Response[ristcachedv1.GetResponse], error) {

	value, found := s.cache.Get(req.Msg.GetKey())
	if !found {
		res := connect.NewResponse(&ristcachedv1.GetResponse{
			Found: false,
		})
		return res, nil
	}
	res := connect.NewResponse(&ristcachedv1.GetResponse{
		Found: true,
		Value: value.([]byte),
	})

	return res, nil
}

func (s *RistcachedServer) Set(ctx context.Context, req *connect.Request[ristcachedv1.SetRequest]) (*connect.Response[ristcachedv1.SetResponse], error) {

	ok := s.cache.Set(
		req.Msg.GetItem().GetKey(),
		req.Msg.GetItem().GetValue(),
		req.Msg.GetItem().GetCost(),
	)
	if !ok {
		res := connect.NewResponse(&ristcachedv1.SetResponse{
			Added: false,
		})
		return res, nil
	}

	res := connect.NewResponse(&ristcachedv1.SetResponse{
		Added: true,
	})

	return res, nil
}

func (s *RistcachedServer) SetWithTTL(ctx context.Context, req *connect.Request[ristcachedv1.SetWithTTLRequest]) (*connect.Response[ristcachedv1.SetWithTTLResponse], error) {
	ok := s.cache.SetWithTTL(
		req.Msg.GetItem().GetKey(),
		req.Msg.GetItem().GetValue(),
		req.Msg.GetItem().GetCost(),
		time.Duration(req.Msg.GetItem().GetTtl()),
	)
	if !ok {
		res := connect.NewResponse(&ristcachedv1.SetWithTTLResponse{
			Added: false,
		})
		return res, nil
	}

	res := connect.NewResponse(&ristcachedv1.SetWithTTLResponse{
		Added: true,
	})

	return res, nil
}

func (s *RistcachedServer) Del(ctx context.Context, req *connect.Request[ristcachedv1.DelRequest]) (*connect.Response[ristcachedv1.DelResponse], error) {
	s.cache.Del(req.Msg.GetKey())

	res := connect.NewResponse(&ristcachedv1.DelResponse{})

	return res, nil
}

func (s *RistcachedServer) GetTTL(ctx context.Context, req *connect.Request[ristcachedv1.GetTTLRequest]) (*connect.Response[ristcachedv1.GetTTLResponse], error) {
	ttl, found := s.cache.GetTTL(req.Msg.GetKey())
	if !found {
		res := connect.NewResponse(&ristcachedv1.GetTTLResponse{
			Found: false,
		})
		return res, nil
	}
	res := connect.NewResponse(&ristcachedv1.GetTTLResponse{
		Found: true,
		Ttl:   int64(ttl),
	})

	return res, nil
}

func (s *RistcachedServer) MaxCost(ctx context.Context, req *connect.Request[ristcachedv1.MaxCostRequest]) (*connect.Response[ristcachedv1.MaxCostResponse], error) {
	res := connect.NewResponse(&ristcachedv1.MaxCostResponse{
		MaxCost: s.cache.MaxCost(),
	})

	return res, nil
}

func (s *RistcachedServer) UpdateMaxCost(ctx context.Context, req *connect.Request[ristcachedv1.UpdateMaxCostRequest]) (*connect.Response[ristcachedv1.UpdateMaxCostResponse], error) {
	s.cache.UpdateMaxCost(req.Msg.GetMaxCost())
	res := connect.NewResponse(&ristcachedv1.UpdateMaxCostResponse{})
	return res, nil
}

func (s *RistcachedServer) Clear(ctx context.Context, req *connect.Request[ristcachedv1.ClearRequest]) (*connect.Response[ristcachedv1.ClearResponse], error) {
	s.cache.Clear()
	res := connect.NewResponse(&ristcachedv1.ClearResponse{})
	return res, nil
}
