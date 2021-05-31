package handler

import (
	"log"
	"net/http"
	"time"

	"ngepet-yuk/config"
	"ngepet-yuk/entity"
	"ngepet-yuk/helper"
	"ngepet-yuk/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

// SHOW ALL USER
func (h *userHandler) ShowAllUser(c *gin.Context) {
	users, err := h.userService.GetAllUser()

	if err != nil {
		responseError := helper.APIResponse("internal server error", http.StatusOK, "error", gin.H{"errors": err.Error()})

		c.JSON(http.StatusOK, responseError)
		return
	}

	response := helper.APIResponse("success get all user", http.StatusOK, "status OK", users)
	c.JSON(http.StatusOK, response)
}

// CREATE NEW USER OR REGISTER
func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var inputUser entity.UserInput

	if err := c.ShouldBindJSON(&inputUser); err != nil {

		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", http.StatusOK, "bad request", gin.H{"errors": splitError})

		c.JSON(http.StatusOK, responseError)
		return
	}

	newUser, err := h.userService.SaveNewUser(inputUser)
	if err != nil {
		responseError := helper.APIResponse("internal server error", http.StatusOK, "error", gin.H{"errors": err.Error()})

		c.JSON(http.StatusOK, responseError)
		return
	}
	response := helper.APIResponse("success create new User", http.StatusOK, "Status OK", newUser)
	c.JSON(http.StatusOK, response)
}

var DB = config.Connection()

// YANG INI SABAR DLU BELUM DITERAPIN CLEAN ARSITEK
func (h *userHandler) GetUserByID(c *gin.Context) {
	var users entity.User

	id := c.Params.ByName("user_id")

	if err := DB.Where("user_id = ?", id).Find(&users).Error; err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, users)
}

func HandleDeleteUser(c *gin.Context) {

	id := c.Params.ByName("user_id")

	if err := DB.Where("user_id = ?", id).Delete(&entity.User{}).Error; err != nil {
		log.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": id,
		"message": "success delete",
	})
}

func HandleUpdateUser(c *gin.Context) {
	var user entity.User
	var userInput entity.UserInput

	id := c.Params.ByName("user_id")
	DB.Where("user_id = ?", id).Find(&user)

	if err := c.ShouldBindJSON(&userInput); err != nil {
		log.Println(err.Error())
		return
	}
	user.UserName = userInput.UserName
	user.Email = userInput.Email
	user.Password = userInput.Password
	user.UpdatedAt = time.Now()

	if err := DB.Where("user_id = ?", id).Save(&user).Error; err != nil {
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
