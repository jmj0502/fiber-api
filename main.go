package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber"
	"github.com/jmj0502/fiber-api/book"
	"github.com/jmj0502/fiber-api/database"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//private function
func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World")
}

func getEnvVars(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Unable to load the .env file!")
	}
	return os.Getenv(key)
}

func dbConnection() {
	var err error
	var dsn string = getEnvVars("DB_USERNAME") + ":" + getEnvVars("DB_PASSWORD") + "@tcp(127.0.0.1:3306)/fiber?charset=utf8mb4&parseTime=True"
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Error while conecting to the database")
	}
	fmt.Println("Connection stablished successfully!")
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Migration performed successfully!")
}

//public function
func setRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/api/books", book.GetBooks)
	app.Post("/api/books", book.CreateBooks)
	app.Put("/api/books/:id", book.UpdateBook)
	app.Get("/api/books/:id", book.GetBook)
	app.Delete("/api/books/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()
	dbConnection()
	//running our first route.
	setRoutes(app)

	app.Listen(4000)
}