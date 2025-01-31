package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{

	{ID: "1", Title: " bangloreboy", Artist: "srujan", Price: 100.0},
	{ID: "2", Title: " shivamogaboy", Artist: "Sathwik", Price: 60.0},
	{ID: "3", Title: " kdboy", Artist: "Kd", Price: 10.0},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)

}

func PostAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)

}

func getAlbumID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {

	router := gin.Default()
	router.GET("/album", GetAlbums)
	router.POST("/album", PostAlbums)
	router.POST("/album/:id", getAlbumID)

	router.Run("localhost:8080")

}
