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
			admin.DELETE("/smb/:id", h.DeleteSymbol)
			admin.POST("/smb", h.CreateSymbol)
			admin.POST("/smb-set-price/:id")
		}
	}
	return router
}
