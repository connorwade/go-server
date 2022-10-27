package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	db database
)

func main() {
	//setup db
	db.Migrate()

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
	// var db database
	// db.Migrate()
	// fmt.Println(db.Data)
	// newAlbum := Music{
	// 	ID:     "4",
	// 	Title:  "Ray",
	// 	Artist: "L'arc-en-ciel",
	// 	Price:  12.99,
	// }
	// db.Create(newAlbum)
	// fmt.Println(db.Data)
	// s, _ := db.All()
	// fmt.Println(s)
	// rMusic, _ := db.GetByID("4")
	// fmt.Println(*rMusic)
	// upMusic, _ := db.Update("4", Music{
	// 	ID:     "4",
	// 	Title:  "Ark",
	// 	Artist: "L'arc-en-ciel",
	// 	Price:  13.99,
	// })
	// fmt.Println(*upMusic)
	// s, _ = db.All()
	// fmt.Println(s)
	// db.Delete("4")
	// s, _ = db.All()
	// fmt.Println(s)
}

func getAlbums(c *gin.Context) {
	c.Bind(&db.Data)
	c.JSON(http.StatusOK, gin.H{
		"data": db.Data.Albums,
	})
}

func postAlbums(c *gin.Context) {
	var newAlbum Music

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	err := db.Create(newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad bad bad"})
	}
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	music, err := db.GetByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
	c.Bind(music)
	c.IndentedJSON(http.StatusOK, gin.H{"music": music})
}
