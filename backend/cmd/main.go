package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AdarshJha-1/Taskify/backend/internal/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	credentials := handlers.AllowCredentials()
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete})
	origins := handlers.AllowedOrigins([]string{os.Getenv("ORIGINS")})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	// Health check route
	routes.Health(router)

	// Todo routes
	routes.TodoRoutes(router)

	// User routes
	routes.UserRoutes(router)

	// Running serer at post 3000

	fmt.Println("Server is running...")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(credentials, methods, origins, headers)(router)))
}
