package rest

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) GetShareInfo(c *gin.Context) {
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

func (h *Handler) GetShareListInfo(c *gin.Context) {
	shares, err := h.dealService.GetShareListInfo()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, shares)
}
