package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type book struct {
	ID string  `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
}

var books = []book{
  {ID: "1", Title: "Search of Lost Time", Author: "Marcel Proust", Quantity:2},
  {ID: "2", Title: "Great Gatsby", Author: "Scott Fitzgerald", Quantity:5},
  {ID: "3", Title: "Atomic Habits", Author: "James Clear", Quantity:9},
}
// get Request
func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}

// func createBook(c *gin.Context){
// 	var newBook book
// 	if err := c.BindJSON(&newBook); err != nil {
// 		return
// 	}
// }

func main(){
	router := gin.Default()
    router.GET("/books" , getBooks)
	router.Run("localhost:8080")
	
}