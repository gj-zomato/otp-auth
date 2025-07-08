package main

import (
	"fmt"
	"net"
	"net/rpc"

	"otp-login/auth"
)

func main() {
	auth.SetupRedis()

	rpc.Register(new(auth.AuthService))
	listener, _ := net.Listen("tcp", ":1234")
	fmt.Println("🔐 OTP Auth Server running on port 1234...")
	rpc.Accept(listener)
}
