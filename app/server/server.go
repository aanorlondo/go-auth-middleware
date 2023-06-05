package server

import (
	"app/server/handlers"
	"app/server/middleware"
	_ "embed"
	"net/http"

	"github.com/flowchartsman/swaggerui"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	Router *http.ServeMux
}

//go:embed api/api.yaml
var apiSpec []byte

func NewServer(redisClient *redis.Client) *Server {
	s := &Server{
		Router: http.NewServeMux(),
	}

	// GET
	s.Router.Handle("/", swaggerui.Handler(apiSpec))

	// POST
	s.Router.HandleFunc("/login", handlers.LoginHandler)

	// POST
	s.Router.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		handlers.SignupHandler(w, r, redisClient)
	})

	// GET
	s.Router.Handle("/user", middleware.JWTMiddleware(http.HandlerFunc(handlers.GetUserHandler)))

	// PUT
	s.Router.Handle("/user/update", middleware.JWTMiddleware(http.HandlerFunc(handlers.UpdateUserHandler)))

	// POST
	s.Router.HandleFunc("/user/promote", func(w http.ResponseWriter, r *http.Request) {
		handlers.PromoteUserHandler(w, r, redisClient)
	})

	// POST
	s.Router.HandleFunc("/user/demote", func(w http.ResponseWriter, r *http.Request) {
		handlers.DemoteUserHandler(w, r, redisClient)
	})

	// GET
	s.Router.HandleFunc("/user/check", func(w http.ResponseWriter, r *http.Request) {
		handlers.CheckUserHandler(w, r, redisClient)
	})

	return s
}
