package handlers

import (
	"encoding/json"
	"net/http"

	"app/models"
	"app/utils"
)

// GET USER HANDLER
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Handling get user request")
	if r.Method != http.MethodGet {
		logger.Error("ERROR: ", r.Method, " method not supported. Expected: ", http.MethodGet)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Extract the token from the request headers
	tokenString := utils.ExtractTokenFromRequest(r)
	if tokenString == "" {
		logger.Error("Unauthorized: tokenString is empty")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
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
	// Retrieve the user from the database by username
	user, err := models.GetUserByUsername(username)
	if err != nil {
		logger.Error("Error retrieving user: ", err)
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}
	// Return the user data as a response
	logger.Info("Returning user data")
	jsonResponse(w, user, http.StatusOK)
}

// UPDATE USER HANDLER
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Handling update user request")
	if r.Method != http.MethodPut {
		logger.Error("ERROR: ", r.Method, " method not supported. Expected: ", http.MethodPut)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// Extract the token from the request headers
	tokenString := utils.ExtractTokenFromRequest(r)
	if tokenString == "" {
		logger.Error("Unauthorized: tokenString is empty")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
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
	// Retrieve the user from the database by username
	user, err := models.GetUserByUsername(username)
	if err != nil {
		logger.Error("Error retrieving user: ", err)
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}
	// Decode the JSON payload from the request body
	logger.Info("Decoding JSON payload from request body...")
	var updatedUser models.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		logger.Error("Invalid request payload: ", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Hash the new password
	hashedPassword, err := hashPassword(updatedUser.Password)
	if err != nil {
		logger.Error("Error hashing new password: ", err)
		http.Error(w, "Error hashing new password", http.StatusInternalServerError)
		return
	}
	// Update the user fields
	user.Password = hashedPassword
	// Update or save the user in the database
	if user.Username != "" {
		logger.Info("User already exists. Updating fields...")
		err = user.Update()
		if err != nil {
			logger.Error("Error updating user: ", err)
			http.Error(w, "Error updating user", http.StatusInternalServerError)
			return
		}
		logger.Info("User updated successfully")
		jsonResponse(w, "User updated successfully", http.StatusOK)
	}
}
