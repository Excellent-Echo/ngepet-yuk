package handler

import (
	"ngepet-yuk/crypto"
	"ngepet-yuk/helper"

	"github.com/gin-gonic/gin"
)

type cryptoHandler struct {
	coinService crypto.Service
}

func NewCryptoHandler(coinService crypto.Service) *cryptoHandler {
	return &cryptoHandler{coinService}
}

func (h *cryptoHandler) ShowTopCoins(c *gin.Context) {
	crypto, err := h.coinService.GetTopCoins()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})
		c.JSON(500, responseError)

		return
	}

	response := helper.APIResponse("success get all top coins", 200, "status OK", crypto)
	c.JSON(200, response)
}
