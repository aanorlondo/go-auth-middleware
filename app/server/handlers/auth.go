package handlers

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"app/models"
	"app/utils"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the credentials
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Retrieve the user from the database by username
	user, err := models.GetUserByUsername(credentials.Username)
	if err != nil {
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}

	// Check if the user exists and the password is correct
	if user == nil || !checkPasswordHash(credentials.Password, user.Password) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generate a JWT token with user data
	tokenData := map[string]interface{}{
		"user_id":  user.ID,
		"username": user.Username,
		// Add other user data or claims as required
	}
	token, err := utils.GenerateJWTToken(tokenData)
	if err != nil {
		http.Error(w, "Error generating JWT token", http.StatusInternalServerError)
		return
	}

	// Return the token as a response
	response := map[string]string{"token": token}
	jsonResponse(w, response, http.StatusOK)
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the credentials
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Check if the username already exists in the database
	existingUser, err := models.GetUserByUsername(credentials.Username)
	if err != nil {
		http.Error(w, "Error checking username availability", http.StatusInternalServerError)
		return
	}
	if existingUser != nil {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	// Hash the password
	hashedPassword, err := hashPassword(credentials.Password)
	if err != nil {
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
	err = user.Save()
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Generate a JWT token with user data
	tokenData := map[string]interface{}{
		"user_id":  user.ID,
		"username": user.Username,
		// Add other user data or claims as required
	}
	token, err := utils.GenerateJWTToken(tokenData)
	if err != nil {
		http.Error(w, "Error generating JWT token", http.StatusInternalServerError)
		return
	}

	// Return the token as a response
	response := map[string]string{"token": token}
	jsonResponse(w, response, http.StatusCreated)
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
