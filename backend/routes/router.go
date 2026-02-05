package routes

import (
	"sijaku-hebat/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Sija Website API",
		})
	})

	api := r.Group("/api")
	r.Static("/uploads", "./uploads")

	projectController := controllers.NewProjectController()
	itcController := controllers.NewItcController()
	moduleController := controllers.NewModuleController()
	companyController := controllers.NewCompanyController()
	userController := controllers.NewUserController()

	{
		projects := api.Group("/project")
		itcs := api.Group("/itc")
		modules := api.Group("/module")
		companies := api.Group("/company")
		users := api.Group("/user")

		{
			projects.GET("", projectController.GetAll)
			projects.GET("/:id", projectController.GetById)
			projects.POST("", projectController.Create)
			projects.PUT("/:id", projectController.Update)
			projects.DELETE("/:id", projectController.Delete)

			itcs.GET("", itcController.GetAll)
			itcs.GET("/:id", itcController.GetById)
			itcs.POST("", itcController.Create)
			itcs.PUT("/:id", itcController.Update)
			itcs.DELETE("/:id", itcController.Delete)

			modules.GET("", moduleController.GetAll)
			modules.GET("/:id", moduleController.GetById)
			modules.POST("", moduleController.Create)
			modules.PUT("/:id", moduleController.Update)
			modules.DELETE("/:id", moduleController.Delete)

			companies.GET("", companyController.GetAll)
			companies.GET("/:id", companyController.GetById)
			companies.POST("", companyController.Create)
			companies.PUT("/:id", companyController.Update)
			companies.DELETE("/:id", companyController.Delete)

			users.GET("", userController.GetAll)
			users.GET("/:id", userController.GetById)
			users.POST("/login", userController.Login)
			users.POST("", userController.Create)
			users.PUT("/:id", userController.Update)
			users.DELETE("/:id", userController.Delete)
		}
	}

	return r
}
