package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createPortfolio(c *gin.Context) {
	var input PortfolioDTO
	userID := getUserID(c)

	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse("Probably, you have an incorrect JSON input.", http.StatusBadRequest, err, c)
	}

	id, err := h.userAccountService.CreatePortfolio(int(userID), input)
	if err != nil {
		newErrorResponse("Server error", http.StatusInternalServerError, errors.New("h.userAccountService.CreatePortfolio(): "+err.Error()), c)
	}
	c.JSON(http.StatusCreated, map[string]int{
		"id": id,
	})
}

func (h *Handler) getPortfolio(c *gin.Context) {
	userID := int(getUserID(c))
	portfolioID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse("Unknown server error", http.StatusInternalServerError, errors.New("unknown parameter error "+err.Error()), c)
	}

	portfolio, err := h.userAccountService.GetPortfolio(userID, portfolioID)
	if err != nil {
		newErrorResponse("Server error", http.StatusInternalServerError, errors.New("h.userAccountService.GetPortfolio(): "+err.Error()), c)
	}
	c.JSON(http.StatusOK, portfolio)
}

func (h *Handler) getPortfolioList(c *gin.Context) {
	userID := int(getUserID(c))
	portfolio, err := h.userAccountService.GetPortfolioList(userID)
	if err != nil {
		newErrorResponse("Server error", http.StatusInternalServerError, errors.New("h.userAccountService.GetPortfolioList(): "+err.Error()), c)
	}

	c.JSON(http.StatusOK, portfolio)
}
