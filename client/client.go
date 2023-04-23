package main

import (
	"google.golang.org/grpc"
	address "grpc-test/proto"
	"log"
)

// Address 连接地址
const Address string = ":8000"

var streamClient address.StreamClient

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	streamClient = address.NewStreamClient(conn)
	conversations()
}
