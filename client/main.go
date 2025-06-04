package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	proto "learn-grpc/grpc"
	"log"
	"time"
)

func main() {
	address := "localhost:9000"
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := proto.NewOrderServiceClient(conn)

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		orderID := "1001"
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		r, err := c.NewOrder(ctx, &proto.NewRequestOrder{OrderRequest: orderID})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Order: %s", r.GetOrderResponse())
		cancel()
	}
}
