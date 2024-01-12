package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
