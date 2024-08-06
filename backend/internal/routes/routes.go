package routes

import (
	"net/http"

	"github.com/AdarshJha-1/Taskify/backend/internal/handler"
	"github.com/AdarshJha-1/Taskify/backend/internal/middleware"
	"github.com/gorilla/mux"
)

// TodoRoutes sets up the routes for todo-related operations
func TodoRoutes(router *mux.Router) {
	router.HandleFunc("/todos", middleware.AuthMiddleware(handler.GetTodo)).Methods(http.MethodGet)
	router.HandleFunc("/todos", middleware.AuthMiddleware(handler.CreateTodos)).Methods(http.MethodPost)
	router.HandleFunc("/todos/{id}", middleware.AuthMiddleware(handler.DeleteTodo)).Methods(http.MethodDelete)
	router.HandleFunc("/todos/{id}", middleware.AuthMiddleware(handler.ToggleIsCompletedTodo)).Methods(http.MethodPut)
}

// UserRoutes sets up the routes for user-related operations
func UserRoutes(router *mux.Router) {
	router.HandleFunc("/sign-up", handler.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/sign-in", handler.SignInUser).Methods(http.MethodPost)
	router.HandleFunc("/sign-out", handler.SignOutUser).Methods(http.MethodPost)
}

// Health sets up a route to check the health of the service
func Health(router *mux.Router) {
	router.HandleFunc("/health", handler.HealthCheck).Methods(http.MethodGet)
}
