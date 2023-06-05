package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"app/models"
	"app/utils"

	"github.com/redis/go-redis/v9"
)

type PromotionRequest struct {
	Username string `json:"username"`
}

func PromoteUserHandler(w http.ResponseWriter, r *http.Request, redisClient *redis.Client) {
	logger.Info("Handling user promotion request")
	if r.Method != http.MethodPost {
		logger.Error("Method Not Allowed")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Check if the request contains a valid token
	validToken := checkAuthToken(r)
	if !validToken {
		logger.Error("Invalid or missing token")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var request PromotionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Error("Invalid request: ", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// Check if the user exists
	user, err := models.GetUserByUsername(request.Username)
	if err != nil {
		logger.Error("Error promoting user: ", err)
		http.Error(w, "Error promoting user", http.StatusInternalServerError)
		return
	}
	if user == nil {
		logger.Error("User not found")
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	// Update user's role to "readWrite"
	logger.Info("Promoting user role...")
	err = redisClient.Set(context.Background(), "user:role:"+user.Username, "readWrite", 0).Err()
	if err != nil {
		logger.Error("Error promoting user: ", err)
		http.Error(w, "Error promoting user", http.StatusInternalServerError)
		return
	}
	jsonResponse(w, map[string]string{"message": "User promoted successfully"}, http.StatusOK)
}

func DemoteUserHandler(w http.ResponseWriter, r *http.Request, redisClient *redis.Client) {
	logger.Info("Handling user demotion request")
	if r.Method != http.MethodPost {
		logger.Error("Method Not Allowed")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Check if the request contains a valid token
	validToken := checkAuthToken(r)
	if !validToken {
		logger.Error("Invalid or missing token")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var request PromotionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Error("Invalid request: ", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// Check if the user exists
	user, err := models.GetUserByUsername(request.Username)
	if err != nil {
		logger.Error("Error demoting user: ", err)
		http.Error(w, "Error demoting user", http.StatusInternalServerError)
		return
	}
	if user == nil {
		logger.Error("User not found")
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	// Update user's role to "readOnly"
	logger.Info("Demoting user role...")
	err = redisClient.Set(context.Background(), "user:role:"+user.Username, "readOnly", 0).Err()
	if err != nil {
		logger.Error("Error demoting user: ", err)
		http.Error(w, "Error demoting user", http.StatusInternalServerError)
		return
	}
	jsonResponse(w, map[string]string{"message": "User demoted successfully"}, http.StatusOK)
}

func CheckUserHandler(w http.ResponseWriter, r *http.Request, redisClient *redis.Client) {
	if r.Method != http.MethodGet {
		logger.Error("Method Not Allowed")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Get the JWT token from the request header
	tokenString := utils.ExtractTokenFromRequest(r)
	if tokenString == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}
	// Verify the token
	token, err := utils.VerifyToken(tokenString)
	if err != nil {
		logger.Error("Unauthorized: ", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	// Get the claims from the token
	claims, err := utils.GetTokenClaims(token)
	if err != nil {
		logger.Error("Invalid token claims: ", err)
		http.Error(w, "Invalid token claims", http.StatusBadRequest)
		return
	}
	// Get the username from the claims
	logger.Info("Getting username from token claims...")
	username, ok := claims["username"].(string)
	if !ok {
		logger.Error("Invalid token claims: ", claims)
		http.Error(w, "Invalid token claims", http.StatusBadRequest)
		return
	}
	// Check if the user has "readWrite" privileges
	logger.Info("Checking user privileges...")
	role, err := redisClient.Get(context.Background(), "user:role:"+username).Result()
	if err != nil {
		http.Error(w, "Error checking user role", http.StatusInternalServerError)
		return
	}
	if role != "readWrite" {
		http.Error(w, "Insufficient privileges", http.StatusForbidden)
		return
	}
	jsonResponse(w, map[string]string{"message": "User has readWrite privileges"}, http.StatusOK)
}

func checkAuthToken(r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return false
	}
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return false
	}
	token := splitToken[1]
	secret := os.Getenv("API_ADMIN_TOKEN")
	return token == secret
}
