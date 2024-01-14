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
	dberr := db.DBConn.AutoMigrate(&Issue{})
	if dberr != nil {
		panic(dberr)
	}
}
