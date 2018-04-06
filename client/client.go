package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"math/rand"
	"strconv"

	"github.com/nat-ventura/rpc-attack/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const serverAddr = "localhost:10000"

var (
	username = flag.String("username", "unknown"+strconv.Itoa(rand.Int()), "Your username is displayed to other users in chat")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := proto.NewChatClient(conn)

	stream, err := client.Connect(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	}

	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf(err.Error())
			}

			fmt.Println(msg.Text)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := stream.Send(&proto.WebMessage{*username + ": " + scanner.Text()}); err != nil {
			log.Fatalf(err.Error())
		}
	}
}
