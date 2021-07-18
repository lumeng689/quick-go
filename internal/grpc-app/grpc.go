package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "quick-go/api/protobuf/helloworld"
	pbs "quick-go/internal/grpc-app/server"
)

var (
	grpcPort = ":50051"
)

func main() {
	fmt.Println("hello protobuf...")
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		fmt.Printf("Failed to listen:%v", grpcPort)
	}

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &pbs.GreeterServer{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("create grpc server failed! err is:%v\n", err)
	}
}
