package middleware

import (
	"fmt"
	"net/http"
)

func (s *ApiServer) handleEmployees(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleListEmployees(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateEmployee(w, r)
	}
	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *ApiServer) handleEmployeeById(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetEmployeeById(w, r)
	}
	if r.Method == "PUT" {
		return s.handleUpdateEmployee(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteEmployee(w, r)
	}
	return fmt.Errorf("method not allowed: %s", r.Method)
}

func (s *ApiServer) handleListEmployees(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleCreateEmployee(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleUpdateEmployee(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleDeleteEmployee(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleGetEmployeeById(w http.ResponseWriter, r *http.Request) error {
	return nil
}
