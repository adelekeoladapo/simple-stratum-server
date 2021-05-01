package main

import (
	"github.com/joho/godotenv"
	"log"
	"luxot.tech/stratum/controller"
	"luxot.tech/stratum/db"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

func main() {
	/* Load configuration */
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Could not load app configuration. %s", err)
	}

	/* Database connection */
	db.InitPostgresDB(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	defer db.Db.Close()

	rpcServer := rpc.NewServer()
	mining := new(controller.Mining)
	err := rpcServer.Register(mining)
	if err != nil {
		log.Fatalf("Could not register service, %s", err)
	}

	rpcServer.HandleHTTP("/", "/debug")
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal("Could not listen on port 3000")
	}
	log.Println("Started RPC Handler on localhost:3000")

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		codec := jsonrpc.NewServerCodec(connection)
		go rpcServer.ServeCodec(codec)
	}

}
