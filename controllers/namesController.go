package controllers

import (
	"strconv"

	"github.com/AlexeyStrekozov/effective_mobile_test/initializers"
	"github.com/AlexeyStrekozov/effective_mobile_test/models"
	"github.com/AlexeyStrekozov/effective_mobile_test/services"
	"github.com/gin-gonic/gin"
)

func NameCreate(c *gin.Context) {
	var body struct {
		Name       string `json:"name" validate:"required"`
		Surname    string `json:"surname" validate:"required"`
		Patronymic string `json:"patronymic,omitempty"`
	}

	c.Bind(&body)

	if body.Name != "" && body.Surname != "" {
		age := services.GetAge(body.Name)
		gender := services.GetGender(body.Name)
		nationality := services.GetNationality(body.Name)

		name := models.Name{
			Name:        body.Name,
			Surname:     body.Surname,
			Age:         age,
			Gender:      gender,
			Nationality: nationality,
			Patronymic:  body.Patronymic,
		}

		result := initializers.DB.Create(&name)

		if result.Error != nil {
			c.Status(400)
			return
		}

		c.JSON(200, gin.H{
			"name": name,
		})
	} else {
		c.JSON(400, gin.H{
			"error": "First and Last name are required",
		})
	}
}

func NamesIndex(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	if page <= 0 {
		page = 1
	}

	sizeStr := c.DefaultQuery("size", "10")
	pageSize, _ := strconv.Atoi(sizeStr)

	var names []models.Name

	initializers.DB.Scopes(services.NewPaginate(pageSize, page).PaginatedResult).Find(&names)

	c.JSON(200, gin.H{
		"names": names,
	})
}

func NameShow(c *gin.Context) {
	id := c.Param("id")

	var name models.Name

	initializers.DB.First(&name, "ID = ?", id)

	c.JSON(200, gin.H{
		"name": name,
	})
}

func NameUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name       string `json:"name" validate:"required"`
		Surname    string `json:"surname" validate:"required"`
		Patronymic string `json:"patronymic,omitempty"`
	}

	c.Bind(&body)

	if body.Name != "" && body.Surname != "" {
		var name models.Name
		initializers.DB.First(&name, id)

		initializers.DB.Model(&name).Updates(models.Name{
			Name:    body.Name,
			Surname: body.Surname,
		})

		c.JSON(200, gin.H{
			"name": name,
		})
	} else {
		c.JSON(400, gin.H{
			"error": "First and Last name are required",
		})
	}
}

func NameDelete(c *gin.Context) {
	id := c.Param("id")

	var name models.Name
	initializers.DB.Delete(&name, id)

	c.Status(200)
}
