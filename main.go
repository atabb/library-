package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Capital names for variables to make it exportable field in go, in our case we need to export as JSON respone so we also added json fields at the end.

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

//typically use DB for sites, but for this tutorial use in memory database we set below.
// Slice of book structs define below
var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
	{ID: "4", Title: "Mimetic Theory", Author: "Abdul Tabbakha", Quantity: 3},
}

//Get request function to get json contents of all values stored in books
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books) //returns json formatted books and http staus of OK.

}

//POST request to add to books slice.
func addBooks(c *gin.Context) {
	var addBooks book //create struct called addBooks from book struct

	if err := c.BindJSON(&addBooks); err != nil { //checks for any error before proceeding to adding a book to book struct
		return
	}

	books = append(books, addBooks)              //apending to the slice of books above.
	c.IndentedJSON(http.StatusCreated, addBooks) //return book we created with status created to let us know it worked!
}

func main() {

	router := gin.Default()        //create router gin
	router.GET("/books", getBooks) //if you visit /books site, it will return function getbooks define above.
	router.POST("/books", addBooks)
	router.Run("localhost:8080")

}
