package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// "html/template"
	models "app/models"
	"fmt"
	"log"
	"time"
)

func somefunction(from string) {

	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
	// time.Sleep(6)

}

func main() {
	// init connection
	models.ConnectDB()

	// set routes
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	router.GET("/users", models.GetUsers)
	router.GET("/users/create", models.CreateUser)
	router.GET("/getFile", func(c *gin.Context) {
		c.File("./main.go")
	})
	router.GET("/goroutine", func(c *gin.Context) {
		go somefunction("goroutine")
		go somefunction("going")
		somefunction("inline")
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	// run application
	router.Run() // listen and serve on 0.0.0.0:8080
	log.Print("open: http://localhost:3000/users")
	log.Print("open: http://localhost:3000/users/create")
	log.Print("open: http://localhost:3000/getFile")
}
