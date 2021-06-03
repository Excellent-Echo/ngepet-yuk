package routes

import (
	"ngepet-yuk/handler"
	"ngepet-yuk/userTransaction"

	"github.com/gin-gonic/gin"
)

var (
	userTransactionRepository = userTransaction.NewRepository(DB)
	userTransactionService    = userTransaction.NewService(userTransactionRepository)
	userTransactionHandler    = handler.NewUserTransactionHandler(userTransactionService, authService)
)

func UserTransactionRoute(r *gin.Engine) {
	r.GET("/users/transaction/:user_id", userTransactionHandler.GetUserTransactionByUserIDHandler)
	r.POST("/users/transactions/add", userTransactionHandler.SaveNewUserTransactionHandler)
}
