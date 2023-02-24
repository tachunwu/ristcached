package main

import (
	"context"
	"net/http"
	"sync"

	"github.com/bufbuild/connect-go"
	ristcachedv1 "github.com/tachunwu/ristcached/pkg/proto/ristcached/v1"
	"github.com/tachunwu/ristcached/pkg/proto/ristcached/v1/ristcachedv1connect"
)

var wg sync.WaitGroup

func main() {
	client := ristcachedv1connect.NewRistcachedServiceClient(
		http.DefaultClient,
		"http://localhost:11212",
	)

	r := &ristcachedv1.SetRequest{
		Item: &ristcachedv1.KeyValue{
			Key:   "key",
			Value: []byte("value"),
			Cost:  1,
			Ttl:   0,
		},
	}
	client.Set(
		context.Background(),
		connect.NewRequest(r),
	)

}

func ClientPool(n int) []*ristcachedv1connect.RistcachedServiceClient {
	pool := []*ristcachedv1connect.RistcachedServiceClient{}
	for i := 0; i < n; i++ {
		c := ristcachedv1connect.NewRistcachedServiceClient(
			http.DefaultClient,
			"http://localhost:11212",
		)
		pool = append(pool, &c)
	}
	return pool
}

func Benchmark(pool []*ristcachedv1connect.RistcachedServiceClient) {

}
