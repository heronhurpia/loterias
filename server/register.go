package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Novo usuário",
	})
}

func SubmitForm(c *gin.Context) {

	name := c.PostForm("name")
	log.Println("name = ", name)

	email := c.PostForm("email")
	log.Println("email = ", email)

	// Return the message as HTML
	response := fmt.Sprintf("Form submitted successfully! Name: %s, Email: %s", name, email)
	c.HTML(http.StatusOK, "response.html", gin.H{
		"message": response,
	})
}
