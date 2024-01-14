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
	router.GET("/welcome", server.Welcome)
	router.GET("/lotofacil", server.Lotofacil)
	router.GET("/group", server.Group)
	router.GET("/quina", server.Quina)
	router.GET("/megasena", server.Megasena)
	router.GET("/login", server.Login)
	router.GET("/register", server.Register)

	router.Run("localhost:8080")
}
