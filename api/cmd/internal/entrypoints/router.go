package entrypoints

import (
	"permutation-game/api/cmd/internal/repositories"
	"permutation-game/api/cmd/internal/usecases"

	"github.com/gin-gonic/gin"
)

// MapURL default url mapper.
func MapURL(router *gin.Engine) {
	// Repo can be reused on multiple Usecases.
	repo := repositories.New()
	uc := usecases.New(repo)
	e := entrypointsImpl{
		uc: *uc,
	}

	router.GET("/ping", e.Pong)
	router.GET("/permutation/list", e.Execute)
}
