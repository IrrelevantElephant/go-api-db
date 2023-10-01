package main

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type albumdb struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
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
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "unable to connect to database :("})
		return
	}

	defer dbpool.Close()

	rows, err := dbpool.Query(context.Background(), "SELECT album FROM albums;")

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "unable to query database :(", "error": err})
		return
	}
	defer rows.Close()

	var rowSlice []albumdb
	for rows.Next() {
		var r albumdb
		err := rows.Scan(&r)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		}
		rowSlice = append(rowSlice, r)
	}
	c.IndentedJSON(http.StatusOK, rowSlice)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "unable to connect to database :("})
		return
	}

	defer dbpool.Close()

	var album albumdb
	err = dbpool.QueryRow(context.Background(), "SELECT album FROM albums where id = $1;", id).Scan(&album)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "specified id not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func healthCheck(c *gin.Context) {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
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
