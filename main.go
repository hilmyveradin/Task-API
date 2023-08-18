package main

import (
	// "task/api"
	"fmt"
	"task/api"
	"task/config"
	"task/model"
)

func main() {
	db, err := config.InitializeDB()
	if err != nil {
		fmt.Println("Failed to connect to database")
		return
	}
	model.Migrate(db)
	r := api.SetupRouter(db)
	r.Run(":8080")
}
