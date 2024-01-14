package handlers

import (
	"io"

	"github.com/gorilla/mux"

	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]
	issue := vars["issue"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You've requested the issue: %s in project %s\n", issue, project)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["userid"]
	issue := vars["issue"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You've requested the issue: %s in project %s\n", issue, project)
}

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]
	issue := vars["issue"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You've requested the issue: %s in project %s\n", issue, project)
}

func ProjectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]
	issue := vars["issue"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You've requested the issue: %s in project %s\n", issue, project)
}

func IssuesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]
	issue := vars["issue"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You've requested the issue: %s in project %s\n", issue, project)
}

func IssueHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	project := vars["project"]
	issue := vars["issue"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You've requested the issue: %s in project %s\n", issue, project)
}
