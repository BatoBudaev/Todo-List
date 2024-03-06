package main

import (
	"fmt"
	db2 "github.com/BatoBudaev/Todo-List/internal/db"
	"github.com/BatoBudaev/Todo-List/internal/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db, err := db2.InitDB("postgres", "1", "todo_database")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	r := mux.NewRouter()
	handlers.SetupRoutes(r, db)

	fmt.Println("Server is starting at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
