package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"gitlab.com/alistairr/spartan/db"
	"gitlab.com/alistairr/spartan/models"
)

// Parse templates once when initializing handlers
var templates *template.Template

func init() {
	templateFiles, err := filepath.Glob("views/*")
	if err != nil {
		panic(err)
	}
	templates = template.Must(template.ParseFiles(templateFiles...))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	var issues []models.Issue
	dbconn := db.DBConn
	if err := dbconn.Find(&issues).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Render the template with data
	w.WriteHeader(http.StatusOK)
	err := templates.ExecuteTemplate(w, "issues_list.tmpl", issues)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	user := r.PathValue("userid")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You've requested user: %s\n", user)
}

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You've requested projects\n")
}

func ProjectHandler(w http.ResponseWriter, r *http.Request) {
	projectId := r.PathValue("projectid")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You've requested the project: %s\n", projectId)
}

func IssuesHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You've requested issues\n")
}

func IssueHandler(w http.ResponseWriter, r *http.Request) {
	issueId := r.PathValue("issueid")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You've requested the issue: %s\n", issueId)
}
