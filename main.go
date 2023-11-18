package main

import (
	"fmt"
	"interview-bank/database"
	"interview-bank/routes"
	"log"
)

func main() {
	// Initialize the database
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Create a new Gin router using SetupRouter
	router := routes.SetupRouter()
	// Start the server
	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	err = router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
