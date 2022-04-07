package controllers

import (
	"fasta/app/models"
	"fasta/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(service services.UserService) UserController {
	return UserController{
		UserService: service,
	}
}

func (controller *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := controller.UserService.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Param("name")
	user, err := controller.UserService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (controller *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := controller.UserService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (controller *UserController) UpdateUser(ctx *gin.Context) {
	var user *models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := controller.UserService.UpdateUser(user); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *UserController) DeleteUser(ctx *gin.Context) {
	name := ctx.Param("name")
	if err := controller.UserService.DeleteUser(&name); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (controller *UserController) RegisterRoutes(rg *gin.RouterGroup) {
	routes := rg.Group("/user")
	routes.POST("/create", controller.CreateUser)
	routes.GET("/get/:name", controller.GetUser)
	routes.GET("/getall", controller.GetAllUsers)
	routes.PATCH("/update", controller.UpdateUser)
	routes.DELETE("/delete/:name", controller.DeleteUser)
}
