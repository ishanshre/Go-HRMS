package models

import "go.starlark.net/lib/time"

type Employee struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Salary     float32   `json:"salary"`
	Age        float32   `json:"age"`
	Created_at time.Time `json:"created_at"`
}

func NewEmployee(name string, salary, age float32) (*Employee, error) {
	return &Employee{
		Name:   name,
		Salary: salary,
		Age:    age,
	}, nil
}
