package main

import (
	"github.com/AlexeyStrekozov/effective_mobile_test/initializers"
	"github.com/AlexeyStrekozov/effective_mobile_test/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Name{})
}
