package routes

import (
	"ngepet-yuk/category"
	"ngepet-yuk/handler"

	"github.com/gin-gonic/gin"
)

var (
	categoryRepository = category.NewRepository(DB)
	categoryService    = category.NewService(categoryRepository)
	categoryHandler    = handler.NewCategoryHandler(categoryService)
)

func CategoryRoute(r *gin.Engine) {
	categoryRepository = category.NewRepository(DB)
	categoryService = category.NewService(categoryRepository)
	categoryHandler = handler.NewCategoryHandler(categoryService)
}
