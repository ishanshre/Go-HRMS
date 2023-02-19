package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/ishanshre/go-hrm/models"
	"github.com/joho/godotenv"
)

type Storage interface {
	CreateEmployee(*models.Employee) error
	DeleteEmployee(int) error
	UpdateEmployee(int, *models.Employee) error
	ListAccounts() ([]*models.Employee, error)
	GetAccountById(int) (*models.Employee, error)
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
	return nil
}

func (s *PostgresStore) DeleteEmployee(id int) error {
	return nil
}

func (s *PostgresStore) UpdateEmployee(id int, employee *models.Employee) error {
	return nil
}
func (s *PostgresStore) ListAccounts() ([]*models.Employee, error) {
	return nil, nil
}
func (s *PostgresStore) GetAccountById(id int) (*models.Employee, error) {
	return nil, nil
}
