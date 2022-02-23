package rest

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService UserService
}

func NewHandler(service UserService) *Handler {
	return &Handler{authService: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api")
	{
		api.POST("", h.userIdentity)
	}
	return router
}
