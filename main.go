package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jere-mie/uwindsor-api/controllers"
	"github.com/jere-mie/uwindsor-api/services"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	ctx                context.Context
	db                 *sqlx.DB
	buildingservice    services.BuildingService
	buildingcontroller controllers.BuildingController
	server             *gin.Engine
	err                error
)

func init() {
	ctx = context.TODO()
	db = sqlx.MustConnect("sqlite3", "file:db.sqlite")
	buildingservice = services.NewBuildingService(db, ctx)
	buildingcontroller = controllers.NewBuildingController(buildingservice)
	server = gin.Default()
}

func main() {
	basePath := server.Group("/api")
	buildingcontroller.RegisterBuildingRoutes(basePath)
	log.Fatal(server.Run(":9090"))
}
