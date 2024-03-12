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
