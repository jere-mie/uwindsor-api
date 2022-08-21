package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jere-mie/uwindsor-api/services"
)

type BuildingController struct {
	BuildingService services.BuildingService
}

func NewBuildingController(buildingservice services.BuildingService) BuildingController {
	return BuildingController{BuildingService: buildingservice}
}

func (bc *BuildingController) GetBuilding(ctx *gin.Context) {
	name := ctx.Param("name")
	building, err := bc.BuildingService.GetBuilding(&name)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, building)
}

func (bc *BuildingController) GetBuildings(ctx *gin.Context) {
	buildings, err := bc.BuildingService.GetBuildings()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, buildings)
}

func (bc *BuildingController) RegisterBuildingRoutes(rg *gin.RouterGroup) {
	buildingroute := rg.Group("/building")
	buildingroute.GET("/get/:name", bc.GetBuilding)
	buildingroute.GET("/getall", bc.GetBuildings)
}
