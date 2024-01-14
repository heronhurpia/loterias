package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Group(c *gin.Context) {

	c.HTML(http.StatusOK, "group.html", gin.H{
		"title": "Benvindo grupo",
	})
}
