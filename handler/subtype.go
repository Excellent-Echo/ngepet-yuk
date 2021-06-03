package handler

import (
	"ngepet-yuk/entity"
	"ngepet-yuk/helper"
	"ngepet-yuk/subtype"

	"github.com/gin-gonic/gin"
)

type subTypeHandler struct {
	subtypeService subtype.Service
}

func NewSubTypeHandler(subtypeService subtype.Service) *subTypeHandler {
	return &subTypeHandler{subtypeService}
}

func (h *subTypeHandler) ShowAllSubType(c *gin.Context) {
	subtypes, err := h.subtypeService.GetAllSubtype()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})
		c.JSON(500, responseError)

		return
	}

	response := helper.APIResponse("success get all subcription type", 200, "status OK", subtypes)
	c.JSON(200, response)
}

func (h *subTypeHandler) CreateSubtypeHandler(c *gin.Context) {
	var inputSubtype entity.SubTypeInput

	if err := c.ShouldBindJSON(&inputSubtype); err != nil {

		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newSubtype, err := h.subtypeService.SaveNewSubtype(inputSubtype)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}
	response := helper.APIResponse("success create subcription type", 200, "Status OK", newSubtype)
	c.JSON(200, response)
}

func (h *subTypeHandler) UpdateSubtypeByIDHandler(c *gin.Context) {
	id := c.Params.ByName("subtype_id")

	var updateSUbtypeInput entity.UpdateSubTypeInput

	if err := c.ShouldBindJSON(&updateSUbtypeInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	user, err := h.subtypeService.UpdateSubTypeByID(id, updateSUbtypeInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update Subcription Type by ID", 200, "success", user)
	c.JSON(200, response)
}
