package main

import (
	"log"
	"luxot.tech/stratum/controller"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {

	rpcServer := rpc.NewServer()
	mining := new(controller.Mining)
	err := rpcServer.Register(mining)
	if err != nil {
		log.Fatalf("Could not register service, %s", err)
	}

	rpcServer.HandleHTTP("/", "")
	listener, err := net.Listen("tcp", "3000")
	if err != nil {
		log.Fatal("Could not listen on port 3000")
	}

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		codec := jsonrpc.NewServerCodec(connection)
		go rpcServer.ServeCodec(codec)
	}

}
