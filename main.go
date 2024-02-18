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
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("GET /status", handlers.StatusHandler)
	mux.HandleFunc("/users", handlers.UsersHandler)
	mux.HandleFunc("/users/{userid}", handlers.UserHandler)
	mux.HandleFunc("/projects", handlers.ProjectsHandler)
	mux.HandleFunc("/projects/{projectid}", handlers.ProjectHandler)
	mux.HandleFunc("/issues", handlers.IssuesHandler)
	mux.HandleFunc("/issues/{issueid}", handlers.IssueHandler)
	// static rroutes
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
