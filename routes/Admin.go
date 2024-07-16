package routes

import (
	"github.com/Hrishikesh-Panigrahi/GoCMS/controllers"
	"github.com/gin-gonic/gin"
)

func adminRoutes(superRoute *gin.RouterGroup) {
	AdminRoutes := superRoute.Group("api/v1/admin")
	{
		adminPostRouts := AdminRoutes.Group("/post")
		{
			adminPostRouts.POST("/", controllers.AdminCreatePost)
			adminPostRouts.PUT("/:id", controllers.AdminUpdatePost)
			adminPostRouts.DELETE("/:id", controllers.AdminDeletePost)
		}

		adminUserRouts := AdminRoutes.Group("/user")
		{
			adminUserRouts.POST("/create", controllers.CreateUser)
			adminUserRouts.PUT("/:id", controllers.UpdateUser)
			adminUserRouts.DELETE("/:id", controllers.DeleteUser)
		}

		
	}

	AdminFormRoutes := superRoute.Group("/admin")
	{
		adminUserGETRouts := AdminFormRoutes.Group("/user")
		{
			adminUserGETRouts.GET("/", controllers.GetUsers)
			adminUserGETRouts.GET("/:id", controllers.GetUser)
			adminUserGETRouts.GET("/create", controllers.CreateUser)
			adminUserGETRouts.GET("/update/:id", controllers.UpdateUser)
		}	

		adminPostGETRouts := AdminFormRoutes.Group("/post")
		{
			adminPostGETRouts.GET("/", controllers.AdminGetPosts)
			adminPostGETRouts.GET("/create", controllers.AdminCreatePost)
			adminPostGETRouts.GET("/update/:id", controllers.AdminUpdatePost)
			adminPostGETRouts.GET("delete/:id", controllers.AdminDeletePost)
		}	
	}
}
