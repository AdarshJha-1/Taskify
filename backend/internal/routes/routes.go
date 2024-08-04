package routes

import (
	"net/http"

	"github.com/AdarshJha-1/Taskify/backend/internal/handler"
	"github.com/AdarshJha-1/Taskify/backend/internal/middleware"
	"github.com/gorilla/mux"
)

func TodoRoutes(router *mux.Router) {
	router.HandleFunc("/todos", middleware.AuthMiddleware(handler.GetTodo)).Methods(http.MethodGet)
	router.HandleFunc("/todos", middleware.AuthMiddleware(handler.CreateTodos)).Methods(http.MethodPost)
}

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/signup", handler.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/login", handler.LoginUser).Methods(http.MethodPost)
}

func Health(router *mux.Router) {
	router.HandleFunc("/health", handler.HealthCheck).Methods(http.MethodGet)
}
