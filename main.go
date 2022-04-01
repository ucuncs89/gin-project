package main

import (
	"gin-project/controllers"
	"gin-project/middlewares"
	"gin-project/models"

	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"description": "REST API",
		})
	})
	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")

}
