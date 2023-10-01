package main

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type album struct {
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

	var rowSlice []album
	for rows.Next() {
		var r album
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

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "unable to connect to database :("})
		return
	}

	defer dbpool.Close()

	// TODO: get it to actually return the id
	id, err := dbpool.Exec(context.Background(), "INSERT INTO albums (album) VALUES ($1) RETURNING id;", newAlbum)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"newid": id})
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "unable to connect to database :("})
		return
	}

	defer dbpool.Close()

	var album album
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
