package controllers

import (
	"equocenterback/pkg/models"
	"equocenterback/pkg/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PractitionerController struct {
	PractitionerRepository repositories.PractitionerRepository
}

func New(pratitionerService repositories.PractitionerRepository) PractitionerController {
	return PractitionerController {
		PractitionerRepository: pratitionerService,
	}
}

func (pc *PractitionerController) CreatePractitioner(ctx *gin.Context) {
	var practitioner models.Practitioner

	if err := ctx.ShouldBindJSON(&practitioner); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error message": err.Error()})
		return
	}

	err := pc.PractitionerRepository.CreatePractitioner(&practitioner)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success message": "Practitioner created successfully!"})
}

func (pc *PractitionerController) GetPractitioner(ctx *gin.Context) {
	document := ctx.Param("document")

	practitioner, err := pc.PractitionerRepository.GetPractitioner((&document))
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error message": err.Error() + " - Practitioner not found!"})
		return
	}

	ctx.JSON(http.StatusOK, practitioner)
}

func (pc *PractitionerController) GetAllPractitioners(ctx *gin.Context) {
	practitioners, err := pc.PractitionerRepository.GetAllPractitioners()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error message": err.Error() + " - No practitioners found!"})
		return
	}
	
	ctx.JSON(http.StatusOK, practitioners)
}

func (pc *PractitionerController) RegisterPractitionerRoutes(router *gin.RouterGroup) {
	practitionerroutes := router.Group("/practitioner")

	practitionerroutes.POST("/create", pc.CreatePractitioner)
	practitionerroutes.GET("/:document", pc.GetPractitioner)
	practitionerroutes.GET("/all", pc.GetAllPractitioners)
}