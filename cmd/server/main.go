package main

import (
	"context"
	"net/http"

	ristcachedv1 "github.com/tachunwu/ristcached/pkg/proto/ristcached/ristcachedv1"
	"github.com/tachunwu/ristcached/pkg/proto/ristcached/ristcachedv1/ristcachedv1connect"
	"github.com/tachunwu/ristcached/pkg/proto/ristcached/v1/ristcachedv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type RistcachedServer struct{}

func main() {
	ristcached := &RistcachedServer{}
	mux := http.NewServeMux()
	path, handler := ristcachedv1connect.NewRistcachedServiceHandler(ristcached)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:11212",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

func (s *RistcachedServer) Get(ctx context.Context, req *connect.Request[ristcachedv1.GetRequest]) (*connect.Response[ristcachedv1.GetResponse], error) {
	return nil, nil
}

func (s *RistcachedServer) Set(ctx context.Context, req *connect.Request[ristcachedv1.SetRequest]) (*connect.Response[ristcachedv1.SetResponse], error) {
	return nil, nil
}
func (s *RistcachedServer) SetWithTTL(ctx context.Context, req *connect.Request[ristcachedv1.SetWithTTLRequest]) (*connect.Response[ristcachedv1.SetWithTTLResponse], error) {
	return nil, nil
}
func (s *RistcachedServer) Del(ctx context.Context, req *connect.Request[ristcachedv1.DelRequest]) (*connect.Response[ristcachedv1.DelResponse], error) {
	return nil, nil
}
func (s *RistcachedServer) GetTTL(ctx context.Context, req *connect.Request[ristcachedv1.GetTTLRequest]) (*connect.Response[ristcachedv1.GetTTLResponse], error) {
	return nil, nil
}
func (s *RistcachedServer) MaxCost(ctx context.Context, req *connect.Request[ristcachedv1.MaxCostRequest]) (*connect.Response[ristcachedv1.MaxCostResponse], error) {
	return nil, nil
}
func (s *RistcachedServer) UpdateMaxCost(ctx context.Context, req *connect.Request[ristcachedv1.UpdateMaxCostRequest]) (*connect.Response[ristcachedv1.UpdateMaxCostResponse], error) {
	return nil, nil
}
func (s *RistcachedServer) Clear(ctx context.Context, req *connect.Request[ristcachedv1.ClearRequest]) (*connect.Response[ristcachedv1.ClearResponse], error) {
	return nil, nil
}
