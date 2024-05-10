package main

import (
				"net/http"
				"github.com/gin-gonic/gin"
				"errors"
				"fmt"
)

type book struct {
				ID	string `json:"id"`
				Title	string `json:"title"`
				Author	string `json:"author"`
				Quantity	int `json:"quantity"`
}


var books = []book{
				{ID: "1", Title: "Aadarsh Jha's book1", Author: "AJ", Quantity : 2},
				{ID: "2", Title:"Platinum J Book 1", Author: "PlatinumJ", Quantity: 5},
}


func getBooks(c *gin.Context){
				c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
				id:=c.Param("id")
				book, err := getBookById(id)

				if err != nil {
								c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Book not found"})
								return 
				}
				c.IndentedJSON(http.StatusOK, book)

}

func getBookById(id string) (*book, error) {
				for i,b:= range books {
								if(b.ID == id){
												return &books[i], nil
								}
				}

				return nil, errors.New("Book Not Found")
}

func createBook(c *gin.Context){
				//c has all query data and body
				var newBook book 
				if err := c.BindJSON(&newBook); err !=nil {return}
				books = append(books,newBook)
				c.IndentedJSON(http.StatusCreated, newBook)
				
}

func checkoutBook(c *gin.Context){
				//Using a query Param
				fmt.Println("Checkout Book")
				id, ok := c.GetQuery("id")
				if ok == false {
								c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing Query Parameter"})
								return 
				}

				book, err := getBookById(id)

				if err != nil {
								c.IndentedJSON(http.StatusNotFound, gin.H{"message" :"Book was not found"})
								return 
				}
				if book.Quantity <= 0 {
								c.IndentedJSON(http.StatusBadRequest, gin.H{"message" :"Book unavailable"})
								return 
				}

				book.Quantity -=1
				c.IndentedJSON(http.StatusOK, book)
				//DONE
}

func returnBook(c *gin.Context){
				id, ok := c.GetQuery("id")

				fmt.Println("Returning a book", id)
					if ok == false {
								c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing Query Parameter"})
								return 
				}

				book, err := getBookById(id)

				if err != nil {
								c.IndentedJSON(http.StatusNotFound, gin.H{"message" :"Book was not found"})
								return 
				}
				book.Quantity +=1
				c.IndentedJSON(http.StatusOK, book)
				

}

func main() {
				router := gin.Default();
				router.GET("/books", getBooks)
				router.POST("/books", createBook)
				router.GET("/book/:id", bookById)
				router.PATCH("/checkout", checkoutBook)
				router.PATCH("/return", returnBook)
				router.Run("localhost:8081")

}
