package rest

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService        UserService
	admSymbolService   AdmSymbolService
	userAccountService UserAccountService
	dealService        DealService
}

func NewHandler(userService UserService, admSymbolService AdmSymbolService, userAccountService UserAccountService, dealService DealService) *Handler {
	return &Handler{authService: userService,
		admSymbolService:   admSymbolService,
		userAccountService: userAccountService,
		dealService:        dealService}
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
		admin := api.Group("/admin", h.userIdentity, h.isAdmin)
		{
			admin.DELETE("/smb/:id", h.deleteSymbol)
			admin.POST("/smb", h.createSymbol)
			admin.POST("/smb-price", h.setPrice)
		}

		mybank := api.Group("/mybank", h.userIdentity)
		{
			mybank.POST("/portfolio", h.createPortfolio)
			mybank.GET("/portfolio", h.getPortfolioList)
			mybank.GET("/portfolio/:id", h.getPortfolio)

		}

		deal := api.Group("/deal", h.userIdentity)
		{
			deal.GET("/share", h.getShareListInfo)
			deal.GET("/share/:id", h.getShareInfo)

			deal.POST("/buy", h.buyShares)
			deal.POST("/sell", h.sellShares)
		}

	}

	return router
}
