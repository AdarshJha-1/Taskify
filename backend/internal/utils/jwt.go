package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// var privateKey *ecdsa.PrivateKey

func CreateJWT(userId string) (string, error) {

	// Create a new token with claims
	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token with the secret key
	token, err := claim.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyJWT(token string) (map[string]interface{}, error) {
	// Parse and validate the token
	verifiedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil // Provide the secret key for validation
	})

	if err != nil {
		return nil, err
	}

	if !verifiedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Extract claims from the token
	claim, _ := verifiedToken.Claims.(jwt.MapClaims)

	return claim, nil // Return the claims
}
