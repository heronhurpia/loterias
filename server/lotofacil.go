package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Lotofacil(c *gin.Context) {

	c.HTML(http.StatusOK, "lotofacil.html", gin.H{
		"title": "Benvindo lotofacil",
	})
}
