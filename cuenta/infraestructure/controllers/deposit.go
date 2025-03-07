package controllers

import (
	"net/http"
	"rabbitMQ/cuenta/application"
	"rabbitMQ/cuenta/infraestructure"
	"strconv"

	"github.com/gin-gonic/gin"
)


type Request struct{
	Saldo float64 `json:"saldo"`
}

func Deposit(c *gin.Context){


	id:=c.Param("id")

	idInt,err := strconv.Atoi(id)

	if err != nil{
		c.JSON(http.StatusForbidden,gin.H{"error":"no se oudo mandar"})
	}

	var req TransferRequest

	// Deserializar el JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el JSON", "details": err.Error()})
		return
	}

	monto:=req.Monto

	repo:=infraestructure.NewMySQLRepository()
	useCase:=application.NewDeposit(repo)

	err = useCase.Execute(idInt,monto)

	if err != nil{
		c.JSON(http.StatusGatewayTimeout,"verifique sus datos")
		return
	}

	c.JSON(http.StatusCreated,gin.H{"ok":"se realizo la transaccion"})

}