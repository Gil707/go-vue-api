package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

type Product struct {
	P_id           int       `json:"p_id"`
	Z_id           int       `json:"z_id"`
	Name           string    `json:"name"`
	Basket         int       `json:"basket"`
	Img            string    `json:"img"`
	Link           string    `json:"link"`
	Count          int       `json:"count"`
	Comment        string    `json:"comment"`
	Price          string    `json:"price"`
	Price_t        string    `json:"price_t"`
	Delivery_price string    `json:"delivery_price"`
	Editor_comm    string    `json:"editor_comm"`
	Zamena         string    `json:"zamena"`
	Adate          time.Time `json:"adate"`
	Status         int       `json:"status"`
	Answer         string    `json:"answer"`
	Delivery_cn    float32   `json:"delivery_cn"`
	Pushed         int       `json:"pushed"`
	Cn_color       int       `json:"cn_color"`
	Updated_at     time.Time `json:"updated_at" type:"datetime"`
}

func GetProducts(c *gin.Context) {

	db, err = gorm.Open("sqlite3", "./data/godb.db")

	if err != nil {
		panic("Could not connect to DB")
	}
	defer db.Close()

	var products []Product

	db.Find(&products)

	c.JSON(200, gin.H{
		"products": products,
	})
}

func GetProduct(c *gin.Context) {

	db, err = gorm.Open("sqlite3", "./data/godb.db")

	if err != nil {
		panic("Could not connect to DB")
	}

	defer db.Close()

	var product Product

	db.Where("p_id = ?", c.Param("p_id")).First(&product)

	c.JSON(200, gin.H{
		"product": product,
	})
}
