package routes


import (
	"github.com/gin-gonic/gin"
	"rabbitMQ/cuenta/infraestructure/controllers"
)


func SetupRoutesCount(router *gin.Engine) {

	routes:=router.Group("/cuenta")

	{
		routes.POST("/", controllers.CreateCount)
		routes.GET("/", controllers.GetCount)
		routes.PUT("/actualizar/:id", controllers.UpdateCount)
		routes.PUT("/:fromId/:toId",controllers.Transfering)
		routes.PUT("/deposito/:id",controllers.Deposit)
	}
}