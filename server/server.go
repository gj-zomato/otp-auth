package main

import (
	"log"
	"net"
	"net/rpc"

	"otp-auth/auth"
)

func main() {
	// Initialize the Redis client and register the AuthService
	auth.SetupRedis()
	// Register the AuthService with the RPC server
	rpc.Register(new(auth.AuthService))

	// Start the RPC server on port 1234
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen error:", err)
	}
	log.Println("RPC server listening on port 1234")
	defer listener.Close()

	// Accept connections and serve them concurrently
	for {
		// Accept a new connection
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
