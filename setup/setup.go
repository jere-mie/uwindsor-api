package setup

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jere-mie/uwindsor-api/controllers"
	"github.com/jere-mie/uwindsor-api/services"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Start() {
	// db and server config
	var ctx context.Context = context.TODO()
	var db *sqlx.DB = sqlx.MustConnect("sqlite3", "file:db.sqlite")
	var server *gin.Engine = gin.Default()
	server.Use(cors.Default())
	server.Static("/assets", "./assets")
	server.LoadHTMLGlob("templates/*.html")

	// building stuff
	buildingservice := services.NewBuildingService(db, ctx)
	buildingcontroller := controllers.NewBuildingController(buildingservice)

	// staff stuff
	staffservice := services.NewStaffService(db, ctx)
	staffcontroller := controllers.NewStaffController(staffservice)

	// course stuff
	courseservice := services.NewCourseService(db, ctx)
	coursecontroller := controllers.NewCourseController(courseservice)

	// html routes
	controllers.RegisterHTMLRoutes(server.Group("/"))

	// REST API Routes
	v1Path := server.Group("/v1")
	buildingcontroller.RegisterBuildingRoutes(v1Path)
	staffcontroller.RegisterStaffRoutes(v1Path)
	coursecontroller.RegisterCourseRoutes(v1Path)

	log.Fatal(server.Run(":9090"))

}
