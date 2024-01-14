package main

import (
	"github.com/gin-gonic/gin"
	"github.com/heronhurpia/loterias/server"
)

func main() {

	// Lista de rotas
	router := gin.Default()
	router.Static("/templates", "./templates")
	router.LoadHTMLGlob("templates/views/*")
	router.GET("/", server.Welcome)
	router.Run("localhost:8080")
}
