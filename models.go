// models.go
package main

type Issue struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"`
    Tags        string `json:"tags"`
    Reporter    string `json:"reporter"`
    Assignee    string `json:"assignee"`
    Links       string `json:"links"`
}

