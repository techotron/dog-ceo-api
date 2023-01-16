package controllers

import (
	"net/http"

	log "github.com/techotron/dog-ceo-api/backend/log"
	"github.com/gin-gonic/gin"
)


// GetInfo returns a basic info response to proof the server is up
func GetInfo(c *gin.Context) {
	log.Debug("GetInfo: ", c)
	c.JSON(http.StatusOK, "Server is up!")
}
