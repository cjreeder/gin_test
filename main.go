package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cjreeder/gin_test/handler"
	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func async() {
	time.Sleep(1 * time.Minute)
	fmt.Println("I've been waiting for 1 Minute for this to return")
}

func middleware(c *gin.Context) {
	fmt.Printf("We are in the middleware layer\n")
	c.Next()
}

/*
func middleware(c *gin.Context) {
	fmt.Printf("Middleware Layer 2: %s\n")
	c.Next()
}
*/
// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	fmt.Printf("Starting Test Service.......\n")
	router := gin.Default()
	router.Use(middleware)
	var hello string
	hello = fmt.Sprint("Hello World")
	//router.Use(middlware2)
	router.GET("/albums", getAlbums)
	router.POST("/async", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Message": "OK"})
		go async()
		fmt.Printf("%s, I'm not waiting for you.......\n", hello)
	})
	router.POST("/handlertest", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Message": "Testing Handler from another package"})
		go handler.AsyncHandler()
		fmt.Printf("I am not waiting for the handler test......\n")
	})

	router.Run("localhost:8080")
}
