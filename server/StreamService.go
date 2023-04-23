package main

import (
	pb "grpc-test/proto"
	"io"
	"log"
	"strconv"
)

type StreamService struct {
	*pb.UnimplementedStreamServer
}

func (s *StreamService) mustEmbedUnimplementedStreamServer() {
}

// Conversations 实现Conversations方法
func (s *StreamService) Conversations(srv pb.Stream_ConversationsServer) error {
	n := 1
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = srv.Send(&pb.StreamResponse{
			Answer: "from stream server answer: the " + strconv.Itoa(n) + " question is " + req.Question,
		})
		if err != nil {
			return err
		}
		n++
		log.Printf("from stream client question: %s", req.Question)
	}
}
