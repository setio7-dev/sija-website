package routes

import (
	"sijaku-hebat/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Sijaku Hebat API",
		})
	})

	api := r.Group("/api")
	r.Static("/uploads", "./uploads")

	projectController := controllers.NewProjectController()
	itcController := controllers.NewItcController()
	moduleController := controllers.NewModuleController()

	{
		projects := api.Group("/project")
		itcs := api.Group("/itc")
		modules := api.Group("/module")

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
		}
	}

	return r
}
