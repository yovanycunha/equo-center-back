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

func (a *ActivityController) GetAllActivities(ctx *gin.Context) {
	activities, err := a.ActivityRepository.GetAllActivities()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error message": err.Error() + " - no activities found"})
		return
	}

	ctx.JSON(http.StatusOK, activities)
}

func (a *ActivityController) GetActivity(ctx *gin.Context){
	id := ctx.Param("id")

	activity, err := a.ActivityRepository.GetActivity((&id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error message": err.Error() + " - no activity found"})
		return
	}

	ctx.JSON(http.StatusOK, activity)
}

func (ac *ActivityController) RegisterActivityRoutes(router *gin.RouterGroup) {
	activityroutes := router.Group("/activity")

	activityroutes.POST("/create", ac.CreateActivity)
	activityroutes.GET("/all", ac.GetAllActivities)
	activityroutes.GET("/:id", ac.GetActivity)
}