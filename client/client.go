package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"

	"otp-auth/auth"
	"otp-auth/validation"
)

func main() {
	var idAttempts, otpAttempts int

	// Connect to the RPC server
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)

	// ✅ Phase 1: Ask for a valid identifier once
	var identifier string
	for {
		if idAttempts == 3 {
			fmt.Println("❌ Too many attempts. Please try again later.")
			return
		}

		fmt.Print("📧 Enter your email or phone: ")
		identifier, _ = reader.ReadString('\n')
		identifier = strings.TrimSpace(identifier)

		identifierType := validation.DetermineIdentifier(identifier)
		if identifierType == "unknown" {
			fmt.Println("❌ Invalid identifier. Please enter a valid email or phone number.")
			idAttempts++
		} else if identifierType == "email" && !validation.ValidEmail(identifier) {
			fmt.Println("❌ Invalid email format. Please enter a valid email address.")
			idAttempts++
		} else if identifierType == "phone" && !validation.ValidPhoneNumber(identifier) {
			fmt.Println("❌ Invalid phone number format. Please enter a valid phone number.")
			idAttempts++
		} else {
			// Valid identifier
			var otpRes auth.AuthResponse
			client.Call("AuthService.RequestOTP", auth.OTPRequest{Identifier: identifier}, &otpRes)
			fmt.Println("📨", otpRes.Message)
			break // Exit the identifier loop
		}
	}

	// ✅ Phase 2: OTP Verification
	for {
		if otpAttempts == 3 {
			fmt.Println("❌ Too many attempts. Please try again later.")
			return
		}

		fmt.Print("\n🔢 Enter the OTP you received: ")
		otp, _ := reader.ReadString('\n')
		otp = strings.TrimSpace(otp)

		if !validation.ValidOTP(otp) {
			fmt.Println("❌ Invalid OTP format. Please enter a 6-digit numeric OTP.")
			otpAttempts++
			continue
		}

		var verifyRes auth.AuthResponse
		client.Call("AuthService.VerifyOTP", auth.OTPVerify{
			Identifier: identifier,
			OTP:        otp,
		}, &verifyRes)

		if !verifyRes.Success {
			fmt.Println(verifyRes.Message)
			otpAttempts++
		} else {
			fmt.Println(verifyRes.Message)
			break
		}
	}
}
