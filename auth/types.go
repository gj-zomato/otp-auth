package auth

type AuthService struct{}

type OTPRequest struct {
	Identifier string
	OTP        string
}

type OTPVerify struct {
	Identifier string
	OTP        string
}

type AuthResponse struct {
	Message string
	Success bool
}
