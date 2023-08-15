package main

import (
	// "task/api"
	"fmt"
	"task/config"
)

func main() {
	// api.setupRouter()
	_, err := config.InitializeDB()
	if err != nil {
		fmt.Println("Failed to connect to database")
		return
	}
}
