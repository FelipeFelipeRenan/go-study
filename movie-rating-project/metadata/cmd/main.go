package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"movieexample/gen"
	"movieexample/metadata/internal/controller/metadata"
	grpchandler "movieexample/metadata/internal/handler/grpc"
	"movieexample/metadata/internal/repository/memory"
	"movieexample/pkg/discovery"
	"movieexample/pkg/discovery/consul"
	"net"
	"time"

	"google.golang.org/grpc"
)

const serviceName = "metadata"
func main() {
	var port int
	flag.IntVar(&port, "port", 8081, "API handler port")
	flag.Parse()
	log.Printf("Starting the metadata service on port %d", port)
	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil{
		panic(err)
	}

	go func() {
		for {
			if err := registry.ReportHealthyState(instanceID, serviceName); err != nil{
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer registry.Deregister(ctx, instanceID, serviceName)

	repo := memory.New()
	svc := metadata.New(repo)
	h := grpchandler.New(svc)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	gen.RegisterMetadataServiceServer(srv, h)
	if err := srv.Serve(lis); err != nil{
		panic(err)
	}
}