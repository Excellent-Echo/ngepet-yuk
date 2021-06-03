package handler

import (
	"log"
	"ngepet-yuk/auth"
	"ngepet-yuk/entity"
	"ngepet-yuk/helper"
	"ngepet-yuk/userDetail"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userDetailHandler struct {
	userDetailService userDetail.Service
	authService       auth.Service
}

func NewUserDetailHandler(userDetailService userDetail.Service, authService auth.Service) *userDetailHandler {
	return &userDetailHandler{userDetailService, authService}
}

func (h *userDetailHandler) GetUserDetailByUserIDHandler(c *gin.Context) {

	userData := int(c.MustGet("currentUser").(int))

	if userData == 0 {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user not authorize / not login"})

		c.JSON(401, responseError)
		return
	}

	userID := strconv.Itoa(userData)

	userDetail, err := h.userDetailService.GetUserDetailByUserID(userID)

	if err != nil {
		responseError := helper.APIResponse("status unauthorize", 401, "error", gin.H{"error": err.Error()})

		c.JSON(401, responseError)
		return
	}

	response := helper.APIResponse("success get user detail by user ID", 200, "success", userDetail)
	c.JSON(200, response)
}

func (h *userDetailHandler) SaveNewUserDetailHandler(c *gin.Context) {
	userData := int(c.MustGet("currentUser").(int))

	if userData == 0 {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user not authorize / not login"})

		c.JSON(401, responseError)
		return
	}

	userID := strconv.Itoa(userData)

	var inputUserDetail entity.UserDetailInput

	if err := c.ShouldBindJSON(&inputUserDetail); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	NewUserDetail, err := h.userDetailService.SaveNewUserDetail(inputUserDetail, userID)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success create new user Detail", 201, "success", NewUserDetail)
	c.JSON(201, response)
}

func (h *userDetailHandler) UpdateUserDetailByIDHandler(c *gin.Context) {
	userData := int(c.MustGet("currentUser").(int))

	if userData == 0 {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user not authorize / not login"})

		c.JSON(401, responseError)
		return
	}

	userID := strconv.Itoa(userData)

	dataUserDetail, err := h.userDetailService.GetUserDetailByUserID(userID)

	if err != nil {
		log.Println("line 99 error get user detail by useriD")
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	userDetailID := strconv.Itoa(dataUserDetail.ID)

	var updateUserDetailInput entity.UpdateUserDetailInput

	if err := c.ShouldBindJSON(&updateUserDetailInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	userDetail, err := h.userDetailService.UpdateUserDetailByID(userDetailID, updateUserDetailInput)

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update user Detail", 200, "success", userDetail)
	c.JSON(200, response)
}
