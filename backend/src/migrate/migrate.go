package main

import (
	"fmt"

	"github.com/Shota-Kurahashi/echo-next/src/db"
	"github.com/Shota-Kurahashi/echo-next/src/models"
)

func main() {
	database := db.Connect()
	defer db.Close(database)
	defer fmt.Println("DB connection closed.")

	database.AutoMigrate(&models.User{})
}
