package handler

import (
	"ngepet-yuk/category"
	"ngepet-yuk/entity"
	"ngepet-yuk/helper"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService category.Service
}

func NewCategoryHandler(categoryService category.Service) *categoryHandler {
	return &categoryHandler{categoryService}
}

func (h *categoryHandler) ShowCategories(c *gin.Context) {
	categories, err := h.categoryService.GetCategories()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})
		c.JSON(500, responseError)

		return
	}

	response := helper.APIResponse("success get all category type", 200, "status OK", categories)
	c.JSON(200, response)
}

func (h *categoryHandler) CreateCategoryHandler(c *gin.Context) {
	var inputCategory entity.CategoryInput

	if err := c.ShouldBindJSON(&inputCategory); err != nil {

		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	newCategory, err := h.categoryService.SaveNewCategory(inputCategory)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}
	response := helper.APIResponse("success create category type", 200, "Status OK", newCategory)
	c.JSON(200, response)
}

func (h *categoryHandler) UpdateCategoryByIDHandler(c *gin.Context) {
	id := c.Params.ByName("category_id")

	var updateCategoryInput entity.UpdateCategoryInput

	if err := c.ShouldBindJSON(&updateCategoryInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	category, err := h.categoryService.UpdateCategoryByID(id, updateCategoryInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update category by ID", 200, "success", category)
	c.JSON(200, response)
}
