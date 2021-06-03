package routes

import (
	"ngepet-yuk/crypto"
	"ngepet-yuk/handler"

	"github.com/gin-gonic/gin"
)

var (
	cryptoService = crypto.NewService()
	cryptoHandler = handler.NewCryptoHandler(cryptoService)
)

func CryptoRoute(r *gin.Engine) {
	r.GET("/top-coins", cryptoHandler.ShowTopCoins)
}
