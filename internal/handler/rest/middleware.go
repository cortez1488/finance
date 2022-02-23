package rest

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authHeader)
	id, err := h.authService.ParseToken(header)
	if err != nil {
		log.Fatal(err)
	}
	c.Set(userCtx, id)
	c.JSON(http.StatusOK, map[string]int64{
		"id": getUserID(c),
	})
}

func getUserID(c *gin.Context) int64 {
	id, _ := c.Get(userCtx)
	return id.(int64)
}
