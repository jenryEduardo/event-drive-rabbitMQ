package controllers

import (
	"net/http"
	"rabbitMQ/cuenta/application"
	"rabbitMQ/cuenta/infraestructure"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteCount(c *gin.Context) {
	
	idParam :=c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	repo:= infraestructure.NewMySQLRepository()
	useCase:=application.NewDeleteCount(repo)

	if err:= useCase.Execute(id);err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"no se pudo eliminar el registro"})
	}

	c.JSON(http.StatusOK,gin.H{"ok":"eliminado correctamente"})

}