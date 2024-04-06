package main

import (
	"fmt"
	"log"
	"crm-go-lang/lead"
	"crm-go-lang/database"
	"github.com/gofiber/fiber"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.AddLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	defer database.DBConn.Close() // Close database connection when main exits

	// Start the server and log any error
	if err := app.Listen(3000); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
