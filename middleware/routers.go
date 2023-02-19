package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ishanshre/go-hrm/database"
)

type ApiServer struct {
	listenAddr string
	store      database.Storage
}

func NewApiServer(listenAddr string, store database.Storage) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/employee", makeHttpHandleFunc(s.handleEmployees))
	router.HandleFunc("/employee/{id}", makeHttpHandleFunc(s.handleEmployeeById))
	log.Printf("Starting server at localhost:%v", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

type ApiFunc func(http.ResponseWriter, *http.Request) error // creating signature for our HandleFunc

func makeHttpHandleFunc(f ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
	Error string `json:"error"`
}

type ApiSuccess struct {
	Success string `json:"Success"`
}
