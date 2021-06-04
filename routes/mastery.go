package routes

import (
	"ngepet-yuk/handler"
	"ngepet-yuk/layer/mastery"

	"github.com/gin-gonic/gin"
)

var (
	masteryRepository = mastery.NewRepository(DB)
	masteryService    = mastery.NewService(masteryRepository)
	masteryHandler    = handler.NewMasteryHandler(masteryService)
)

func MasteryRoute(r *gin.Engine) {
	r.GET("/masteries", masteryHandler.ShowMasteries)
	r.POST("/masteries", masteryHandler.CreateMasteryHandler)
	r.PUT("/masteries/:mastery_id", masteryHandler.UpdateMasteryByIDHandler)
}
