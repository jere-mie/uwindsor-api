package setup

import (
	"context"
	"log"
	"net/http"

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
	server.Static("/assets", "./assets")
	server.LoadHTMLGlob("templates/*.html")

	// building stuff
	buildingservice := services.NewBuildingService(db, ctx)
	buildingcontroller := controllers.NewBuildingController(buildingservice)

	// staff stuff
	staffservice := services.NewStaffService(db, ctx)
	staffcontroller := controllers.NewStaffController(staffservice)

	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Home",
		})
	})

	// registering routes and running
	basePath := server.Group("/api")
	buildingcontroller.RegisterBuildingRoutes(basePath)
	staffcontroller.RegisterStaffRoutes(basePath)
	log.Fatal(server.Run(":9090"))

}
