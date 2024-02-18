package models

import (
	"gitlab.com/alistairr/spartan/db"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Username  string
}

type Project struct {
	gorm.Model
	Name        string
	Description string
	Status      string
	Issues      []Issue `gorm:"foreignKey:ProjectRefer"`
}

type Issue struct {
	gorm.Model
	Title        string `json:"title"`
	Description  string `json:"description"`
	Status       string `json:"status,omitempty,default:\"open\""`
	ProjectRefer uint   `json:"project_refer,omitempty,default:0"`
}

func Migrate() {
	dberr := db.DBConn.AutoMigrate(&User{}, &Project{}, &Issue{})
	if dberr != nil {
		panic(dberr)
	}
}
