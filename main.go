package main

import (
    "context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    v1 "ivanrylach.com/hello-grpc/gen/go/ivanrylach/hello_grpc/v1"
    "log"
    "net"
)

const (
    port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
    v1.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloResponse, error) {
    log.Printf("Retrieve: %v", in.GetName())
    return &v1.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func (s *server) CreateHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloResponse, error) {
    log.Printf("Create: %v", in.GetName())
    return &v1.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    reflection.Register(s)
    v1.RegisterHelloServiceServer(s, &server{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
