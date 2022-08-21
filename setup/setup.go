package setup

import (
	"context"
	"log"

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

	// building stuff
	buildingservice := services.NewBuildingService(db, ctx)
	buildingcontroller := controllers.NewBuildingController(buildingservice)

	// staff stuff
	staffservice := services.NewStaffService(db, ctx)
	staffcontroller := controllers.NewStaffController(staffservice)

	// registering routes and running
	basePath := server.Group("/api")
	buildingcontroller.RegisterBuildingRoutes(basePath)
	staffcontroller.RegisterStaffRoutes(basePath)
	log.Fatal(server.Run(":9090"))

}
