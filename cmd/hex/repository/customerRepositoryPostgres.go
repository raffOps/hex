package repository

import (
	"errors"
	"fmt"
	"github.com/rjribeiro/hex/cmd/hex/domain"
	"github.com/rjribeiro/hex/cmd/hex/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

// customer represents a domain.Customer entity in Postgres Database.
type customer struct {
	gorm.Model
	Id          string `gorm:"primaryKey"`
	Name        string
	City        string
	ZipCode     string
	DateOfBirth time.Time
	Status      bool
}

// CustomerRepositoryPostgres is a struct that represents a Postgres implementation of the CustomerRepository interface.
type CustomerRepositoryPostgres struct {
	conn *gorm.DB
}

// NewCustomerRepositoryPostgres initializes a new instance of CustomerRepositoryPostgres with the provided gorm.DB connection.
// It performs auto migration on the customer table and panics if an error occurs.
// It returns the initialized CustomerRepositoryPostgres instance.
func NewCustomerRepositoryPostgres(conn *gorm.DB) CustomerRepositoryPostgres {
	err := conn.AutoMigrate(&customer{})
	if err != nil {
		panic(err)
	}
	return CustomerRepositoryPostgres{conn: conn}
}

// FindAll returns a slice of domain.Customer and an error.
// It retrieves all customers from the database and maps them into domain.Customer structs.
func (s CustomerRepositoryPostgres) FindAll() ([]domain.Customer, error) {
	var customers []customer
	result := s.conn.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	var domainCustomers []domain.Customer
	for _, c := range customers {
		domainCustomers = append(domainCustomers, domain.Customer{
			Id:          c.Id,
			Name:        c.Name,
			City:        c.City,
			ZipCode:     c.ZipCode,
			DateOfBirth: c.DateOfBirth,
			Status:      c.Status,
		})

	}
	return domainCustomers, nil
}

// FindById gets a customer by id and returns a domain.Customer and an error.
func (s CustomerRepositoryPostgres) FindById(id string) (*domain.Customer, error) {
	var c customer
	result := s.conn.First(&c, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logger.Debug(fmt.Sprintf("Customer %s not found", id))
		return nil, nil
	} else if result.Error != nil {
		logger.Error("Connection with postgres not working")
		return nil, result.Error
	}

	return &domain.Customer{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		ZipCode:     c.ZipCode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.Status,
	}, nil
}

func sanityCheck() {
	if _, ok := os.LookupEnv("POSTGRES_HOST"); !ok {
		panic("POSTGRES_HOST environment variable is not set")
	}

	if _, ok := os.LookupEnv("POSTGRES_USER"); !ok {
		panic("POSTGRES_USER environment variable is not set")
	}

	if _, ok := os.LookupEnv("POSTGRES_PASSWORD"); !ok {
		panic("POSTGRES_PASSWORD environment variable is not set")
	}

	if _, ok := os.LookupEnv("POSTGRES_DB"); !ok {
		panic("POSTGRES_DB environment variable is not set")
	}
}

// GetPostgresConnection returns a PostgreSQL database connection using the environment variables:
// POSTGRES_HOST, POSTGRES_USER, POSTGRES_PASSWORD and POSTGRES_DB.
func GetPostgresConnection() *gorm.DB {
	sanityCheck()
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DB")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=UTC",
		host,
		user,
		password,
		database,
	)
	DB, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		panic("Unable to connect to database")
	}

	logger.Info("Connected to database")
	return DB
}

func ClosePostgresConnection(conn *gorm.DB) {
	sqlDB, _ := conn.DB()
	_ = sqlDB.Close()
	logger.Info("Closed database connection")
}
