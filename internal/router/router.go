package router

import (
	"go-crud-api/internal/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handlers.UserHandler, orderHandler *handlers.OrderHandler) *mux.Router {
	r := mux.NewRouter()

	// Маршруты для пользователей
	r.HandleFunc("/users/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/users/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.GetByID).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.Update).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.Delete).Methods("DELETE")
	r.HandleFunc("/users", userHandler.List).Methods("GET")

	// Маршруты для заказов
	r.HandleFunc("/orders", orderHandler.Create).Methods("POST")
	r.HandleFunc("/orders/{id}", orderHandler.GetByID).Methods("GET")
	r.HandleFunc("/orders/{id}", orderHandler.Update).Methods("PUT")
	r.HandleFunc("/orders/{id}", orderHandler.Delete).Methods("DELETE")
	r.HandleFunc("/orders", orderHandler.List).Methods("GET")
	r.HandleFunc("/users/{user_id}/orders", orderHandler.GetByUserID).Methods("GET")

	return r
}
