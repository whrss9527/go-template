package util

import jwtv4 "github.com/golang-jwt/jwt/v4"

type AccountClaims struct {
	jwtv4.RegisteredClaims
	Channel string `json:"channel"`
}
