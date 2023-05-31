package handlers

import (
	"encoding/json"
	"net/http"

	"app/models"
	"app/utils"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the token from the request headers
	tokenString := utils.ExtractTokenFromRequest(r)
	if tokenString == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Verify the token
	token, err := utils.VerifyToken(tokenString)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get the claims from the token
	claims, err := utils.GetTokenClaims(token)
	if err != nil {
		http.Error(w, "Invalid token claims", http.StatusBadRequest)
		return
	}

	// Get the username from the claims
	username, ok := claims["username"].(string)
	if !ok {
		http.Error(w, "Invalid token claims", http.StatusBadRequest)
		return
	}

	// Retrieve the user from the database by username
	user, err := models.GetUserByUsername(username)
	if err != nil {
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}

	// Return the user data as a response
	jsonResponse(w, user, http.StatusOK)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the token from the request headers
	tokenString := utils.ExtractTokenFromRequest(r)
	if tokenString == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Verify the token
	token, err := utils.VerifyToken(tokenString)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get the claims from the token
	claims, err := utils.GetTokenClaims(token)
	if err != nil {
		http.Error(w, "Invalid token claims", http.StatusBadRequest)
		return
	}

	// Get the username from the claims
	username, ok := claims["username"].(string)
	if !ok {
		http.Error(w, "Invalid token claims", http.StatusBadRequest)
		return
	}

	// Retrieve the user from the database by username
	user, err := models.GetUserByUsername(username)
	if err != nil {
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}

	// Decode the JSON payload from the request body
	var updatedUser models.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Update the user fields
	user.Username = updatedUser.Username
	user.Password = updatedUser.Password

	// Save the updated user to the database
	err = user.Save()
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	// Return a success response
	jsonResponse(w, "User updated successfully", http.StatusOK)
}
