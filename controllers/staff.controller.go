package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jere-mie/uwindsor-api/services"
)

type StaffController struct {
	StaffService services.StaffService
}

func NewStaffController(staffservice services.StaffService) StaffController {
	return StaffController{StaffService: staffservice}
}

func (sc *StaffController) GetStaffMember(ctx *gin.Context) {
	name := ctx.Param("name")
	staffMember, err := sc.StaffService.GetStaffMember(&name)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, staffMember)
}

func (sc *StaffController) GetAllStaff(ctx *gin.Context) {
	staff, err := sc.StaffService.GetAllStaff()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, staff)
}

func (sc *StaffController) RegisterStaffRoutes(rg *gin.RouterGroup) {
	staffroute := rg.Group("/staff")
	staffroute.GET("/:name", sc.GetStaffMember)
	staffroute.GET("/", sc.GetAllStaff)
}
