package server

import (
	"context"
	"fmt"

	pb "quick-go/api/protobuf/helloworld"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Printf("receive client greeter, name is: %v \n", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
