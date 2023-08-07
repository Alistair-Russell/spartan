package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB Connection String
func dbvar() string {
	// // CONFIG VARS
	// DB_HOST := os.Getenv("DB_HOST")
	// DB_PORT := os.Getenv("DB_PORT")
	// DB_PASSWORD := os.Getenv("DB_PASSWORD")
	// DB_NAME := os.Getenv("DB_NAME")
	// DB_USER := os.Getenv("DB_USER")

	// dsn := ("host=" + DB_HOST +
	// 	" user=" + DB_USER +
	// 	" password=" + DB_PASSWORD +
	// 	" dbname=" + DB_NAME +
	// 	" port=" + DB_PORT +
	// 	" sslmode=disable TimeZone=Asia/Shanghai")

	// Local DB
	dsn := "host=localhost user=postgres password=postgres dbname=spartandb port=5432 sslmode=disable"

	return dsn
}

// Database Connection
var DB = func() (db *gorm.DB) {
	if db, err := gorm.Open("postgres", dbvar()); err != nil {
		fmt.Println("Connection to database failed", err)
		panic(err)
	} else {
		fmt.Println("Connected to database")
		return db
	}
}()
