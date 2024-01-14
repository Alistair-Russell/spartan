package models

import (
	"gitlab.com/alistairr/spartan/db"

	"gorm.io/gorm"
)

type Issue struct {
	gorm.Model
	Title       string
	Description string
	Status      string
}

func Migrate() {
	db.DBConn.AutoMigrate(&Issue{})
}
