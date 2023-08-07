// main.go
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/alistairr/spartan/models"
)

var (
	router       *gin.Engine
	ListenAddr   = "localhost:8080"
	PostgresAddr = "localhost:5432"
)

func main() {
	// MIGRATIONS
	models.Migrate()
	// END MIGRATIONS

	router := gin.Default()

	router.Static("/static", "./static")

	router.LoadHTMLGlob("templates/*")
	log.Println("Post template load")

	// ROUTES
	router.GET("/", Home)
	router.POST("/", CreateProject)
	router.GET("/:url", ProjectList)
	// issue crud
	router.POST("/:url/issue", Issue)

	// Run server
	router.Run()
}

// Render either HTML or JSON depending on 'Accept' header
func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}

}
