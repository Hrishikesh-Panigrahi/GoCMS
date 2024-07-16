package controllers

import (
	"fmt"
	"net/http"

	"github.com/Hrishikesh-Panigrahi/GoCMS/connections"
	"github.com/Hrishikesh-Panigrahi/GoCMS/models"
	"github.com/Hrishikesh-Panigrahi/GoCMS/render"
	view404 "github.com/Hrishikesh-Panigrahi/GoCMS/templates/404"
	views "github.com/Hrishikesh-Panigrahi/GoCMS/templates/Admin"
	processedviews "github.com/Hrishikesh-Panigrahi/GoCMS/templates/Processed"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	if c.Request.Method == "GET" {
		render.Render(c, http.StatusOK, views.CreateUser())
		return
	}

	var userbody struct {
		Name     string
		Email    string
		Password string
		Role_ID  uint
	}
	if c.Bind(&userbody) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
	}

	user := models.User{
		Email:    userbody.Email,
		Password: userbody.Password,
		RoleID:   userbody.Role_ID,
	}

	result := connections.DB.Create(&user)

	if result.Error != nil {
		render.Render(c, http.StatusInternalServerError, processedviews.Failure("Error while creating the user", "user not created", result.Error.Error()))
	}
	render.Render(c, http.StatusInternalServerError, processedviews.Success("User Created Successfully", "", "", ""))
}

func DeleteUser(c *gin.Context) {
	fmt.Println("Delete User")
}

func UpdateUser(c *gin.Context) {
	if c.Request.Method == "GET" {
		id := c.Param("id")
		var user models.User
		result := connections.DB.First(&user, id)
		if result.Error != nil {
			render.Render(c, http.StatusNotFound, view404.Page404("User not found"))
		}
		render.Render(c, http.StatusOK, views.EditUser(user))
	}

	if c.Request.Method == "PUT" {
		name := c.Request.FormValue("name")
		fmt.Println(name)
	}

}

func UpdatePassword(c *gin.Context) {
	
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	result := connections.DB.First(&user, id)
	if result.Error != nil {
		render.Render(c, http.StatusNotFound, view404.Page404("User not found"))
	}
	render.Render(c, http.StatusOK, views.GetUser(user))
}

func GetUsers(c *gin.Context) {
	var users []models.User
	result := connections.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while fetching the users",
		})
		return
	}
	result = connections.DB.Preload("Role").Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while preloading the role",
		})
		return
	}

	render.Render(c, http.StatusOK, views.Users(users))
}

func GetUsersByRole(c *gin.Context) {
}
