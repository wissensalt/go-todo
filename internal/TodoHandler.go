package internal

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type TodoHandler struct {
}

var todoService = TodoService{}

func (receiver TodoHandler) FindById(writer http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idParam)
	todo := todoService.FindById(id)
	if todo == nil {
		http.Error(writer, "Todo Not Found", http.StatusNotFound)
	}

	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(todo)
	if err != nil {
		http.Error(writer, "Failed to Get Todo", http.StatusInternalServerError)
	}
}

func (receiver TodoHandler) ListTodos(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(todoService.ListTodos())
	if err != nil {
		http.Error(writer, "Failed to List Todos", http.StatusInternalServerError)

		return
	}
}

func (receiver TodoHandler) CreateTodo(writer http.ResponseWriter, request *http.Request) {
	var todo Todo
	err := json.NewDecoder(request.Body).Decode(&todo)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(todoService.createTodo(todo))
}

func (receiver TodoHandler) UpdateTodo(writer http.ResponseWriter, request *http.Request) {
	var todo Todo
	err := json.NewDecoder(request.Body).Decode(&todo)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(todoService.updateTodo(todo))
}

func (receiver TodoHandler) DeleteTodo(writer http.ResponseWriter, request *http.Request) {
	idParam := chi.URLParam(request, "id")
	id, _ := strconv.Atoi(idParam)
	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(todoService.deleteTodo(id))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
