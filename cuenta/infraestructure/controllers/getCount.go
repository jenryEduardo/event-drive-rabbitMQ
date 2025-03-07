package controllers

import (
	"net/http"
	"rabbitMQ/cuenta/application"
	"rabbitMQ/cuenta/infraestructure"

	"github.com/gin-gonic/gin"
)

func GetCount(c *gin.Context) {

	repo:= infraestructure.NewMySQLRepository()
	useCase:=application.NewGetCount(repo)

	count,_:= useCase.Execute()


	c.JSON(http.StatusOK,count)


}