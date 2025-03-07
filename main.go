package main

import (
	"log"
	"time"
	countsRoutes "rabbitMQ/cuenta/infraestructure/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default()

	// Configurar CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000","http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	countsRoutes.SetupRoutesCount(router)
	

	port := ":8080"
	log.Println("Servidor escuchando en el puerto", port)
	log.Fatal(router.Run(port))

		
}