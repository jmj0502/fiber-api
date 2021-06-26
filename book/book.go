package book

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jmj0502/fiber-api/database"
	"gorm.io/gorm"
)

type Response struct {
	Success bool `json:success`
	Error bool 	`json:error`
}

type Book struct {
	gorm.Model
	Title string `json:title`
	Description string `json:description`
	Author string `json:author`
	Rating int8 `json:rating`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.Status(200).JSON(books)
}

func CreateBooks(c *fiber.Ctx) {
	book := new(Book)
	db := database.DBConn
	if err := c.BodyParser(book); err != nil {
		c.Status(503).JSON(Response{Success: false, Error: true})
		return
	}
	db.Create(&book)
	c.Status(201).JSON(book)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.Status(200).JSON(book)
}

func compareChanges(originalBook *Book, updatedBook *Book, field string) {
	switch field {
	case "author":
		if originalBook.Author != updatedBook.Author {
			originalBook.Author = updatedBook.Author
		}
	case "title":
		if originalBook.Title != updatedBook.Title {
			originalBook.Title = updatedBook.Title
		}
	case "description":
		if originalBook.Description != updatedBook.Description {
			originalBook.Description = updatedBook.Description
		}
	case "rating":
		if originalBook.Rating != updatedBook.Rating {
			originalBook.Rating = updatedBook.Rating
		}
	default:
		fmt.Println("Nothing to update.")
	}
}

func UpdateBook(c *fiber.Ctx) {
	id := c.Params("id") 
	db := database.DBConn 
	var updatedBook Book
	var book Book
	if err := c.BodyParser(&updatedBook); err != nil {
		c.Status(500).JSON(Response{Success: false, Error: true})
		return
	}
	db.Find(&book, id)
	var changesArr [4]string
	changesArr[0] = "title"
	changesArr[1] = "author"
	changesArr[2] = "description"
	changesArr[3] = "rating"
	for i := 0; i < len(changesArr); i++ {
		compareChanges(&book, &updatedBook, changesArr[i])
	}
	db.Save(&book)
	c.Status(202).JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	db.Delete(Book{}, id)
	c.Status(204).JSON(Response{Success: true, Error: false})
}