package rest

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService        UserService
	admSymbolService   AdmSymbolService
	userAccountService UserAccountService
}

func NewHandler(userService UserService, admSymbolService AdmSymbolService, userAccountService UserAccountService) *Handler {
	return &Handler{authService: userService,
		admSymbolService:   admSymbolService,
		userAccountService: userAccountService}
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
			admin.DELETE("/smb/:id", h.deleteSymbol)
			admin.POST("/smb", h.createSymbol)
			admin.POST("/smb-price", h.setPrice)
		}

		mybank := api.Group("/mybank")
		{
			mybank.POST("/portfolio", h.createPortfolio)
			mybank.GET("/portfolio", h.getPortfolioList)
			mybank.GET("/portfolio/:id", h.getPortfolio)

		}

	}

	return router
}
