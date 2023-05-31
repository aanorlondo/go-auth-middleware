package handlers

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"app/models"
	"app/utils"
)

var logger = utils.GetLogger()

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Handling login request")

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
		"user_id":  user.ID,
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

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Handling signup request")

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
	logger.Info("Hashing password...")
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
		// Add other user data as required
	}

	// Save the user to the database
	logger.Info("Saving user to the database...")
	err = user.Save()
	if err != nil {
		logger.Error("Error creating user: ", err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Generate a JWT token with user data
	logger.Info("Generating JWT token...")
	tokenData := map[string]interface{}{
		"user_id":  user.ID,
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
	jsonResponse(w, response, http.StatusCreated)
}

func hashPassword(password string) (string, error) {
	logger.Info("Hashing password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Error hashing password: ", err)
		return "", err
	}
	return string(hashedPassword), nil
}

func checkPasswordHash(password, hash string) bool {
	logger.Info("Checking password hash")

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	logger.Info("Creating JSON response")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		logger.Error("Error encoding JSON response: ", err)
	}
}
