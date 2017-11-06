package server

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/im-auld/users/proto"
)

var (
	port = flag.Int("port", 8081, "The server port")
)

func newServer() *UserService {
	s := &UserService{DB: *NewDBClient()}
	return s
}

func StartServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterUserServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
