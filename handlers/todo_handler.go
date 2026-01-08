package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var todos = []map[string]interface{}{}
var idCounter = 1

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, todo := range todos {
		if todo["id"].(int) == id {
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo map[string]interface{}
	json.NewDecoder(r.Body).Decode(&todo)

	todo["id"] = idCounter
	idCounter++
	todo["status"] = "pending"
	todo["created_at"] = time.Now()

	// Pastikan ada field description
	if todo["description"] == nil {
		todo["description"] = ""
	}

	todos = append(todos, todo)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, todo := range todos {
		if todo["id"].(int) == id {
			var requestData map[string]interface{}
			err := json.NewDecoder(r.Body).Decode(&requestData)
			
			if err == nil && len(requestData) > 0 {
				// EDIT: Update title dan description
				if title, ok := requestData["title"].(string); ok && title != "" {
					todos[i]["title"] = title
				}
				if description, ok := requestData["description"].(string); ok {
					todos[i]["description"] = description
				}
			} else {
				// TOGGLE: Toggle status tanpa body
				if todo["status"].(string) == "pending" {
					todos[i]["status"] = "done"
				} else {
					todos[i]["status"] = "pending"
				}
			}
			
			json.NewEncoder(w).Encode(todos[i])
			return
		}
	}
	
	w.WriteHeader(http.StatusNotFound)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, todo := range todos {
		if todo["id"].(int) == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	
	w.WriteHeader(http.StatusNotFound)
}