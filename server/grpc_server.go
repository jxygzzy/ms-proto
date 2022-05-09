package main

import (
	"fmt"
	"log"
	"ms-proto/service"
	"net"
	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	service.RegisterProdServiceServer(rpcServer, service.ProductService)
	listen, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal("启动监听出错")
	}
	err = rpcServer.Serve(listen)
	if err != nil {
		log.Fatal("启动服务出错")
	}
	fmt.Println("微服务启动成功")
	
}
