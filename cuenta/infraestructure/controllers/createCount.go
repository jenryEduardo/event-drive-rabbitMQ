package controllers

import (
	"rabbitMQ/cuenta/application"
	"rabbitMQ/cuenta/domain"
	"rabbitMQ/cuenta/infraestructure"
	"net/http"
	"github.com/gin-gonic/gin"
)

func CreateCount(c *gin.Context) {
    var cuenta domain.Cuenta

    // Intenta deserializar el JSON
    if err := c.ShouldBindJSON(&cuenta); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el JSON", "details": err.Error()})
        return
    }

    // Continua con la lógica si no hubo errores
    repo := infraestructure.NewMySQLRepository()
    useCase := application.NewCreateCount(repo)

    if err := useCase.Execute(cuenta); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el usuario"})
        return
    }

    // Responde con éxito
    c.JSON(http.StatusOK, gin.H{"message": "Cuenta creada con éxito"})
}
