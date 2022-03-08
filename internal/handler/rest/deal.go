package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getShareInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse("Probably, you have an incorrect JSON input.", http.StatusBadRequest, err, c)
	}

	share, err := h.dealService.GetShareInfo(id)
	if err != nil {
		newErrorResponse("Server error", http.StatusInternalServerError,
			errors.New("h.dealService.GetShareInfo(id): "+err.Error()), c)
	}
	c.JSON(http.StatusOK, share)
}

func (h *Handler) getShareListInfo(c *gin.Context) {
	shares, err := h.dealService.GetShareListInfo()
	if err != nil {
		newErrorResponse("Server error", http.StatusInternalServerError,
			errors.New("h.dealService.GetShareListInfo(): "+err.Error()), c)
	}
	c.JSON(http.StatusOK, shares)
}

func (h *Handler) buyShares(c *gin.Context) {
	var input CreateDealDTO
	userID := int(getUserID(c))

	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse("Probably, you have an incorrect JSON input.", http.StatusBadRequest, err, c)
	}
	money, err := h.dealService.BuyShares(input.ShareID, input.PortfolioID, userID, input.Quantity)
	if err != nil {
		newErrorResponse("Server error", http.StatusInternalServerError,
			errors.New("h.dealService.BuyShares(): "+err.Error()), c)
	}
	c.JSON(http.StatusOK, money)

}

func (h *Handler) sellShares(c *gin.Context) {
	var input CreateDealDTO
	userID := int(getUserID(c))

	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse("Probably, you have an incorrect JSON input.", http.StatusBadRequest, err, c)
	}
	money, err := h.dealService.SellShares(input.ShareID, input.PortfolioID, userID, input.Quantity)
	if err != nil {
		newErrorResponse("Server error", http.StatusInternalServerError,
			errors.New("h.dealService.SellShares(): "+err.Error()), c)
	}
	c.JSON(http.StatusOK, money)
}
