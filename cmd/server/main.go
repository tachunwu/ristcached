package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/tachunwu/ristcached/pkg/server"
)

func main() {

	ristcached := server.NewRistcachedServer()
	ristcached.Start()

	// Shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	os.Exit(0)
}
