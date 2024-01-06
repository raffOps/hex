package domain

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindById(id string) (*Customer, error)
}

type AccountRepository interface {
	Save(account Account) (*Account, error)
	FindById(id string) (*Account, error)
	UpdateAmount(string, float64) (*Account, error)
	BeginTransaction() interface{}
	RollbackTransaction() interface{}
	CommitTransaction() interface{}
}
