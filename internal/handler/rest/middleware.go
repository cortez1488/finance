package rest

import (
	"github.com/gin-gonic/gin"
	"log"
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
}

func getUserID(c *gin.Context) int64 {
	id, _ := c.Get(userCtx)
	return id.(int64)
}
