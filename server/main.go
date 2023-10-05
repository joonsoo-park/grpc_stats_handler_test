package main

import (
	"context"
	pb "dice/proto"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/protobuf/proto"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type Server struct {
	pb.UnimplementedDiceServer
}

func (s *Server) Greet(_ context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {

	// log.Printf("Received: %v\n", in.GetName())
	log.Println("Received bytes:", proto.Size(in))

	return &pb.GreetResponse{Greeting: "Hello "}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	handler := MyHandler{}
	opts := []grpc.ServerOption{
		grpc.StatsHandler(&handler),
	}
	s := grpc.NewServer(opts...)
	pb.RegisterDiceServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
