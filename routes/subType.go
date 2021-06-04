package routes

import (
	"ngepet-yuk/handler"
	"ngepet-yuk/layer/subtype"

	"github.com/gin-gonic/gin"
)

var (
	subtypeRepository = subtype.NewRepository(DB)
	subtypeService    = subtype.NewService(subtypeRepository)
	subtypeHandler    = handler.NewSubTypeHandler(subtypeService)
)

func SubTypeRoute(r *gin.Engine) {
	r.GET("/sub-types", subtypeHandler.ShowAllSubType)
	r.POST("/sub-types", subtypeHandler.CreateSubtypeHandler)
	r.PUT("/sub-types/:subtype_id", subtypeHandler.UpdateSubtypeByIDHandler)
}
