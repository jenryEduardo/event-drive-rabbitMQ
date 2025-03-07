package controllers

import (
	"net/http"
	"rabbitMQ/cuenta/application"
	"rabbitMQ/cuenta/domain"
	"rabbitMQ/cuenta/infraestructure"
	"strconv"

	"github.com/gin-gonic/gin"
)


func UpdateCount(c *gin.Context){
	idParam :=c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	
	var count domain.Cuenta


	if err := c.ShouldBindJSON(&count); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el JSON"})
		return
	}

	repo := infraestructure.NewMySQLRepository()
	useCase:=application.NewUpdate(repo)


	if err := useCase.Execute(id,count); err !=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"error al actualizar la cuenta"})
		return
	}

	c.JSON(http.StatusOK,gin.H{"ok":"dato actualizado correctamentee"})


}