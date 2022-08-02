package main

import (
	"myapp/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	templateRepo := controllers.New()
	r.POST("/templates", templateRepo.CreateTemplate)
	r.GET("/templates", templateRepo.GetTemplates)
	r.GET("/templates/:id", templateRepo.GetTemplate)
	r.PUT("/templates/:id", templateRepo.UpdateTemplate)
	r.DELETE("/templates/:id", templateRepo.DeleteTemplate)

	return r
}
