package main

import (
	"fmt"
	"io"
	"net"

	"github.com/nat-ventura/rpc-attack/proto"

	"google.golang.org/grpc"

	"google.golang.org/grpc/grpclog"
)

const port = 10000

type chatServer struct {
	streams []proto.Chat_ConnectServer
}

func (s *chatServer) Connect(stream proto.Chat_ConnectServer) error {
	s.streams = append(s.streams, stream)
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Println("%#v \n", msg)

		go func() {
			for i, stream_client := range s.streams {
				if err := stream_client.Send(msg); err != nil {
					s.streams = append(s.streams[:i], s.streams[i+1:]...)
				}
			}
		}()
	}
}

func newServer() *chatServer {
	s := new(chatServer)
	return s
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	proto.RegisterChatServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
