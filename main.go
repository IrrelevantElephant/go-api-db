package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
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
	router := gin.Default()

	router.GET("/health", healthCheck)

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)

	router.Run("0.0.0.0:8080")
}

func getAlbums(c *gin.Context) {
	var rowSlice []album
	var dsn = getDatabaseUrl()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Find(&rowSlice)

	c.IndentedJSON(http.StatusOK, rowSlice)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	var dsn = getDatabaseUrl()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&album{})
	db.Create(&newAlbum)

	// protocol should be configurable
	var location = fmt.Sprintf("http://%s/%s/%d", c.Request.Host, "albums", newAlbum.Id)

	c.Header("location", location)
	c.IndentedJSON(http.StatusCreated, gin.H{"newid": newAlbum.Id})
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	var dsn = getDatabaseUrl()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var album album
	db.First(&album, id)

	c.IndentedJSON(http.StatusOK, album)
}

func healthCheck(c *gin.Context) {
	dbpool, err := pgxpool.New(context.Background(), getDatabaseUrl())
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "unable to connect to database :("})
		return
	}

	defer dbpool.Close()

	var message string
	err = dbpool.QueryRow(context.Background(), "SELECT 'healthy';").Scan(&message)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "unable to query database :(", "error": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "healthy :)"})
}

func getDatabaseUrl() string {
	return os.Getenv("DATABASE_URL")
}
