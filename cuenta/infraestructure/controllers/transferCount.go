package controllers

import (
	"fmt"
	"net/http"
	"rabbitMQ/cuenta/application"
	"rabbitMQ/cuenta/infraestructure"
	"rabbitMQ/cuenta/infraestructure/adaptadores"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TransferRequest struct {
	Monto float64 `json:"monto"`
}
func Transfering(c *gin.Context) {
    fmt.Println("Solicitud recibida en el backend")

    fromIdP := c.Param("fromId")
    fromId, err := strconv.Atoi(fromIdP)
    if err != nil {
        fmt.Println("Error en fromId:", err) // 🛠 Log de error
        c.JSON(http.StatusBadRequest, gin.H{"error": "fromId debe ser un número"})
        return
    }

    toIdP := c.Param("toId")
    toId, err := strconv.Atoi(toIdP)
    if err != nil {
        fmt.Println("Error en toId:", err) // 🛠 Log de error
        c.JSON(http.StatusBadRequest, gin.H{"error": "toId debe ser un número"})
        return
    }

	fmt.Println(toId)

    var req TransferRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        fmt.Println("Error al procesar JSON:", err) // 🛠 Log de error
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar JSON"})
        return
    }

    fmt.Println("fromId:", fromId, "toId:", toId, "Monto:", req.Monto) 

    // Publicar la transacción en RabbitMQ
    transaction := adaptadores.Transaction{
        ID:        fmt.Sprintf("%d-%d-%d", fromId, toId, time.Now().UnixNano()),
        From:      fromId,
        To:        toId,
        Amount:    req.Monto,
        Timestamp: time.Now().Format(time.RFC3339),
    }

    success, err := adaptadores.PublishTransaction(transaction)
    if err != nil || !success {
        fmt.Println("Error en la publicación de la transacción:", err) // 🛠 Log de error
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error en la transferencia"})
        return
    }

    repo := infraestructure.NewMySQLRepository()
    useCase := application.NewTransfer(repo)

    if err := useCase.Execute(fromId, toId, req.Monto); err != nil {
        fmt.Println("Error ejecutando la transferencia:", err) // 🛠 Log de error
        c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo mandar los datos"})
        return
    }

    fmt.Println("Transferencia realizada con éxito") // ✅ Confirmación en backend
    c.JSON(http.StatusOK, gin.H{"message": "Transferencia realizada con éxito"})
}
