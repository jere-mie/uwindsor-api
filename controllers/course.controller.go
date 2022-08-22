package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jere-mie/uwindsor-api/services"
)

type CourseController struct {
	CourseService services.CourseService
}

func NewCourseController(courseservice services.CourseService) CourseController {
	return CourseController{CourseService: courseservice}
}

func (cc *CourseController) GetCourse(ctx *gin.Context) {
	code := ctx.Param("code")
	course, err := cc.CourseService.GetCourse(&code)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, course)
}

func (cc *CourseController) GetCourses(ctx *gin.Context) {
	courses, err := cc.CourseService.GetCourses()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, courses)
}

func (cc *CourseController) RegisterCourseRoutes(rg *gin.RouterGroup) {
	courseroute := rg.Group("/course")
	courseroute.GET("/:code", cc.GetCourse)
	courseroute.GET("/", cc.GetCourses)
}
