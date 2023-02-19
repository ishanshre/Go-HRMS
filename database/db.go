package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/ishanshre/go-hrm/models"
	"github.com/joho/godotenv"
)

type Storage interface {
	CreateEmployee(*models.Employee) error
	DeleteEmployee(int64) error
	UpdateEmployee(int64, *models.Employee) error
	ListEmployees() ([]*models.Employee, error)
	GetEmployeeById(int64) (*models.Employee, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	// loading the environment files
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error in loading environment files: %s", err)
	}
	// connecting to postgres
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_CONN_STRING"))
	if err != nil {
		return nil, err
	}
	// now using db.Ping() we check if we are connected to the database
	if err := db.Ping(); err != nil {
		return nil, err
	}
	// if every thing is alright we create and return a PostgresStore struct variable
	return &PostgresStore{
		db: db,
	}, nil
}

// Creating a constructor Init. We will invoke it after database connection is successfull
func (s *PostgresStore) Init() error {
	return s.createEmployeeTable()
}

func (s *PostgresStore) createEmployeeTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS employee (
			id BIGSERIAL PRIMARY KEY,
			name VARCHAR(255),
			salary DECIMAL,
			age DECIMAL,
			create_at TIMESTAMPTZ
		);
	`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateEmployee(employee *models.Employee) error {
	query := `
		INSERT INTO employee (
			name, salary, age, create_at
		) VALUES ($1, $2, $3, $4);
	`
	_, err := s.db.Query(
		query,
		employee.Name,
		employee.Salary,
		employee.Age,
		employee.Created_at,
	)
	if err != nil {
		log.Printf("error in creating new employee: %s", err)
		return err
	}
	log.Printf("successfull inserting new employee")
	return nil
}

func (s *PostgresStore) DeleteEmployee(id int64) error {
	query := `
		DELETE FROM employee
		WHERE id = $1
	`
	s.db.Exec("COMMIT")
	rows, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows_affected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rows_affected == 0 {
		return fmt.Errorf("error: id %v does not exists", id)
	}
	return nil
}

func (s *PostgresStore) UpdateEmployee(id int64, employee *models.Employee) error {
	query := `
		UPDATE employee
		SET name = $2, salary = $3, age = $4
		WHERE id = $1;
		
	`
	s.db.Exec("COMMIT")
	_, err := s.db.Exec(
		query,
		id,
		employee.Name,
		employee.Salary,
		employee.Age,
	)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostgresStore) ListEmployees() ([]*models.Employee, error) {
	query := `
		SELECT * FROM employee;
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	employees := []*models.Employee{}
	for rows.Next() {
		employee, err := scanEmployee(rows)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil

}
func (s *PostgresStore) GetEmployeeById(id int64) (*models.Employee, error) {
	query := `
		SELECT * FROM employee
		WHERE id = $1;
	`
	rows, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanEmployee(rows)
	}
	return nil, fmt.Errorf("account with id %v not found", id)
}

func scanEmployee(rows *sql.Rows) (*models.Employee, error) {
	employee := new(models.Employee)
	err := rows.Scan(
		&employee.ID,
		&employee.Name,
		&employee.Salary,
		&employee.Age,
		&employee.Created_at,
	)
	return employee, err
}
