package main

import (
	"fmt"
	"log"
	"luxot.tech/stratum/dto"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	connection, err := net.Dial("tcp", ":3000")
	if err != nil {
		log.Fatalf("Could not connect to localhost:3000. %s", err)
	}
	defer connection.Close()
	client := jsonrpc.NewClient(connection)

	var authorizationResponse bool
	if err := client.Call("Mining.Authorize", dto.AuthorizeRequest{"dapoadeleke", "password123"}, &authorizationResponse); err != nil {
		log.Fatalf("An error occurred, %s", err)
	}

	fmt.Printf("Authorization response: %v", authorizationResponse)
	fmt.Println()

}
