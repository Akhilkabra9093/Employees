package main

import (
	"Employees/api"
	"Employees/internal"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello, World!")
	internal.InitDB()
	defer internal.CloseDB()

	router := api.SetupRouter(internal.GetDB())

	//Setup API routes and start server
	if err := api.StartServer(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
