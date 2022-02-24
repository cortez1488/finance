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
	api := router.Group("/api", h.userIdentity)
	{
		admin := api.Group("/admin", h.userIdentity)
		{
			admin.GET("/check", h.isAdmin)
			admin.POST("/sbm-create/:id")
			admin.POST("/sbm-set-price/:id")
			admin.DELETE("/:id")
		}
	}
	return router
}
