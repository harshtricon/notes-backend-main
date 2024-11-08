package main

import (
	"log"

	"github.com/harshtricon/notes-backend-main/database"
	"github.com/harshtricon/notes-backend-main/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	routes.Setup(app)
	log.Println("Server is running at http://localhost:8000")
	log.Fatal(app.Listen(":8000"))
}
