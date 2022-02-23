package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input UserDTO
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"err": err.Error(),
		})
	}
	id, err := h.authService.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"err": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]int{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input UserDTO
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"err": err.Error(),
		})
	}
	token, err := h.authService.GenerateToken(input.Name, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"err": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
