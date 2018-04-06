package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/context"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/nat-ventura/rpc-party/proto"
	"google.golang.org/grpc"

	"google.golang.org/grpc/grpclog"
)

const port = 10000

var (
	tlsCertFilePath = flag.String("tls_cert_file", "./../ts/localhost.crt", "Path to the CRT/PEM file.")
	tlsKeyFilePath  = flag.String("tls_key_file", "./../ts/localhost.key", "Path to the private key file.")
)

type chatServer struct {
	streams []proto.Chat_ConnectServer
}

func (s *chatServer) Connect(stream proto.Chat_ConnectServer) error {
	fmt.Println("message received")
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

func (s *chatServer) GetMessage(ctx context.Context, msg *proto.GetMessageRequest) (*proto.WebMessage, error) {
	fmt.Printf("here\n")
	return &proto.WebMessage{}, nil
}

func newServer() *chatServer {
	s := new(chatServer)
	return s
}

func main() {
	fmt.Println("starting")
	// lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	// if err != nil {
	// grpclog.Fatalf("failed to listen: %v", err)
	// }
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	proto.RegisterChatServer(grpcServer, &chatServer{})
	wrappedServer := grpcweb.WrapServer(grpcServer)

	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}

	if err := httpServer.ListenAndServeTLS(*tlsCertFilePath, *tlsKeyFilePath); err != nil {
		grpclog.Fatalf("failed starting http2 server: %v", err)
	}

	// if err := httpServer.ListenAndServe(); err != nil {
	// 	grpclog.Fatalf("failed starting http server: %v", err)
	// }

	// grpcServer.Serve(lis)
}
