package main

import (
	"fmt"
	"github.com/rohit/go-fiber/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/rohit/go-fiber/database"
)

func setuproutes(app *fiber.App)  {
	app.Get("/api/v1/lead",lead.GetLeads)
	app.Post("/api/v1/lead",lead.NewLead)
	app.Delete("/api/v1/lead/:id",lead.DeleteLead)
	app.Get("/api/v1/lead/:id",lead.GetLead)

}

func initDatabase()  {
	var err error
	database.DBconn,err = gorm.Open("sqlite3","lead.db")
	if err != nil{
		panic("failed to connect database")
	}
	fmt.Println("connection opened to database")
	database.DBconn.AutoMigrate(&lead.Lead{})
	fmt.Println("database Migrated")
}

func main()  {
	app := fiber.New()
	initDatabase()
    setuproutes(app)
	app.Listen(3000)
	defer database.DBconn.Close()
}