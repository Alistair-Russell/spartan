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

	"github.com/gorilla/mux"
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

func initRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.IndexHandler)
	r.HandleFunc("/users", handlers.UsersHandler).Methods("GET")
	r.HandleFunc("/users/{userid:[0-9]+}", handlers.UserHandler)
	r.HandleFunc("/projects", handlers.ProjectsHandler).Methods("GET")
	r.HandleFunc("/projects/{projectid}", handlers.ProjectHandler)
	r.HandleFunc("/issues", handlers.IssuesHandler).Methods("GET")
	r.HandleFunc("/issues/{issueid}", handlers.IssueHandler)
	// static rroutes
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	return r
}

func main() {
	initDB()
	// init router
	router := initRoutes()

	// Initiate server
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:3000",
		// Good practice: enforce timeouts
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
