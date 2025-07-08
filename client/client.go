package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"

	"otp-login/auth"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("📧 Enter your email or phone: ")
	identifier, _ := reader.ReadString('\n')
	identifier = strings.TrimSpace(identifier)

	var otpRes auth.AuthResponse
	client.Call("AuthService.RequestOTP", auth.OTPRequest{Identifier: identifier}, &otpRes)
	fmt.Println("📨", otpRes.Message)

	fmt.Print("🔢 Enter the OTP you received: ")
	otp, _ := reader.ReadString('\n')
	otp = strings.TrimSpace(otp)

	var verifyRes auth.AuthResponse
	client.Call("AuthService.VerifyOTP", auth.OTPVerify{
		Identifier: identifier,
		OTP:        otp,
	}, &verifyRes)

	fmt.Println(verifyRes.Message)
}
