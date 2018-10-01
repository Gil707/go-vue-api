package main

import (
	"./common"
	"./models"
	"github.com/gin-gonic/gin"
)

const API = "api/v1/"

func main() {

	r := gin.Default()
	r.Use(common.LiberalCORS)
	r.GET(API+"/ping", pong)
	r.GET(API+"/users", models.GetUsers)
	r.GET(API+"/users/:id", models.GetUser)
	r.GET(API+"/products", models.GetProducts)
	r.GET(API+"/products/:p_id", models.GetProduct)
	r.Run(":9090")
}

/* Simple test api request */

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
