package server

import (
	"app/server/handlers"
	"app/server/middleware"
	_ "embed"
	"net/http"

	"github.com/flowchartsman/swaggerui"
)

type Server struct {
	Router *http.ServeMux
}

//go:embed api/api.yaml
var apiSpec []byte

func NewServer() *Server {
	s := &Server{
		Router: http.NewServeMux(),
	}
	s.Router.Handle("/", swaggerui.Handler(apiSpec))                                                        // GET
	s.Router.HandleFunc("/login", handlers.LoginHandler)                                                    // POST
	s.Router.HandleFunc("/signup", handlers.SignupHandler)                                                  // POST
	s.Router.Handle("/user", middleware.JWTMiddleware(http.HandlerFunc(handlers.GetUserHandler)))           // GET
	s.Router.Handle("/user/update", middleware.JWTMiddleware(http.HandlerFunc(handlers.UpdateUserHandler))) // PUT

	return s
}
