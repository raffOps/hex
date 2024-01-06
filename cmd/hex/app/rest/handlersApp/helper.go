package handlersApp

import (
	"encoding/json"
	"github.com/rjribeiro/hex/cmd/hex/errs"
	"github.com/rjribeiro/hex/cmd/hex/logger"
	"net/http"
	"reflect"
)

var knownErrorTypes = []reflect.Type{
	reflect.TypeOf(errs.CustomerNotFoundError{}),
	reflect.TypeOf(errs.DuplicateAccountError{}),
	reflect.TypeOf(errs.InvalidAmountError{}),
	reflect.TypeOf(errs.AccountNotFoundError{}),
	reflect.TypeOf(errs.InvalidAccountError{}),
	reflect.TypeOf(errs.InvalidAccountConfigError{}),
}

func IsKnownError(err error) bool {
	errType := reflect.TypeOf(err)
	logger.Debug("error type: " + errType.String())
	for _, knownType := range knownErrorTypes {
		if errType == knownType {
			return true
		}
	}
	return false
}

func HandleError(w http.ResponseWriter, err error) {
	if IsKnownError(err) {
		writeResponse(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}
	writeResponse(w, http.StatusInternalServerError, map[string]string{"message": "something goes wrong"})
	return
}

func writeResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}
