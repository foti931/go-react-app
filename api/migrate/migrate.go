package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/models"
	"log"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfuly Migratde")
	defer db.CloseDB(dbConn)
	if err := dbConn.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		log.Fatalln("migration falied")
	}
}
