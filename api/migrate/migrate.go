package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/models"
	"log"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrate")
	defer db.CloseDB(dbConn)
	if err := dbConn.AutoMigrate(&models.User{}, &models.Task{}, &models.PasswordReset{}); err != nil {
		log.Fatalln("migration failed: ", err)
	}
}
