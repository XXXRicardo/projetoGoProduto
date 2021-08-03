package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		mercado := main.Group("mercado")
		{

			mercado.GET("/:id", controllers.ShowProduto)
			mercado.GET("/", controllers.ShowProduto)
			mercado.POST("/", controllers.CreateProduto)
			mercado.PUT("/", controllers.UpdateProduto)
			mercado.DELETE("/:id", controllers.DeleteMercado)
		}
	}
	return router
}
