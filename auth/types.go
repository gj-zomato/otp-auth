package auth

type OTPRequest struct {
	Identifier string
}

type OTPVerify struct {
	Identifier string
	OTP        string
}

type AuthResponse struct {
	Message string
	Success bool
}
