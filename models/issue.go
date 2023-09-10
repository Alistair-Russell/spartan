package models

import (
	"gitlab.com/alistairr/spartan/db"

    "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Issue struct {
	gorm.Model
	Title       string
	Description string
	Status      string
}

func GetIssues(c *fiber.Ctx) error {
	database := db.DBConn
	var issues []Issue
	database.Find(&issues)
	return c.JSON(issues)
}

func GetIssue(c *fiber.Ctx) error {
    id := c.Params("id")
	database := db.DBConn
	var issue Issue
	database.Find(&issue, id)
	return c.JSON(issue)
}

func CreateIssue(c *fiber.Ctx) error {
	database := db.DBConn
    issue := new(Issue)
    if err := c.BodyParser(issue); err != nil {
        return c.Status(503).SendString(err.Error())
    }
	database.Create(&issue)
	return c.JSON(issue)
}

func DeleteIssue(c *fiber.Ctx) error {
    id := c.Params("id")
	database := db.DBConn
	var issue Issue
	database.First(&issue, id)
    if issue.Title == "" {
        return c.Status(500).SendString("No issue found with ID")
    }
    database.Delete(&issue)
    return c.SendString("Issue successfully deleted")
}

func Migrate() {
	db.DBConn.AutoMigrate(&Issue{})
}

