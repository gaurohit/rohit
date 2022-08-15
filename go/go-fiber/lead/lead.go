package lead

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/rohit/go-fiber/database"
)

type Lead struct{
	gorm.Model
	Name 	string `json:"name"`
	Company string `json:"company"`
	Email 	string	`json:"email"`
	Phone 	int		`json:"phone"`
}

func GetLeads(c *fiber.Ctx)  {
	db := database.DBconn
	var lead []Lead
    db.Find(&lead)
	c.JSON(lead)
}

func GetLead(c *fiber.Ctx) {
   id := c.Params("id")
   db := database.DBconn
   var lead Lead
   db.Find(&lead,id)
   c.JSON(lead)

}

func NewLead(c *fiber.Ctx)  {
	db := database.DBconn

	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil{
		c.Status(503).Send(err)
	}
	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx)  {
	id := c.Params("id")
	db := database.DBconn

	var lead Lead
	db.First(&lead,id)
	if lead.Name == ""{
		c.Status(503).Send("No lead found with given id")
	}
    db.Delete(&lead)
	c.JSON(lead)
}