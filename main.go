// main.go
package main

import (
	"gitlab.com/alistairr/spartan/db"
	"gitlab.com/alistairr/spartan/handlers"
	"gitlab.com/alistairr/spartan/models"

	"fmt"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDB() {
	var err error
	db.DBConn, err = gorm.Open(sqlite.Open("spartan.db"))
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Print("Connection opened to database\n")
	models.Migrate()
	fmt.Print("Database migrated\n")
}

func initRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// GET and POST routes for /users
	mux.HandleFunc("GET /users", handlers.ListUsersHandler)
	mux.HandleFunc("POST /users", handlers.CreateUsersHandler)
	mux.HandleFunc("GET /users/{userid}", handlers.GetUserByIdHandler)

	// GET and POST routes for /projects
	mux.HandleFunc("GET /projects", handlers.ListProjectsHandler)
	mux.HandleFunc("POST /projects", handlers.CreateProjectsHandler)
	mux.HandleFunc("GET /projects/{projectid}", handlers.GetProjectByIdHandler)

	// GET and POST routes for /issues
	mux.HandleFunc("GET /issues", handlers.ListIssuesHandler)
	mux.HandleFunc("POST /issues", handlers.CreateIssuesHandler)
	mux.HandleFunc("GET /issues/{issueid}", handlers.GetIssueByIdHandler)

	// Other routes...
	mux.HandleFunc("/", handlers.IndexHandler)

	// Static file routing
	mux.HandleFunc("GET /status", handlers.StatusHandler)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	return mux
}

func main() {
	initDB()
	handler := initRoutes()

	srv := &http.Server{
		Addr:         "127.0.0.1:3000",
		Handler:      handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
