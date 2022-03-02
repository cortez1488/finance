package rest

import (
	"github.com/gin-gonic/gin"
	"log"
)

func newErrorResponse(clientMessage string, status int, err error, c *gin.Context) {
	log.Println(err)
	c.AbortWithStatusJSON(status, map[string]string{
		"error": clientMessage,
	})
}
