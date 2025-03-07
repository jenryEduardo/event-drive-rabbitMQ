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

	fromIdP := c.Param("fromId")
	fromId, err := strconv.Atoi(fromIdP)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "fromId debe ser un número"})
		return
	}

	toIdP := c.Param("toId")
	toId, err := strconv.Atoi(toIdP)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "toId debe ser un número"})
		return
	}

	var req TransferRequest


	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar JSON"})
		return
	}

	// crear la transacción y publicarla en RabbitMQ
	transaction := adaptadores.Transaction{
		ID:        fmt.Sprintf("%d-%d-%d", fromId, toId, time.Now().UnixNano()),
		From:      fromId,
		To:        toId,
		Amount:    req.Monto,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	success, err := adaptadores.PublishTransaction(transaction)
	if err != nil || !success {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error en la transferencia"})
		return
	}

	repo:=infraestructure.NewMySQLRepository()
	useCase:=application.NewTransfer(repo)
	fmt.Println("fromId:", fromId, "toId:", toId, "Monto:", req.Monto)

	if err:=useCase.Execute(fromId,toId,req.Monto);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"no se pudo mandar los datos"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transferencia realizada con éxito"})
}