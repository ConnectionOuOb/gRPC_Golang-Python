package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "ConnectgRPC/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("連接失敗: %v", err)
	}
	defer conn.Close()
	client := pb.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &pb.CalcRequest{
		Inputs: []string{"hello", "world"},
	}

	resp, err := client.Calculate(ctx, req)
	if err != nil {
		log.Fatalf("計算失敗: %v", err)
	}

	for i, outer := range resp.Outers {
		fmt.Printf("Outer %d:\n", i)
		for j, inner := range outer.Inners {
			fmt.Printf("  Inner %d: %v\n", j, inner.Values)
		}
	}
}
