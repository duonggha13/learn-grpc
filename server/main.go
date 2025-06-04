package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	proto "learn-grpc/grpc"
	"log"
	"net"
)

type server struct {
	proto.UnimplementedOrderServiceServer
}

func (s *server) NewOrder(ctx context.Context, in *proto.NewRequestOrder) (*proto.NewResponseOrder, error) {
	log.Printf("Receive order: %v", in.GetOrderRequest())
	return &proto.NewResponseOrder{OrderResponse: "new id order" + in.GetOrderRequest()}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":9000"))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	proto.RegisterOrderServiceServer(s, &server{})
	log.Printf("server listening on port %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
