package routes

import (
	"ticketing/backend-api/controllers"
	"ticketing/backend-api/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	api := router.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		admin := api.Group("", middlewares.AuthMiddleware(), middlewares.SuperAdminOnly())
		{
			users := admin.Group("/users")
			{
				users.GET("", controllers.FindUser)
				users.POST("", controllers.CreateUser)
				users.GET("/:id", controllers.FindUserById)
				users.PUT("/:id", controllers.UpdateUser)
				users.DELETE("/:id", controllers.DeleteUser)
			}

			branches := admin.Group("/branches")
			{
				branches.GET("", controllers.FindBranch)
				branches.POST("", controllers.CreateBranch)
				branches.GET("/:id", controllers.FindBranchById)
				branches.PUT("/:id", controllers.UpdateBranch)
				branches.DELETE("/:id", controllers.DeleteBranch)
			}

			divisions := admin.Group("/divisions")
			{
				divisions.GET("", controllers.FindDivision)
				divisions.POST("", controllers.CreateDivision)
				divisions.GET("/:id", controllers.FindDivisionByid)
				divisions.PUT("/:id", controllers.UpdateDivision)
				divisions.DELETE("/:id", controllers.DeleteDivision)
			}

			departments := admin.Group("/departments")
			{
				departments.GET("", controllers.FindDepartment)
				departments.POST("", controllers.CreateDepartment)
				departments.GET("/:id", controllers.FindDepartmentById)
				departments.PUT("/:id", controllers.UpdateDepartment)
				departments.DELETE("/:id", controllers.DeleteDepartment)
			}

			categories := admin.Group("/categories")
			{
				categories.GET("", controllers.FindCategory)
				categories.POST("", controllers.CreateCategory)
				categories.GET("/:id", controllers.FindCategoryById)
				categories.PUT("/:id", controllers.UpdateCategory)
				categories.DELETE("/:id", controllers.DeleteCategory)
			}
		}
	}
	return router
}
