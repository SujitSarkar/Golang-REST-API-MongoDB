package controllers

import (
	"net/http"

	"example.mongo-api/models"
	"example.mongo-api/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(userService services.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Param("name")
	user, err := uc.UserService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAllUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	ctx.JSON(200, "")
}

func (uc UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	userRoute.POST("/create", uc.CreateUser)
	userRoute.GET("/get/:name", uc.GetAllUser)
	userRoute.GET("/getall", uc.GetAllUser)
	userRoute.PATCH("/update", uc.UpdateUser)
	userRoute.DELETE("/delete", uc.DeleteUser)
}
