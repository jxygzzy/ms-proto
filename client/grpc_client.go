package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"ms-proto/service"
)

func main() {
	conn, err := grpc.Dial(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("连接失败")
	}
	defer conn.Close()
	client := service.NewProdServiceClient(conn)
	requset := &service.ProductRequest{
		ProdId: 123,
	}
	resp, err := client.GetProductStock(context.Background(), requset)
	if err != nil {
		log.Fatal("查询库存出错")
	}
	fmt.Println("查询成功:", resp.ProdStock)
}
