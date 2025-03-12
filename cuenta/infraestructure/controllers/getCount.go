package controllers

import (
	"log"
	"net/http"
	"rabbitMQ/cuenta/application"
	"rabbitMQ/cuenta/infraestructure"

	"github.com/gin-gonic/gin"
)

func GetCount(c *gin.Context) {
	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewGetCount(repo)

	count, err := useCase.Execute()
	if err != nil {
		log.Println("❌ Error ejecutando el caso de uso:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("✅ Datos obtenidos:", count)
	c.JSON(http.StatusOK, count)
}
