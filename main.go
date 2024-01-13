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
	router.GET("/", welcome)
	router.Run("localhost:8080")
}

func welcome(c *gin.Context) {
	server.Welcome(c)
}
