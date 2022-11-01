package entrypoints

//go:generate mockgen -source=${GOFILE} -destination=$PWD/internal/mocks/$GOPACKAGE/mock_${GOFILE} -package=mocks

import (
	"context"
	"errors"
	"net/http"

	"permutation-game/api/cmd/internal/entities"
	"permutation-game/api/cmd/internal/usecases"

	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// Interface.
type Executer interface {
	Execute(c *gin.Context)
}

type entrypointsImpl struct {
	uc usecases.UseCase
}

func (ep entrypointsImpl) Pong(c *gin.Context) {
	log.Debugf("Pong")
	c.String(http.StatusOK, "pong")
}

// Execute executes the router.
func (ep entrypointsImpl) Execute(c *gin.Context) {
	log.Debugf("[Handler] Starting.")

	ctx := context.TODO()

	number, _, errP := validateParams(c)
	if errP != nil {
		log.Infof("[Handler] URI: %v. Headers: %v. Body: %v.",
			c.FullPath(), c.Request.Header, c.Request.Body)
		c.JSON(http.StatusBadRequest, errP.Error())

		return
	}

	m := entities.Build().SetNumber(number)
	newCtx := context.WithValue(ctx, entities.CtxKey, m)
	response, e := ep.uc.Execute(newCtx)

	if e != nil {
		log.Errorf("[Handler] Response Error: %s", e.Error())
		c.JSON(http.StatusNotFound, "Cannot made permutation")

		return
	}

	value := response.Value(entities.CtxKey)
	if value == nil {
		c.JSON(http.StatusInternalServerError, "")
	}

	m, ok := value.(entities.Context)
	if !ok {
		c.JSON(http.StatusInternalServerError, "")
	}

	if m.APIError() == "" {
		log.Infof("[Handler] Status: %d. %s",
			http.StatusOK, "")
		c.JSON(http.StatusOK, m.Result())

		return
	}

	if m.APIError() != "" {
		log.Errorf("[Handler] Response Error: %s", m.APIError())
		c.JSON(http.StatusInternalServerError, m.APIError())

		return
	}
}

func validateParams(c *gin.Context) (int64, string, error) {
	s, ok := c.GetQuery("number")
	if !ok {
		log.Errorf("[Handler validateParams] Error missing params number.")
		return 0, "", errors.New("query params 'number' is required")
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Errorf("[Handler validateParams] Error converting number:%v.", s)
		return i, "", err
	}

	firstname := c.DefaultQuery("firstname", "Anonymous")

	log.Debugf("[Handler] User %s", firstname)

	return i, s, nil
}
