package rest

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CreateSymbol(c *gin.Context) {
	var input AdmSymbolDTO
	err := c.BindJSON(&input)
	if err != nil {
		log.Fatalln("incorrect input")
	}
	id, err := h.admSymbolService.CreateSymbol(input)
	if err != nil {
		log.Fatalln("error with creating")
	}
	c.JSON(http.StatusCreated, map[string]int{
		"id": id,
	})
}

func (h *Handler) DeleteSymbol(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln("incorrect input")
	}
	err = h.admSymbolService.DeleteSymbol(id)
	if err != nil {
		log.Fatalln("error with deleting" + err.Error())
	}
}
