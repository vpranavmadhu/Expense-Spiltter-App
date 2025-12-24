package main

import (
	"esapp/internal/database"
	"esapp/internal/routes"
	"log"
)

func main() {

	//connected to postgres and created db
	db, err := database.ConnectPostgres()

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	//created table inside db
	err = database.InitDb(db)

	//created gin frame and passed the db address
	r := routes.SetupRoutes(db)

	log.Println("Server running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
