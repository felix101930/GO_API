package main

import (

	// "github.com/bytedance/sonic"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)


type book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
}

var books = []book{


	{
		ID: "1",
		Title: "The Catcher in the Rye",
		Author: "J.D. Salinger",
		Quantity: 10,
	},
	{

		ID: "2",
		Title: "To Kill a Mockingbird",
		Author: "Harper Lee",
		Quantity: 5,
	},
	{
			ID: "3",
			Title: "1984",
			Author: "George Orwell",
			Quantity: 3,
	},

	
}


func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)


}

func checkoutBook(c *gin.Context) {

	id,ok := c.GetQuery("id")
	if !ok{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"Missing Book"})
		return
	}

	book, err:= getBookById(id)

	if err!=nil {
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"Missing Book"})
		return

	}

	if book.Quantity <= 0{
		c.IndentedJSON(http.StatusBadRequest,gin.H{"message":"Book out of stock"})
		return
	}

	book.Quantity =-1
	c.IndentedJSON(http.StatusOK,book)

}

func returnBook (c *gin.Context){
	// obtain id
	id, ok := c.GetQuery("id")

	// if id not found
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "Missing id"})	
		return

	}

	// obtain book through id
	book, err := getBookById(id)

	// if err exists
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return	
	}

	//add 1 to book count 
	book.Quantity +=1 

	c.IndentedJSON(http.StatusOK, book)	
}

func bookById(c *gin.Context){
	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}


func getBookById(id string ) (*book, error) {

	for i, b := range books {
		if b.ID == id{
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")

}

func createBook(c *gin.Context) {

	var newBook book

	if err := c.BindJSON(&newBook); err != nil{
		return	
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}


func main(){
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")


}

