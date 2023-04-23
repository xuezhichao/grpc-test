package main

import (
	"context"
	"io"
	"log"
)

// conversations 调用服务端的Conversations方法
func conversations() {
	//调用服务端的Conversations方法，获取流
	stream, err := streamClient.Conversations(context.Background())
	if err != nil {
		log.Fatalf("get conversations stream err: %v", err)
	}
	for n := 0; n < 5; n++ {
		//err := stream.Send(&pb.StreamRequest{Question: "stream client rpc " + strconv.Itoa(n)})
		//if err != nil {
		//	log.Fatalf("stream request err: %v", err)
		//}
		//time.Sleep(time.Second * 15)

		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Conversations get stream err: %v", err)
		}
		// 打印返回值
		log.Println(res.Answer)
	}
	//最后关闭流
	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("Conversations close stream err: %v", err)
	}
}
