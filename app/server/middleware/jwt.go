package middleware

import (
	"app/utils"
	"net/http"
)

var logger = utils.GetLogger()

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("JWTMiddleware: Handling request")

		// Extract the JWT token from the request headers or cookies
		logger.Info("Middleware: Extracting JWT token from request...")
		tokenString := utils.ExtractTokenFromRequest(r)
		if tokenString == "" {
			logger.Error("Missing JWT token")
			http.Error(w, "Missing JWT token", http.StatusUnauthorized)
			return
		}

		// Verify and validate the JWT token
		logger.Info("Middleware: Verifying JWT token...")
		token, err := utils.VerifyToken(tokenString)
		if err != nil || !token.Valid {
			logger.Error("Invalid JWT token: ", err)
			http.Error(w, "Invalid JWT token", http.StatusUnauthorized)
			return
		}

		// If the token is valid, call the next handler
		logger.Info("JWT token is valid !")
		next.ServeHTTP(w, r)
	})
}
