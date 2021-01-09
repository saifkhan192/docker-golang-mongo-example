package models

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
)

type User struct {
	gorm.Model
	Name  string `json:"name",omitempty;gorm:"default:'name'"`
	Email string `json:"email",omitempty;gorm:"default:'email'"`
}

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dbInst, err := gorm.Open("sqlite3", "./sqlite.db")
	log.Print(dbInst)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	log.Info("SQLite DB opened!")
	dbInst.AutoMigrate(&User{})
	DB = dbInst
	// defer db.Close()
	return dbInst
}

func GetUsers(c *gin.Context) {
	var users []User
	DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	DB.Create(&User{Name: "Saif", Email: "saif@gmail.com"})
}
