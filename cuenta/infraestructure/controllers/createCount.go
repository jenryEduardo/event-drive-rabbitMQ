package controllers

import (
	"log"
	"net/http"
	"rabbitMQ/cuenta/application"
	"rabbitMQ/cuenta/domain"
	"rabbitMQ/cuenta/infraestructure"

	"github.com/gin-gonic/gin"
)
func CreateCount(c *gin.Context) {
    var cuenta domain.Cuenta

    // Intenta deserializar el JSON
    if err := c.ShouldBindJSON(&cuenta); err != nil {
        log.Println("❌ Error al procesar el JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el JSON", "details": err.Error()})
        return
    }

    log.Println("✅ Datos recibidos:", cuenta)

    // Crea el repositorio y el caso de uso
    repo := infraestructure.NewMySQLRepository()
    useCase := application.NewCreateCount(repo)

    if err := useCase.Execute(cuenta); err != nil {
        log.Println("❌ Error al guardar en la BD:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el usuario", "details": err.Error()})
        return
    }

    // Responde con éxito
    log.Println("✅ Cuenta creada con éxito")
    c.JSON(http.StatusOK, gin.H{"message": "Cuenta creada con éxito"})
}
