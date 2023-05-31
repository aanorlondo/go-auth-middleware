package utils

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"app/config"
)

var secretKey []byte

func init() {
	// Retrieve the secret key from the config
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}
	secretKey = []byte(cfg.GetAppSecretKey())
}

func ExtractTokenFromRequest(r *http.Request) string {
	// Check the Authorization header for the token
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		// The header value should be in the format: Bearer <token>
		authValue := strings.Split(authHeader, " ")
		if len(authValue) == 2 && authValue[0] == "Bearer" {
			return authValue[1]
		}
	}

	// Check the cookie for the token
	cookie, err := r.Cookie("token")
	if err == nil {
		return cookie.Value
	}

	// Check the query parameters for the token
	token := r.URL.Query().Get("token")
	if token != "" {
		return token
	}

	// Return an empty string if the token is not found
	return ""
}

func GenerateJWTToken(data map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
		// Add your custom claims or data here
	}
	for key, value := range data {
		claims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method used in the token
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func GetTokenClaims(token *jwt.Token) (map[string]interface{}, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
