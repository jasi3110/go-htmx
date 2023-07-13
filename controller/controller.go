package controller

import (
	"go-htmx/domain"
	"html/template"
	"log"
	// "net/http"
	"github.com/gin-gonic/gin"
)

func Hometemp(c *gin.Context) {
	tmp := template.Must(template.ParseFiles("C:/Users/JASIM MANSOOR/go/src/go-htmx/index.html")) // Update the file path accordingly
ucl := map[string][]domain.UCL{
		"ucl": {
			{Club_Name: "REAL MADRID", Players: "32"},
			{Club_Name: "CHELSEA", Players: "29"},
			{Club_Name: "AJAX", Players: "35"},
		},
		
	}

	tmp.Execute(c.Writer, ucl)
}

func Addtemp(c *gin.Context) {
	log.Println(c.GetHeader("HX-Request"))

	Club_Name := c.PostForm("Club_Name")
	Players := c.PostForm("Players")
	log.Println(Club_Name, Players)
}