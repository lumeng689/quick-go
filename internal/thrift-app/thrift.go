package main

import (
	"flag"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"os"
	s "quick-go/internal/thrift-app/server"
	"time"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

func main() {
	fmt.Println("hello thrift...")
	flag.Usage = Usage

	//server := flag.Bool("server", true, "Run server")
	protocol := flag.String("P", "binary", "Specify the protocol (binary, compact, json, simplejson)")
	framed := flag.Bool("framed", false, "Use framed transport")
	buffered := flag.Bool("buffered", false, "Use buffered transport")
	addr := flag.String("addr", "localhost:9090", "Address to listen to")
	secure := flag.Bool("secure", false, "Use tls secure transport")

	flag.Parse()

	conf := &thrift.TConfiguration{
		ConnectTimeout:     time.Second,
		SocketTimeout:      time.Second,
		MaxFrameSize:       1024 * 256,
		TBinaryStrictRead:  thrift.BoolPtr(true),
		TBinaryStrictWrite: thrift.BoolPtr(true),
	}

	var protocolFactory thrift.TProtocolFactory
	switch *protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactoryConf(conf)
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactoryConf(conf)
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryConf(conf)
	default:
		fmt.Fprint(os.Stderr, "Invalid protocol specified", protocol, "\n")
		Usage()
		os.Exit(1)
	}

	var transportFactory thrift.TTransportFactory
	if *buffered {
		transportFactory = thrift.NewTBufferedTransportFactory(8192)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}

	if *framed {
		transportFactory = thrift.NewTFramedTransportFactoryConf(transportFactory, conf)
	}

	if err := s.RunServer(transportFactory, protocolFactory, *addr, *secure); err != nil {
		fmt.Println("error running server:", err)
	}

	//if err := s.RunClient(transportFactory, protocolFactory, *addr, *secure); err != nil {
	//	fmt.Println("error running client:", err)
	//}
}
