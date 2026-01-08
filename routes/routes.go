package routes

import (
	"github.com/gorilla/mux"
	"todo-api/handlers"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", handlers.GetTodoByID).Methods("GET") // NEW
	router.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")
}