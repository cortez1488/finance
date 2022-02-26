package rest

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) createSymbol(c *gin.Context) {
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

func (h *Handler) setPrice(c *gin.Context) {
	var input AdmPriceDTO
	err := c.BindJSON(&input)
	if err != nil {
		log.Fatal("incorrect input" + err.Error())
	}
	err = h.admSymbolService.SetPrice(input)
	if err != nil {
		log.Fatalln("error with creating" + err.Error())
	}
	c.Status(http.StatusCreated)
}

func (h *Handler) deleteSymbol(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln("incorrect input")
	}
	err = h.admSymbolService.DeleteSymbol(id)
	if err != nil {
		log.Fatalln("error with deleting" + err.Error())
	}
}
