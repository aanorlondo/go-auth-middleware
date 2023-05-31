package server

import (
	"app/server/handlers"
	"app/server/middleware"
	"net/http"
)

type Server struct {
	Router *http.ServeMux
}

func NewServer() *Server {
	s := &Server{
		Router: http.NewServeMux(),
	}

	// Register your routes here
	s.Router.HandleFunc("/login", handlers.LoginHandler)
	s.Router.HandleFunc("/signup", handlers.SignupHandler)
	s.Router.Handle("/user", middleware.JWTMiddleware(http.HandlerFunc(handlers.GetUserHandler)))
	s.Router.Handle("/user/update", middleware.JWTMiddleware(http.HandlerFunc(handlers.UpdateUserHandler)))

	return s
}
