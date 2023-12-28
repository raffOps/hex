package repository

import (
	"github.com/rjribeiro/hex/cmd/hex/domain"
	"gorm.io/gorm"
	"time"
)

type customer struct {
	gorm.Model
	Id          string `gorm:"primaryKey"`
	Name        string
	City        string
	ZipCode     string
	DateOfBirth time.Time
	Status      string
}

type CustomerRepositoryPostgres struct {
	conn *gorm.DB
}

func NewCustomerRepositoryPostgres(conn interface{}) CustomerRepositoryPostgres {
	db, ok := conn.(*gorm.DB)
	if !ok {
		panic("Failed to assert interface{} to *gorm.DB")
	}
	err := db.AutoMigrate(&customer{})
	if err != nil {
		panic(err)
	}
	return CustomerRepositoryPostgres{conn: db}
}

func (r CustomerRepositoryPostgres) FindAll() ([]domain.Customer, error) {
	var customers []customer
	err := r.conn.Find(&customers).Error
	if err != nil {
		return nil, err
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
