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
	Title        string
	Description  string
	Status       string
	ProjectRefer uint
}

func Migrate() {
	dberr := db.DBConn.AutoMigrate(&User{}, &Project{}, &Issue{})
	if dberr != nil {
		panic(dberr)
	}
}
