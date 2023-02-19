package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	id, err := getID(r)
	if err != nil {
		return err
	}
	employee := new(models.Employee)
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		log.Printf("error in parsing data: %s", err)
		return err
	}
	if err := s.store.UpdateEmployee(id, employee); err != nil {
		log.Printf("error in updating the employee: %s", err)
		return err
	}
	log.Printf("Account with id %v successfully updated", id)
	return writeJSON(w, http.StatusOK, ApiSuccess{Success: "successfullly updated the employee"})

}

func (s *ApiServer) handleDeleteEmployee(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}
	if err := s.store.DeleteEmployee(id); err != nil {
		log.Printf("error in deleting employee with id %v: error: %s", id, err)
		return err
	}
	return writeJSON(w, http.StatusOK, ApiSuccess{Success: "success in deleting employee"})
}

func (s *ApiServer) handleGetEmployeeById(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}
	employee, err := s.store.GetEmployeeById(id)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, employee)
}

func getID(r *http.Request) (int64, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return int64(id), fmt.Errorf("error in parsing the id: %s", err)
	}
	return int64(id), nil
}
