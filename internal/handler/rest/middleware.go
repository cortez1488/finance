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
	if header == "" {
		log.Fatal("empty header")
	}
	id, err := h.authService.ParseToken(header)
	if err != nil {
		log.Fatal(err)
	}
	c.Set(userCtx, id)
}

func (h *Handler) isAdmin(c *gin.Context) {
}

func getUserID(c *gin.Context) int64 {
	id, exists := c.Get(userCtx)
	if !exists {
		log.Fatal("user don't exists")
	}
	return id.(int64)
}
