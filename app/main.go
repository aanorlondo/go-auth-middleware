package main

import (
	"app/config"
	"app/server"
	"app/utils"
	"log"
	"net/http"
)

var logger = utils.GetLogger()
var conf, err = config.LoadConfig()

func main() {
	if err != nil {
		logger.Error("Error reading server configuration: ", err)
		return
	}
	s := server.NewServer()
	logger.Info("Server initialized.")
	logger.Info("Serving configuration: ", conf)
	logger.Info("Server listening on port: 3456")
	log.Fatal(http.ListenAndServe(":3456", s.Router))
}
