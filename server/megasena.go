package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Megasena(c *gin.Context) {

	c.HTML(http.StatusOK, "megasena.html", gin.H{
		"title": "Benvindo",
	})
}
