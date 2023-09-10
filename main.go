// main.go
package main

import (
	"gitlab.com/alistairr/spartan/models"
    "gitlab.com/alistairr/spartan/db"
    "fmt"
	"log"

	"github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
    app.Get("/", func(c *fiber.Ctx) error {
        return c.Render("index", fiber.Map{"Issues": models.GetIssues})
    })
    app.Get("/issues", func(c *fiber.Ctx) error {
        return c.Render("issue-list", fiber.Map{})
    })
    app.Get("/api/v1/issue", models.GetIssues)
    app.Get("/api/v1/issue/:id", models.GetIssue)
    app.Post("/api/v1/issue", models.CreateIssue)
    app.Delete("/api/v1/issue/:id", models.DeleteIssue)
}

func initDB () {
    var err error
    db.DBConn, err = gorm.Open(sqlite.Open("spartan.db"))
    if err != nil {
        panic("failed to connect database")
    }
    fmt.Print("Connection opened to database\n")
    models.Migrate()
    fmt.Print("Database migrated\n")
}

func main() {
    engine := html.New("./views", ".html")
    app := fiber.New(fiber.Config{Views: engine})
    app.Use(cors.New())

    initDB()

    setupRoutes(app)

    log.Fatal(app.Listen(":3000"))
}

