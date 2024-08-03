package routes

import (
	"net/http"

	"github.com/AdarshJha-1/Taskify/backend/internal/handler"
	"github.com/gorilla/mux"
)

func TodoRoutes(router *mux.Router) {
	router.HandleFunc("/todos/{id}", handler.GetTodo()).Methods(http.MethodGet)
	router.HandleFunc("/todos/{id}", handler.CreateTodos()).Methods(http.MethodPost)
}

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/user", handler.CreateUser()).Methods(http.MethodPost)
}

func Health(router *mux.Router) {
	router.HandleFunc("/health", handler.HealthCheck()).Methods(http.MethodGet)
}
