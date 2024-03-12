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

func CreateUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	var newUser models.User
	newUser.Name = r.FormValue("name")
	newUser.Email = r.FormValue("email")

	// TODO: Validate and sanitize user data (if needed)
	err = db.DBConn.Create(&newUser).Error
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	w.Header().Set("HX-Trigger", "refresh-user-list")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser) // Return the newly created user
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	dbconn := db.DBConn
	if err := dbconn.Find(&users).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Render the template with data
	err := templates.ExecuteTemplate(w, "user-list.html", users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userid")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You've requested the user: %s\n", userId)
}

