package main

import (
	"net/http"

	"gitlab.com/alistairr/spartan/db"
	"gitlab.com/alistairr/spartan/helpers"
	"gitlab.com/alistairr/spartan/models"

	"github.com/gin-gonic/gin"
)

const (
	ContentTypeBinary = "application/octet-stream"
	ContentTypeForm   = "application/x-www-form-urlencoded"
	ContentTypeJSON   = "application/json"
	ContentTypeHTML   = "text/html; charset=utf-8"
	ContentTypeText   = "text/plain; charset=utf-8"
)

func Home(c *gin.Context) {
	render(c, gin.H{}, "index.html")
}

func GenerateUrl() string {
	var url string

	urlExist := true
	for urlExist {
		url = helpers.GenerateRandomString(7)
		//Verify is the url is unique
		project := models.Project{}
		db.DB.Where("url = ?", url).First(&project)
		if project.ID == 0 {
			urlExist = false
		}
	}
	return url
}

func CreateProject(c *gin.Context) {
	url := GenerateUrl()
	var project models.Project
	project.Url = url
	db.DB.Create(&project)

	c.Redirect(http.StatusMovedPermanently, "/"+url)
	c.Abort()
}

func ProjectList(c *gin.Context) {
	url := c.Param("url")
	var project models.Project
	var copyAlert bool // DEFAULT -> FALSE

	db.DB.Preload("Project.Issues").Where("ProjectID = ?", url).First(&project)

	if project.ID != 0 {
		if true {
			copyAlert = true
		}
		render(c, gin.H{"title": "Project List", "project": project.Issues, "copyAlert": copyAlert}, "project.html")
	}
}

func Issue(c *gin.Context) {
	url := c.Param("url")

	switch c.Request.Method {
	case "POST":
		{
			title := c.PostForm("name")
			description := c.PostForm("description")
			status := "Open"
			issue := models.Issue{
				Title:       title,
				Description: description,
				Status:      status,
				ProjectID:   url,
			}
			db.DB.Create(&issue)

			c.HTML(http.StatusOK, "issue.html", issue)
		}
	}
}
