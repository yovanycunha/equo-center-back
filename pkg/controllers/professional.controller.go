package controllers

import (
	"equocenterback/pkg/models"
	"equocenterback/pkg/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfessionalController struct {
	ProfessionalRepository repositories.ProfessionalRepository
}

func NewProfessionalController(professionalService repositories.ProfessionalRepository) ProfessionalController {
	return ProfessionalController{
		ProfessionalRepository: professionalService,
	}
}

func (pc *ProfessionalController) CreateProfessional(ctx *gin.Context) {
	var professional models.Professional

	if err := ctx.ShouldBindJSON(&professional); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error message": err.Error()})
		return
	}

	err := pc.ProfessionalRepository.CreateProfessional(&professional)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "professional created successfully"})
}

func (pc *ProfessionalController) RegisterProfessionalRoutes(router *gin.RouterGroup){
	professionalroutes := router.Group("/professional")

	professionalroutes.POST("/create", pc.CreateProfessional)
}