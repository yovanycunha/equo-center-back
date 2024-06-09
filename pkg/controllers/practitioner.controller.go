package controllers

import (
	"equocenterback/pkg/models"
	"equocenterback/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PractitionerController struct {
	PractitionerService services.PractitionerService
}

func New(pratitionerService services.PractitionerService) PractitionerController {
	return PractitionerController {
		PractitionerService: pratitionerService,
	}
}

func (pc *PractitionerController) CreatePractitioner(ctx *gin.Context) {
	var practitioner models.Practitioner

	if err := ctx.ShouldBindJSON(&practitioner); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error message": err.Error()})
		return
	}

	err := pc.PractitionerService.CreatePractitioner(&practitioner)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success message": "Practitioner created successfully!"})
}

func (pc *PractitionerController) RegisterPractitionerRoutes(router *gin.RouterGroup) {
	practitionerroutes := router.Group("/practitioner")

	practitionerroutes.POST("/create", pc.CreatePractitioner)
}