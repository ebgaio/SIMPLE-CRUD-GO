package main

import (
	"fmt"
	"github.com/ebgaio/simple-go-mod/config"
	"github.com/ebgaio/simple-go-mod/handlers"
	"github.com/ebgaio/simple-go-mod/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Entry point of the application

func main() {

	// Initialize DB connection
	db := config.SetupDB()

	// // Ensure the DB connection is closed when the application exits
	defer db.Close()

	_, err := db.Exec(models.CreateTableSQL)

	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	} else {
		fmt.Println("Table created")
	}

	router := mux.NewRouter()

	taskhandler := handlers.NewTaskHandler(db)

	router.HandleFunc("/tasks", taskhandler.ReadTasks).Methods("GET")
	router.HandleFunc("/tasks", taskhandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskhandler.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", taskhandler.DeleteTask).Methods("DELETE")

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
