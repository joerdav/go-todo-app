package api

import (
	"encoding/json"
	"net/http"

	"github.com/joerdav/go-todo-app/db"
)

func (s *Server) handleTodoPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t db.Todo
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
