package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authHeader)
	if header == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, "You are unauthorized.")
	}
	id, err := h.authService.ParseToken(header)
	if err != nil {
		newErrorResponse("Server error with parsing your auth token. ", http.StatusInternalServerError,
			errors.New("h.authService.ParseToken(header): "+err.Error()), c)
	}
	c.Set(userCtx, id)
}

func getUserID(c *gin.Context) int64 {
	id, exists := c.Get(userCtx)
	if !exists {
		newErrorResponse("Server auth error.", http.StatusInternalServerError,
			errors.New("no user context"), c)
	}
	return id.(int64)
}

func (h *Handler) isAdmin(c *gin.Context) {
	id := getUserID(c)
	if !h.authService.IsAdmin(id) {
		c.AbortWithStatusJSON(http.StatusForbidden, "No access.")
	}
}
