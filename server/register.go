package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Novo usuário",
	})
}
