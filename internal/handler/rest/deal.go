package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) getShareInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}

	share, err := h.dealService.GetShareInfo(id)
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, share)
}

func (h *Handler) getShareListInfo(c *gin.Context) {
	shares, err := h.dealService.GetShareListInfo()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, shares)
}

func (h *Handler) buyShares(c *gin.Context) {
	var input CreateDealDTO
	userID := int(getUserID(c))

	err := c.BindJSON(&input)
	if err != nil {
		log.Fatalln(errors.New("BindJSON(): " + err.Error()))
	}
	money, err := h.dealService.BuyShares(input.ShareID, input.PortfolioID, userID, input.Quantity)
	if err != nil {
		log.Fatalln(errors.New("h.dealService.BuyShares(): " + err.Error()))
	}
	c.JSON(http.StatusOK, money)

}

func (h *Handler) sellShares(c *gin.Context) {
	var input CreateDealDTO
	userID := int(getUserID(c))

	err := c.BindJSON(&input)
	if err != nil {
		log.Fatalln(err)
	}
	money, err := h.dealService.SellShares(input.ShareID, input.PortfolioID, userID, input.Quantity)
	if err != nil {
		log.Fatalln(errors.New("SERVICE SellShares(): " + err.Error()))
	}
	c.JSON(http.StatusOK, money)
}
