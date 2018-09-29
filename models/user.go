package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type User struct {
	Id      int    `json:"id"`
	Login   string `json:"login"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
	Type    string `json:"type"`
	Car     string `json:"car"`
	Status  int    `json:"status"`
}

func GetUsers(c *gin.Context) {

	db, err = gorm.Open("sqlite3", "./data/godb.db")

	if err != nil {
		panic("Could not connect to DB")
	}
	defer db.Close()

	var users []User

	db.Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})
}

func GetUser(c *gin.Context) {

	db, err = gorm.Open("sqlite3", "./data/godb.db")

	if err != nil {
		panic("Could not connect to DB")
	}

	defer db.Close()

	var user User

	db.First(&user, c.Param("id"))

	c.JSON(200, gin.H{
		"user": user,
	})
}
