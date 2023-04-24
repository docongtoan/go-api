package router

import (
	controllers "goserverapi/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	router := gin.Default()

	router.Static("/public", "./public")

	public := router.Group("/api")
	{
		public.POST("/login", controllers.Login)
		public.POST("/register")
	}

	protected := router.Group("/api/admin")
	{
		// unit
		protected.GET("/unit/get-list", controllers.GetListUnit)
		protected.GET("/unit/get-row/:id", controllers.GetRowUnit)
		protected.POST("/unit/insert", controllers.InsertUnit)
		protected.POST("/unit/update/:id", controllers.UpdateUnit)
		protected.DELETE("/unit/delete/:id", controllers.DeleteUnit)
	}

	return router
}
