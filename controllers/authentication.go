package controllers

import (
	"net/http"

	"github.com/Hrishikesh-Panigrahi/GoCMS/connections"
	"github.com/Hrishikesh-Panigrahi/GoCMS/models"
	"github.com/Hrishikesh-Panigrahi/GoCMS/render"
	"github.com/Hrishikesh-Panigrahi/GoCMS/services"

	authViews "github.com/Hrishikesh-Panigrahi/GoCMS/templates/authentication"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	if c.Request.Method == "GET" {
		render.Render(c, http.StatusOK, authViews.Index())
	}
	if c.Request.Method == "POST" {
		var body struct {
			Email    string
			Password string
		}

		if c.Bind(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid input",
			})
			return
		}

		// find the user
		var user models.User
		connections.DB.First(&user, "email = ?", body.Email)

		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid credentials",
			})
			return
		}

		ValidatePassword(user.Password, body.Password, c)

		//set cookies and jwt token
		services.JwtToken(c, user)
		render.Redirect(c, "/user/post", http.StatusFound)
	}

}

func Register(c *gin.Context) {
	if c.Request.Method == "GET" {
		render.Render(c, http.StatusOK, authViews.Registration())
	}

	if c.Request.Method == "POST" {
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid input",
			})
			return
		}
		defer file.Close()
		var profileImagePath string

		if file != nil {
			profileImagePath = ProfileImageUpload(c)
		}

		var body struct {
			Email    string
			Password string
			Role_ID  uint
		}

		if c.Bind(&body) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid input",
			})
			return
		}

		user := models.User{
			Email:          body.Email,
			Password:       body.Password,
			ProfileImgPath: profileImagePath,
			RoleID:         body.Role_ID,
		}

		result := connections.DB.Create(&user)

		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error while creating the user",
			})
			return
		}

		render.Redirect(c, "/login", http.StatusFound)
	}

}