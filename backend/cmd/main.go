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

	routes.Health(router)
	routes.TodoRoutes(router)
	routes.UserRoutes(router)

	fmt.Println("Server is running...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
