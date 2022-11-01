package main

import (
	"os"

	"permutation-game/api/cmd/internal/entrypoints"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := run(port); err != nil {
		log.Errorf("error running server: %v", err)
	}
}

func run(port string) error {
	// Default Router includes, attributes filter, jsonp and pprof middlewares.
	Router := gin.Default()

	entrypoints.MapURL(Router)
	log.SetLevel(log.DebugLevel)
	log.Debug("Api Start.")

	return Router.Run(":" + port)
}
