package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	readTimeout   = 10 * time.Second
	writeTimeout  = 10 * time.Second
	maxHeaderByte = 1 << 20
)

func main() {
	gin.ForceConsoleColor()

	router := gin.Default()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	r := router.Group("/permutation")

	r.GET("/ping", func(c *gin.Context) {
		handlePing(c)
	})

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderByte,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic("error on listen and server")
	}
}

func handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
