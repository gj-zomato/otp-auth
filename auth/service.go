package auth

import (
	"fmt"
	"math/rand"
	"time"

	redis "github.com/redis/go-redis/v9"
)

// OTPRequest is used to request an OTP for a given identifier (email or phone).
type AuthService struct{}

// OTPRequest is used to request an OTP for a given identifier (email or phone).
func generateOTP() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// OTPVerify is used to verify the OTP for a given identifier.
func (a *AuthService) RequestOTP(req OTPRequest, res *AuthResponse) error {
	otp := generateOTP()
	key := "otp:" + req.Identifier

	err := Rdb.Set(Ctx, key, otp, 2*time.Minute).Err()
	if err != nil {
		res.Message = ("Error saving OTP: " + err.Error())
		res.Success = false
		return nil
	}

	fmt.Println("📲 OTP sent to", req.Identifier, ":", otp)
	res.Message = "OTP sent successfully"
	res.Success = true
	return nil
}

// VerifyOTP checks the provided OTP against the stored OTP for the identifier.
func (a *AuthService) VerifyOTP(req OTPVerify, res *AuthResponse) error {
	key := "otp:" + req.Identifier
	storedOTP, err := Rdb.Get(Ctx, key).Result()

	if err == redis.Nil {
		res.Message = "OTP expired or not found"
		res.Success = false
	} else if err != nil {
		res.Message = "Redis error"
		res.Success = false
	} else if storedOTP == req.OTP {
		res.Message = "✅ Login successful!"
		res.Success = true
		_ = Rdb.Del(Ctx, key)
	} else {
		res.Message = "❌ Incorrect OTP"
		res.Success = false
	}
	return nil
}
