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

	// TODO: add a GET "/" route to display a front-end (help) page to list the known routes and expected payloads/headers

	s.Router.HandleFunc("/login", handlers.LoginHandler)                                                    // POST
	s.Router.HandleFunc("/signup", handlers.SignupHandler)                                                  // POST
	s.Router.Handle("/user", middleware.JWTMiddleware(http.HandlerFunc(handlers.GetUserHandler)))           // GET
	s.Router.Handle("/user/update", middleware.JWTMiddleware(http.HandlerFunc(handlers.UpdateUserHandler))) // PUT

	return s
}
