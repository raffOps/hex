package handlersApp

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rjribeiro/hex/cmd/hex/dto"
	"github.com/rjribeiro/hex/cmd/hex/logger"
	"github.com/rjribeiro/hex/cmd/hex/service"
	"net/http"
	"strconv"
)

type AccountHandlers struct {
	accountService  service.AccountService
	customerService service.CustomerService
}

func NewAccountHandlers(accountService service.AccountService, customerService service.CustomerService) AccountHandlers {
	return AccountHandlers{accountService: accountService, customerService: customerService}
}

func (ah AccountHandlers) Save(w http.ResponseWriter, r *http.Request) {
	var account dto.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		HandleError(w, err)
		return
	}
	savedAccount, err := ah.accountService.NewAccount(account)
	if err != nil {
		HandleError(w, err)
		return
	}

	writeResponse(w, http.StatusCreated, savedAccount)
}

func (ah AccountHandlers) Deposit(w http.ResponseWriter, r *http.Request) {
	var id = mux.Vars(r)["id"]
	var amount = mux.Vars(r)["amount"]
	amountParsed, _ := strconv.ParseFloat(amount, 64)
	logger.Debug("depositing amount " + amount + " to account " + id)

	depositedAccount, err := ah.accountService.Deposit(id, amountParsed)
	if err != nil {
		HandleError(w, err)
		return
	}

	writeResponse(w, http.StatusOK, depositedAccount)
}
