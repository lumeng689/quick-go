package server

import (
	"context"
	"log"
	"testing"
	"time"

	"google.golang.org/grpc"

	pb "quick-go/api/protobuf/helloworld"
)

func Test_SayHello1(t *testing.T) {
	address := "localhost:50051"

	opts := []grpc.DialOption{
		// credentials.
		grpc.WithInsecure(),
		//grpc.WithTransportCredentials(
		//	credentials.NewTLS(&tls.Config{InsecureSkipVerify: true}),
		//),
	}
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "zhangsan"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
