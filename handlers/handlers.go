package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
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

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := io.WriteString(w, `{"alive": true}`)
	if err != nil {
		panic(err)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateUsersHandler(w http.ResponseWriter, r *http.Request) {
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

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
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

func CreateIssuesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	var newIssue models.Issue
	newIssue.Title = r.FormValue("title")
	newIssue.Description = r.FormValue("description")

	// TODO: Validate and sanitize issue data (if needed)
	err = db.DBConn.Create(&newIssue).Error
	if err != nil {
		http.Error(w, "Failed to create issue", http.StatusInternalServerError)
		return
	}
	w.Header().Set("HX-Trigger", "refresh-issue-list")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newIssue) // Return the newly created issue
}

func ListIssuesHandler(w http.ResponseWriter, r *http.Request) {
	var issues []models.Issue
	dbconn := db.DBConn
	if err := dbconn.Find(&issues).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Render the template with data
	err := templates.ExecuteTemplate(w, "issue-list.html", issues)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func IssueHandler(w http.ResponseWriter, r *http.Request) {
	issueId := r.PathValue("issueid")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You've requested the issue: %s\n", issueId)
}
