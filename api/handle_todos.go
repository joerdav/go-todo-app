package api

import (
	"encoding/json"
	"net/http"

	"github.com/a-h/pathvars"
	"github.com/google/uuid"
	"github.com/joerdav/go-todo-app/db"
)

func (s *Server) handleTodoPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t db.Todo
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(struct {
				Message string `json:"message"`
			}{err.Error()})
			return
		}
		t.ID = uuid.NewString()
		err = s.store.Update(r.Context(), t)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(struct {
				Message string `json:"message"`
			}{err.Error()})
			return
		}
		err = json.NewEncoder(w).Encode(t)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(struct {
				Message string `json:"message"`
			}{err.Error()})
			return
		}
	}
}

func (s *Server) handleTodosGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// userId would usually come from some authentication
		userId := "u1"
		ts, err := s.store.LoadByUserID(r.Context(), userId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(struct {
				Message string `json:"message"`
			}{err.Error()})
			return
		}
		err = json.NewEncoder(w).Encode(ts)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(struct {
				Message string `json:"message"`
			}{err.Error()})
			return
		}
	}
}

var deleteTodoPath = pathvars.NewExtractor("*/todo/{id}")

func (s *Server) handleTodosDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v, ok := deleteTodoPath.Extract(r.URL)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(struct {
				Message string `json:"message"`
			}{"invalid path"})
			return
		}
		err := s.store.Delete(r.Context(), v["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(struct {
				Message string `json:"message"`
			}{err.Error()})
			return
		}
	}
}
