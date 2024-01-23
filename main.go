package main

import (
	"github.com/AlexeyStrekozov/effective_mobile_test/controllers"
	"github.com/AlexeyStrekozov/effective_mobile_test/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.CreateHttpClient()
}

func main() {
	r := gin.Default()

	r.POST("/name", controllers.NameCreate)
	r.PUT("/name/:id", controllers.NameUpdate)
	r.GET("/name", controllers.NamesIndex)
	r.GET("/name/:id", controllers.NameShow)
	r.DELETE("/name/:id", controllers.NameDelete)

	r.Run()
}
