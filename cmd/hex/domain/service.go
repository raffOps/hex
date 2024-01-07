package domain

type CustomerService interface {
	GetAllCustomer() ([]Customer, error)
	GetCustomerById(string) (*Customer, error)
}

type AccountService interface {
	NewAccount(account Account) (*Account, error)
	Deposit(accountId string, amountToDeposit float64) (*Account, error)
}
