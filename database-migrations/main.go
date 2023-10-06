package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type album struct {
	Id     uint    `gorm:"primaryKey"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	var dsn = getDatabaseUrl()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&album{})

	router := gin.Default()
	router.GET("/health", func(c *gin.Context) { c.IndentedJSON(http.StatusOK, gin.H{"message": "healthy :)"}) })
	router.Run("0.0.0.0:8080")
}

func getDatabaseUrl() string {
	return os.Getenv("DATABASE_URL")
}
