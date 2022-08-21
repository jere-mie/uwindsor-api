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
	var ctx context.Context = context.TODO()
	var db *sqlx.DB = sqlx.MustConnect("sqlite3", "file:db.sqlite")
	var buildingservice services.BuildingService = services.NewBuildingService(db, ctx)
	var buildingcontroller controllers.BuildingController = controllers.NewBuildingController(buildingservice)
	var server *gin.Engine = gin.Default()

	basePath := server.Group("/api")
	buildingcontroller.RegisterBuildingRoutes(basePath)
	log.Fatal(server.Run(":9090"))

}
