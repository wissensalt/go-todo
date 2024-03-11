package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/wissensalt/go-todo/internal"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Golang Todo Api"))
	})

	r.Mount("/todos", TodoRouter())

	_ = http.ListenAndServe("localhost:8080", r)
}

func TodoRouter() http.Handler {
	todoHandler := internal.TodoHandler{}
	r := chi.NewRouter()
	r.Get("/", todoHandler.ListTodos)
	r.Get("/{id}", todoHandler.FindById)
	r.Post("/", todoHandler.CreateTodo)
	r.Put("/", todoHandler.UpdateTodo)
	r.Delete("/{id}", todoHandler.DeleteTodo)

	return r
}
