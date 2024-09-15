package controllers

import (
	"equocenterback/pkg/models"
	"equocenterback/pkg/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ActivityController struct {
	ActivityRepository repositories.ActivityRepository
}

func NewActivityController(activityRepository repositories.ActivityRepository) ActivityController {
	return ActivityController{
		ActivityRepository: activityRepository,
	}
}

func (a *ActivityController) CreateActivity(ctx *gin.Context) {
	var activity models.Activity

	if err := ctx.ShouldBindJSON(&activity); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error message": err.Error()})
		return
	}

	err := a.ActivityRepository.CreateActivity(&activity)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error message": err.Error()})
		return 
	}

	ctx.JSON(http.StatusCreated, gin.H{"message":"activity created successfully"})
}

func (ac *ActivityController) RegisterActivityRoutes(router *gin.RouterGroup) {
	activityroutes := router.Group("/activity")

	activityroutes.POST("/create", ac.CreateActivity)
}