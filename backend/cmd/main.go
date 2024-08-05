package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AdarshJha-1/Taskify/backend/internal/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Health check route
	routes.Health(router)

	// Todo routes
	routes.TodoRoutes(router)

	// User routes
	routes.UserRoutes(router)

	// Running serer at post 3000
	fmt.Println("Server is running...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
