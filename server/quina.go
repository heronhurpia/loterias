package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Quina(c *gin.Context) {

	c.HTML(http.StatusOK, "quina.html", gin.H{
		"title": "Benvindo",
	})
}
