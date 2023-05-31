package middleware

import (
	"app/utils"
	"net/http"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the JWT token from the request headers or cookies
		tokenString := utils.ExtractTokenFromRequest(r)
		if tokenString == "" {
			http.Error(w, "Missing JWT token", http.StatusUnauthorized)
			return
		}

		// Verify and validate the JWT token
		token, err := utils.VerifyToken(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Invalid JWT token", http.StatusUnauthorized)
			return
		}

		// If the token is valid, call the next handler
		next.ServeHTTP(w, r)
	})
}
