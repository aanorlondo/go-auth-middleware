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
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}
	secretKey = []byte(cfg.GetAppSecretKey())
}

func ExtractTokenFromRequest(r *http.Request) string {
	logger.Info("Extracting token from request...")
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		authValue := strings.Split(authHeader, " ")
		if len(authValue) == 2 && authValue[0] == "Bearer" {
			logger.Info("Token extracted from Authorization header")
			return authValue[1]
		}
	}
	cookie, err := r.Cookie("token")
	if err == nil {
		logger.Info("Token extracted from cookie")
		return cookie.Value
	}
	token := r.URL.Query().Get("token")
	if token != "" {
		logger.Info("Token extracted from query parameters")
		return token
	}
	logger.Error("Token not found in request")
	return ""
}

func GenerateJWTToken(data map[string]interface{}) (string, error) {
	logger.Info("Generating JWT token...")
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	}
	for key, value := range data {
		claims[key] = value
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		logger.Error("Error generating JWT token: ", err)
		return "", err
	}
	return signedToken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	logger.Info("Verifying JWT token...")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.Error(("ERROR when verifying JWT toke: invalid signing method"))
			return nil, errors.New("invalid signing method")
		}
		return secretKey, nil
	})
	if err != nil {
		logger.Error("Error verifying JWT token: ", err)
		return nil, err
	}
	return token, nil
}

func GetTokenClaims(token *jwt.Token) (map[string]interface{}, error) {
	logger.Info("Getting token claims...")
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		logger.Error("Invalid token claims: ", claims)
		return nil, errors.New("invalid token claims")
	}
	return claims, nil
}
