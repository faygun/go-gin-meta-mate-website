package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

func registerRoutes() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	r.Any("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/contact", func(c *gin.Context) {

		email, _ := c.GetPostForm("email")
		username, _ := c.GetPostForm("username")
		message, _ := c.GetPostForm("message")

		date := time.Now().String()

		info := fmt.Sprintf("date: %s \nusername: %s \nemail: %s \nmessage: %s\n\n", date, username, email, message)

		writeText(info)
		c.JSON(http.StatusOK, "Your message has been successfully sent.")
	})

	r.Static("/public", "./public")

	return r
}

func writeText(s string) {
	base := string(http.Dir("public"))

	path := path.Join(base, "file", "email.txt")
	// file, err := os.Stat(path)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	f, err := os.OpenFile(path, os.O_APPEND, 0644)
	// if file != nil {
	// 	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	// }
	// if file == nil {
	// 	f, err := os.Create(path)
	// }

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(s)

	if err2 != nil {
		log.Fatal(err2)
	}
}
