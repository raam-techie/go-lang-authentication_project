package security

import (
	"os"
	"time"

	"crypto/rand"
	"encoding/base64"

	"github.com/golang-jwt/jwt"
)

// Generate JWT token after athenticate the user
func GenerateJWT(username string) (string, error) {

	secretKey := []byte(os.Getenv("SecretYouShouldHide"))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username :"] = username
	claims["expiration :"] = time.Now().Add(15 * time.Minute)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

// Generate Refresh token
func GenerateRefreshToken() (string, error) {
	// Generate a random byte slice to be used as the token
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes to base64 to create the token string
	refreshToken := base64.URLEncoding.EncodeToString(tokenBytes)

	return refreshToken, nil
}
