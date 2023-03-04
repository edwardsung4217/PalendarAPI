package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type events struct {
	message    string
	otherField bool
}

func eventsHandler(c *gin.Context) {

	respObject := &events{
		message:    "events",
		otherField: true,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": respObject.message,
	})
}
