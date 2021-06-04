package handler

import (
	"ngepet-yuk/entity"
	"ngepet-yuk/helper"
	"ngepet-yuk/layer/mastery"

	"github.com/gin-gonic/gin"
)

type masteryHandle struct {
	masteryService mastery.Service
}

func NewMasteryHandler(masteryService mastery.Service) *masteryHandle {
	return &masteryHandle{masteryService}
}

func (h *masteryHandle) ShowMasteries(c *gin.Context) {
	masteries, err := h.masteryService.GetMasteries()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})
		c.JSON(500, responseError)

		return
	}

	response := helper.APIResponse("success get all mastery type", 200, "status OK", masteries)
	c.JSON(200, response)
}

func (h *masteryHandle) CreateMasteryHandler(c *gin.Context) {
	var inputMastery entity.MasteryInput

	if err := c.ShouldBindJSON(&inputMastery); err != nil {

		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newMastery, err := h.masteryService.SaveNewMastery(inputMastery)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}
	response := helper.APIResponse("success create mastery", 200, "Status OK", newMastery)
	c.JSON(200, response)
}

func (h *masteryHandle) UpdateMasteryByIDHandler(c *gin.Context) {
	id := c.Params.ByName("mastery_id")

	var updateMasteryInput entity.UpdateMasteryInput

	if err := c.ShouldBindJSON(&updateMasteryInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	mastery, err := h.masteryService.UpdateMasteryByID(id, updateMasteryInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update mastery by ID", 200, "success", mastery)
	c.JSON(200, response)
}
