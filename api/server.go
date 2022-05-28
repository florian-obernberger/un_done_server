package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"unDoneServer/dtypes"
	"unDoneServer/pwd"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	*mux.Router

	entries []dtypes.TodoEntry
}

func InitServer() *Server {
	s := &Server{
		Router:  mux.NewRouter(),
		entries: []dtypes.TodoEntry{},
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/todo-entries", s.addTodoEntries()).Queries("key").Methods("POST")
	s.HandleFunc("/todo-entries", s.getTodoEntries()).Queries("key").Methods("GET")
}

func (s *Server) addTodoEntries() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		k := r.FormValue("key")
		if !pwd.ValidatePasswordWithStored(k, pwd.HashFile) {
			http.Error(w, fmt.Sprintf("Password %s could not be verified", k), http.StatusUnauthorized)
			log.Warn("Request was denied due to wrong password")
			return
		}

		var t []dtypes.TodoEntry
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf("Couldn't decode incoming TodoEntry: %s", err.Error())
			return
		}
		s.entries = append(s.entries, t...)
		log.Infof("Added %d TodoEntries", len(t))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(t); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatalf("Error encoding incoming TodoEntry: %s", err.Error())
			return
		}
	}

}

func (s *Server) getTodoEntries() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		k := r.FormValue("key")
		if !pwd.ValidatePasswordWithStored(k, pwd.HashFile) {
			http.Error(w, fmt.Sprintf("Password %s could not be verified", k), http.StatusUnauthorized)
			log.Warn("Request was denied due to wrong password")
			return
		}

		t := s.entries
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(t); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatalf("Error encoding outgoing TodoEntries: %s", err.Error())
			return
		}
		log.Infof("Sent %d TodoEntries", len(t))
	}
}
