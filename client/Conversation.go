package main

import (
	"context"
	"log"

	address "grpc-test/v1"
)

// conversations 调用服务端的Conversations方法
func conversations() {
	// 调用服务端的Conversations方法，获取流
	payload := &address.Payload{
		Metadata: &address.Metadata{
			RequestType: address.RequestType_ROUTE,
			Key:         "apptest",
		},
	}
	stream, err := streamClient.GetStrategy(context.Background(), payload)
	if err != nil {
		log.Fatalf("get conversations stream err: %v", err)
	}
	res := stream.GetBody()
	// 打印返回值
	log.Println(res.String())
}
