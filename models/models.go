package models

import (
	"gitlab.com/alistairr/spartan/db"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Issue struct {
	gorm.Model
	Title       string
	Description string
	Status      string

	ProjectID string
}

type Project struct {
	gorm.Model
	Url    string `gorm:"not null; unique"`
	Name   string
	Issues []Issue `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func Migrate() {
	db.DB.AutoMigrate(&Issue{}, &Project{})
}
