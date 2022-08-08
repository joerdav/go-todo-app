package api

import (
	"net/http"

	"github.com/joerdav/go-todo-app/db"
	"go.uber.org/zap"
)

type Server struct {
	mux   *http.ServeMux
	log   *zap.Logger
	store db.TodoStore
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
