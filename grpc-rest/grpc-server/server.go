package main

import (
	"context"
	"log"
	"net"

	pb "github.com/learning-go/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
}

func (s *server) Add(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	a, b := in.GetA(), in.GetB()

	result := a + b

	return &pb.Response{Result: result}, nil
}
func (s *server) Multiply(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	a, b := in.GetA(), in.GetB()

	result := a * b

	return &pb.Response{Result: result}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	log.Println("🚀 gRPC server listening on localhost:4040")
	defer listener.Close()

	srv := grpc.NewServer()
	pb.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}
