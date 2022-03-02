package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input UserDTO
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse("Probably, you have an incorrect JSON input.", http.StatusBadRequest, err, c)
	}
	id, err := h.authService.CreateUser(input)
	if err != nil {
		newErrorResponse("Server error", http.StatusInternalServerError, errors.New("h.authService.CreateUser(input): "+err.Error()), c)
	}
	c.JSON(http.StatusOK, map[string]int{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input UserDTO
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse("Probably, you have an incorrect JSON input.", http.StatusBadRequest, err, c)
	}
	token, err := h.authService.GenerateToken(input.Name, input.Password)
	if err != nil {
		newErrorResponse("Server error", http.StatusInternalServerError, errors.New("h.authService.GenerateToken(input.Name, input.Password): "+err.Error()), c)
	}
	c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
