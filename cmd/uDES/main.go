package main

import (
	"log"
	"sync"
	"uDES/internal/grpc"
	"uDES/internal/http"
)

func main() {
	var wg sync.WaitGroup

	newGrpcserver := grpcserver.NewUserManagementServer()
	newGrpcserver.WorkerPool.Start()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := newGrpcserver.Run(); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	newHttpserver := httpserver.NewHTTPServer()
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := newHttpserver.Start(); err != nil {
			log.Fatalf("failed to serve HTTP: %v", err)
		}
	}()

	wg.Wait()
}
