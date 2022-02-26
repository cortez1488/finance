package rest

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) createPortfolio(c *gin.Context) {
	var input PortfolioDTO
	userID := getUserID(c)

	err := c.BindJSON(&input)
	if err != nil {
		log.Fatalln("error with binding" + err.Error())
	}

	id, err := h.userAccountService.CreatePortfolio(int(userID), input)
	if err != nil {
		log.Fatalln("error with service" + err.Error())
	}
	c.JSON(http.StatusCreated, map[string]int{
		"id": id,
	})
}

func (h *Handler) getPortfolio(c *gin.Context) {
	userID := int(getUserID(c))
	portfolioID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}

	portfolio, err := h.userAccountService.GetPortfolio(userID, portfolioID)
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, portfolio)
}

func (h *Handler) getPortfolioList(c *gin.Context) {
	userID := int(getUserID(c))
	portfolio, err := h.userAccountService.GetPortfolioList(userID)
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, portfolio)
}
