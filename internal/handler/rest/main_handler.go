package rest

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService      UserService
	admSymbolService AdmSymbolService
}

func NewHandler(userService UserService, admSymbolService AdmSymbolService) *Handler {
	return &Handler{authService: userService,
		admSymbolService: admSymbolService}
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
		admin := api.Group("/admin", h.isAdmin)
		{
			admin.POST("/sbm-create", h.CreateSymbol)
			admin.POST("/sbm-set-price/:id")
			admin.DELETE("/:id", h.DeleteSymbol)
		}
	}
	return router
}
