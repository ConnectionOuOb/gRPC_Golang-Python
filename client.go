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
	// 連接到服務端
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("連接失敗: %v", err)
	}
	defer conn.Close()
	client := pb.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 建立請求：傳入 []string 輸入
	req := &pb.CalcRequest{
		Inputs: []string{"hello", "world"},
	}

	// 呼叫 Calculate 方法
	resp, err := client.Calculate(ctx, req)
	if err != nil {
		log.Fatalf("計算失敗: %v", err)
	}

	// 印出返回的三層嵌套 float32 陣列
	// resp.Outers 為 []Outer，每個 Outer 有 Inners []Inner，每個 Inner 裡的 Values 為 []float32
	for i, outer := range resp.Outers {
		fmt.Printf("Outer %d:\n", i)
		for j, inner := range outer.Inners {
			fmt.Printf("  Inner %d: %v\n", j, inner.Values)
		}
	}
}
