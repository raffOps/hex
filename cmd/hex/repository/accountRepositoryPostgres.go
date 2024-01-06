package repository

import (
	"errors"
	"fmt"
	"github.com/rjribeiro/hex/cmd/hex/domain"
	"github.com/rjribeiro/hex/cmd/hex/errs"
	"github.com/rjribeiro/hex/cmd/hex/logger"
	"gorm.io/gorm"
	"time"
)

const errPostgresUniqueViolation = 23505

type account struct {
	gorm.Model
	AccountId   string `gorm:"not null;uniqueIndex;primary"`
	CustomerId  string `gorm:"not null"`
	OpeningDate time.Time
	AccountType string
	Amount      float64
	Status      bool
}

func (a account) ToDomainAccount() domain.Account {
	return domain.Account{
		AccountId:   a.AccountId,
		CustomerId:  a.CustomerId,
		OpeningDate: a.OpeningDate,
		AccountType: a.AccountType,
		Amount:      a.Amount,
		Status:      a.Status,
	}
}

type AccountConfig struct {
	AccountId   string
	CustomerId  string
	OpeningDate time.Time
	AccountType string
	Amount      float64
	Status      bool
}

func newAccount() account {
	return account{}
}

type AccountRepositoryPostgres struct {
	conn         *gorm.DB
	customerRepo domain.CustomerRepository
}

func NewAccountRepositoryPostgres(conn *gorm.DB) AccountRepositoryPostgres {
	err := conn.AutoMigrate(&account{})
	if err != nil {
		panic(err)
	}
	return AccountRepositoryPostgres{conn: conn}
}

func (s AccountRepositoryPostgres) Save(account domain.Account) (*domain.Account, error) {
	result := s.conn.Create(&account)
	if result.Error != nil && result.Error.Error() == "ERROR: duplicate key value violates unique constraint \"idx_accounts_account_id\" (SQLSTATE 23505)" {
		return nil, errs.DuplicateAccountError{AccountId: account.AccountId}
	}
	if result.Error != nil {
		return nil, errs.RepositoryError{Err: result.Error}
	}
	return &account, nil
}

func (s AccountRepositoryPostgres) FindById(id string) (*domain.Account, error) {
	var account account
	result := s.conn.First(&account, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logger.Debug(fmt.Sprintf("Account %s not found", id))
		return nil, errs.AccountNotFoundError{Id: id}
	}
	if result.Error != nil {
		return nil, errs.RepositoryError{Err: result.Error}
	}
	domainAccount := account.ToDomainAccount()
	return &domainAccount, nil
}

func (s AccountRepositoryPostgres) UpdateAmount(accountId string, amount float64) (*domain.Account, error) {
	var account account
	result := s.conn.First(&account, accountId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errs.AccountNotFoundError{Id: accountId}
	}
	if result.Error != nil {
		return nil, errs.RepositoryError{Err: result.Error}
	}
	account.Amount = amount
	result = s.conn.Save(&account)
	if result.Error != nil {
		return nil, errs.RepositoryError{Err: result.Error}
	}
	domainAccount := account.ToDomainAccount()
	return &domainAccount, nil
}

func (s AccountRepositoryPostgres) BeginTransaction() interface{} {
	return s.conn.Begin()
}

func (s AccountRepositoryPostgres) RollbackTransaction() interface{} {
	return s.conn.Rollback()
}

func (s AccountRepositoryPostgres) CommitTransaction() interface{} {
	return s.conn.Commit()
}
