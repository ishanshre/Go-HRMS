package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ishanshre/go-hrm/models"
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
	employees, err := s.store.ListEmployees()
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, employees)
}

func (s *ApiServer) handleCreateEmployee(w http.ResponseWriter, r *http.Request) error {
	newEmployee := new(models.Employee)
	if err := json.NewDecoder(r.Body).Decode(&newEmployee); err != nil {
		return err
	}
	employee, err := models.NewEmployee(newEmployee.Name, newEmployee.Salary, newEmployee.Age)
	if err != nil {
		return err
	}
	if err := s.store.CreateEmployee(employee); err != nil {
		return err
	}
	return writeJSON(w, http.StatusCreated, ApiSuccess{Success: "employee successfully created"})
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
