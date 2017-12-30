package main

import (
	"flag"
	"fmt"
	"log"

	pb "github.com/im-auld/users/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:8081", "The server address in the format of host:port")
	userID     = flag.Int64("userID", 123, "The user ID")
)

type UserServiceStub struct {
	client pb.UserServiceClient
}

func NewStub(addr *string, opts []grpc.DialOption) UserServiceStub {
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)
	return UserServiceStub{client: client}
}

func GetUser(client pb.UserServiceClient, userID int64) {
	request := &pb.GetUserRequest{UserId: userID}
	response, _ := client.GetUser(context.Background(), request)
	fmt.Println(response)
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	stub = NewStub(*serverAddr, opts)
	stub.GetUser(client, *userID)
}
