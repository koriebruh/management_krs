package conf

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("fIDJA90wju190djkqqlwpwqqwieuhd90q32")

type JWTClaim struct {
	NIM string `json:"nim"`
	jwt.RegisteredClaims
}
