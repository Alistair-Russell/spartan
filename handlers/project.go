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

func CreateProjectsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	var newProject models.Project
	newProject.Name = r.FormValue("name")
	newProject.Description = r.FormValue("description")

	// TODO: Validate and sanitize project data (if needed)
	err = db.DBConn.Create(&newProject).Error
	if err != nil {
		http.Error(w, "Failed to create project", http.StatusInternalServerError)
		return
	}
	w.Header().Set("HX-Trigger", "refresh-project-list")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProject) // Return the newly created project
}

func ListProjectsHandler(w http.ResponseWriter, r *http.Request) {
	var projects []models.Project
	dbconn := db.DBConn
	if err := dbconn.Find(&projects).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Render the template with data
	err := templates.ExecuteTemplate(w, "project-list.html", projects)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetProjectByIdHandler(w http.ResponseWriter, r *http.Request) {
	id = r.PathValue("id")
	var project models.Project
	if err := db.DBConn.First(&project, id).Error; err != nil {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(project)
}

