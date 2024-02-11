package handlers

import (
	"io"

	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := io.WriteString(w, `{"alive": true}`)
	if err != nil {
		panic(err)
	}
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You've requested users\n")
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
