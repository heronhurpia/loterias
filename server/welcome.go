package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {

	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"title": "Benvindo",
	})
}
