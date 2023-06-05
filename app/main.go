package main

import (
	"app/config"
	"app/server"
	"app/utils"
	"context"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

var logger = utils.GetLogger()
var conf, err = config.LoadConfig()

func main() {
	// Test loading the server configuration
	if err != nil {
		logger.Error("Error reading server configuration: ", err)
		return
	}
	log.Println(("Successfully loaded server configuration."))

	// Test connecting to REDIS RBAC database
	redisOptions := &redis.Options{
		Addr:     conf.GetRedisURL(),
		Password: conf.GetRedisPassword(),
		DB:       0,
	}
	redisClient := redis.NewClient(redisOptions)
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		logger.Error("Failed to connect to Redis:", err)
	}
	log.Println("Connection to Redis OK.")
	defer redisClient.Close()

	// Start Server
	s := server.NewServer(redisClient)
	log.Print("Go-auth Server initialized.")
	log.Print("Server listening on port: 3456...")
	log.Fatal(http.ListenAndServe(":3456", s.Router))
}
