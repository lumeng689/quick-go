package main

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
	"quick-go/api/thrift/gen-go/tutorial"
)

func handleClient(client *tutorial.CalculatorClient) {
	res, _ := client.Add(context.Background(), 13, 25)
	fmt.Println("result is ", res)
}

func SimpleClient() {
	var transport thrift.TTransport
	transport, _ = thrift.NewTSocketConf("localhost:8090", &thrift.TConfiguration{
		ConnectTimeout: time.Second, // Use 0 for no timeout
		SocketTimeout:  time.Second, // Use 0 for no timeout
	})

	conf := &thrift.TConfiguration{
		ConnectTimeout: time.Second,
		SocketTimeout:  time.Second,

		MaxFrameSize: 1024 * 256,

		TBinaryStrictRead:  thrift.BoolPtr(true),
		TBinaryStrictWrite: thrift.BoolPtr(true),
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(conf)

	transportFactory := thrift.NewTTransportFactory()
	transport, _ = transportFactory.GetTransport(transport)
	defer transport.Close()
	transport.Open()

	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)

	handleClient(tutorial.NewCalculatorClient(thrift.NewTStandardClient(iprot, oprot)))
}
