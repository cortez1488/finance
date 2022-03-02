package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createSymbol(c *gin.Context) {
	var input AdmSymbolDTO
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse("Probably, you have an incorrect JSON input.", http.StatusBadRequest, err, c)
	}
	id, err := h.admSymbolService.CreateSymbol(input)
	if err != nil {
		newErrorResponse("Server error", http.StatusInternalServerError, errors.New("h.admSymbolService.CreateSymbol(): "+err.Error()), c)
	}
	c.JSON(http.StatusCreated, map[string]int{
		"id": id,
	})
}

func (h *Handler) setPrice(c *gin.Context) {
	var input AdmPriceDTO
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse("Probably, you have an incorrect JSON input.", http.StatusBadRequest, err, c)
	}
	err = h.admSymbolService.SetPrice(input)
	if err != nil {
		newErrorResponse("Server error", http.StatusInternalServerError, errors.New("h.admSymbolService.SetPrice(): "+err.Error()), c)
	}
	c.Status(http.StatusCreated)
}

func (h *Handler) deleteSymbol(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse("Unknown error with id parameter", http.StatusInternalServerError, err, c)
	}
	err = h.admSymbolService.DeleteSymbol(id)
	if err != nil {
		newErrorResponse("Server error", http.StatusInternalServerError, errors.New("h.admSymbolService.DeleteSymbol(): "+err.Error()), c)
	}
}
