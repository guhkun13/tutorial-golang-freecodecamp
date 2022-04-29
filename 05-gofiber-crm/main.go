package main

import (
	"fmt"
	"log"
	
	"github.com/jinzhu/gorm"
	"github.com/gofiber/fiber"
	"github.com/guhkun13/tutorial/freeCodeCamp/05-gofiber-crm/database"
	"github.com/guhkun13/tutorial/freeCodeCamp/05-gofiber-crm/lead"
)


func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead/", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead/", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase(){
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	
	if err != nil {
		panic("failed to open database")
	}
	fmt.Println("connected to database")
	
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	
	initDatabase()
	setupRoutes(app)
	
	app.Listen(8010)
	defer database.DBConn.Close()
	
}