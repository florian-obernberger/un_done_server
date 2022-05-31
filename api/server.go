package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"unDoneServer/dtypes"
	"unDoneServer/pwd"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	*mux.Router

	entries map[string]dtypes.TodoEntry
}

func InitServer() *Server {
	s := &Server{
		Router:  mux.NewRouter(),
		entries: map[string]dtypes.TodoEntry{},
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/api/add", s.addTodoEntries()).Methods("POST")
	// TODO: create seperate function
	s.HandleFunc("/api/update", s.addTodoEntries()).Methods("POST")

	s.HandleFunc("/api/get/all", s.getAllTodoEntries()).Methods("GET")
	s.HandleFunc("/api/get/new", s.getNewTodoEntries()).Methods("GET")

	s.HandleFunc("/api/exists/{ids:((\\w+,?)+}", s.doesTodoEntryExist()).Methods("GET")
}

func validateRequest(w http.ResponseWriter, r *http.Request) bool {
	k := r.FormValue("key")
	fmt.Printf("key: %s\n", k)

	if len(k) == 0 {
		http.Error(w, "Provide an API key using ?key=key", http.StatusUnauthorized)
		log.Warn("Request denied: wrong password")
		return false
	} else if !pwd.ValidatePasswordWithStored(k, pwd.HashFile) {
		http.Error(w, fmt.Sprintf("Password %s could not be verified", k), http.StatusUnauthorized)
		log.Warn("Request denied: wrong password")
		return false
	}
	return true
}

func (s *Server) addTodoEntries() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !validateRequest(w, r) {
			return
		}

		var t []dtypes.TodoEntry
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf("Couldn't decode incoming TodoEntry: %s", err.Error())
			return
		}

		for _, e := range t {
			s.entries[e.ID] = e
		}

		log.Infof("Added %d TodoEntries", len(t))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(t); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatalf("Error encoding incoming TodoEntry: %s", err.Error())
			return
		}
	}

}

func (s *Server) getAllTodoEntries() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !validateRequest(w, r) {
			return
		}

		t := s.entries
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(t); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatalf("Error encoding outgoing TodoEntries: %s", err.Error())
			return
		}
	}
}

func (s *Server) doesTodoEntryExist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !validateRequest(w, r) {
			return
		}

		ids := strings.Split(mux.Vars(r)["id"], ",")
		var e bool

		for _, id := range ids {
			_, e = s.entries[id]

			if e {
				break
			}
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(e); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatalf("Error encoding outgoing bool: %s", err.Error())
			return
		}
	}
}

func (s *Server) getNewTodoEntries() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !validateRequest(w, r) {
			return
		}

	}
}
