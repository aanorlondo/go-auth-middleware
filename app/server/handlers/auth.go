package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"

	"app/models"
	"app/utils"
)

var logger = utils.GetLogger()

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LOGIN HANDLER
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Handling login request")
	if r.Method != http.MethodPost {
		logger.Error("ERROR: ", r.Method, " method not supported. Expected: ", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Parse the request body to get the credentials
	logger.Info("Parsing request body...")
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		logger.Error("Invalid request: ", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// Retrieve the user from the database by username
	logger.Info("Retrieving user from database...")
	user, err := models.GetUserByUsername(credentials.Username)
	if err != nil {
		logger.Error("Error retrieving user: ", err)
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}
	// Check if the user exists and the password is correct
	if user == nil || !checkPasswordHash(credentials.Password, user.Password) {
		logger.Info("Invalid username or password")
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	// Generate a JWT token with user data
	logger.Info("Generating JWT token...")
	tokenData := map[string]interface{}{
		"username": user.Username,
		// Add other user data or claims as required
	}
	token, err := utils.GenerateJWTToken(tokenData)
	if err != nil {
		logger.Error("Error generating JWT token: ", err)
		http.Error(w, "Error generating JWT token", http.StatusInternalServerError)
		return
	}
	// Return the token as a response
	response := map[string]string{"token": token}
	jsonResponse(w, response, http.StatusOK)
}

// SIGNUP HANDER
func SignupHandler(w http.ResponseWriter, r *http.Request, redisClient *redis.Client) {
	logger.Info("Handling signup request")
	if r.Method != http.MethodPost {
		logger.Error("ERROR: ", r.Method, " method not supported. Expected: ", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Parse the request body to get the credentials
	logger.Info("Parsing request body...")
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		logger.Error("Invalid request: ", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// Check if the username already exists in the database
	logger.Info("Checking username availability...")
	existingUser, err := models.GetUserByUsername(credentials.Username)
	if err != nil {
		logger.Error("Error checking username availability: ", err)
		http.Error(w, "Error checking username availability", http.StatusInternalServerError)
		return
	}
	if existingUser != nil {
		logger.Info("Username already exists")
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}
	// Hash the password
	hashedPassword, err := hashPassword(credentials.Password)
	if err != nil {
		logger.Error("Error hashing password: ", err)
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	// Create a new user
	user := &models.User{
		Username: credentials.Username,
		Password: hashedPassword,
	}
	// Save the user to the database
	err = user.Save()
	if err != nil {
		logger.Error("Error creating user: ", err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}
	// Set the user's role in Redis
	logger.Info("Writing user role in Redis...")
	err = redisClient.Set(context.Background(), "user:role:"+user.Username, "readOnly", 0).Err()
	if err != nil {
		logger.Error("Error setting user role in Redis: ", err)
		http.Error(w, "Error setting user role in Redis", http.StatusInternalServerError)
		return
	}
	// Generate a JWT token with user data
	tokenData := map[string]interface{}{
		"username": user.Username,
	}
	token, err := utils.GenerateJWTToken(tokenData)
	if err != nil {
		logger.Error("Error generating JWT token: ", err)
		http.Error(w, "Error generating JWT token", http.StatusInternalServerError)
		return
	}
	// Return the token as a response
	response := map[string]string{"token": token}
	jsonResponse(w, response, http.StatusCreated)
}

// PASSWORD HASH
func hashPassword(password string) (string, error) {
	logger.Info("Hashing password...")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Error hashing password: ", err)
		return "", err
	}
	return string(hashedPassword), nil
}

func checkPasswordHash(password, hash string) bool {
	logger.Info("Checking password hash...")
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// JSON RESPONSE
func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	logger.Info("Creating JSON response...")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		logger.Error("Error encoding JSON response: ", err)
	}
}
