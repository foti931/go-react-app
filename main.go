package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/models"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfuly Migratde")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&models.User{}, &models.Task{})
}
