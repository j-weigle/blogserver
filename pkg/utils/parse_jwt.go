// Package utils offers utility functions for parsing information from headers and cookies
package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

// JWTMetadata contains any metadata about a JWT that is contained in its claims
type JWTMetadata struct {
	Expiration int64
}

// ParseJWTClaims extracts metadata that is contained within a JWT's claims
func ParseJWTClaims(claims jwt.MapClaims) (*JWTMetadata, error) {
	claimsOk := []bool

	exp, ok := claims["exp"]
	claimsOk = append(claimsOk, ok)

	if allClaimsOk(claimsOk) {
		expiration := int64(exp.(float64))

		return &JWTMetadata{
			Expiration: expiration,
		}, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}

func allClaimsOk(claimsOk []bool) bool {
	for _, ok := range claimsOk {
		if !ok {
			return false
		}
	}

	return true
}
