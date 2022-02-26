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
	if header == "" {
		log.Fatal("empty header")
	}
	id, err := h.authService.ParseToken(header)
	if err != nil {
		log.Fatal(err)
	}
	c.Set(userCtx, id)
}

func getUserID(c *gin.Context) int64 {
	id, exists := c.Get(userCtx)
	if !exists {
		log.Fatal("user don't exists")
	}
	return id.(int64)
}

func (h *Handler) isAdmin(c *gin.Context) {
	id := getUserID(c)
	if !h.authService.IsAdmin(id) {
		c.JSON(http.StatusForbidden, map[string]string{
			"error": "no access",
		})
	}
}
