package main

import (
	"fmt"
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
	listener, _ := net.Listen("tcp", ":1234")
	fmt.Println("🔐 OTP Auth Server running on port 1234...")
	rpc.Accept(listener)
}
