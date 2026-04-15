package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/baniksudipta/transaction-manager/src/repository"
	"github.com/baniksudipta/transaction-manager/src/service"
)

type RequestHandler struct {
	accountService     service.AccountService
	transactionService service.TransactionService
}

func GetRequestMapping(accountService service.AccountService, transactionService service.TransactionService) *http.ServeMux {
	h := &RequestHandler{
		accountService:     accountService,
		transactionService: transactionService,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /accounts", h.createAccount)
	mux.HandleFunc("GET /accounts/{accountId}", h.getAccount)
	mux.HandleFunc("POST /transactions", h.createTransaction)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "endpoint not found", http.StatusNotFound)
	})

	return mux
}

func (h *RequestHandler) createAccount(w http.ResponseWriter, r *http.Request) {
	var req CreateAccountRequest

	if err := Decode(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.DocumentNumber == "" {
		http.Error(w, "document_number is required", http.StatusBadRequest)
		return
	}

	acc, err := h.accountService.CreateAccount(req.DocumentNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(acc)
}

func (h *RequestHandler) getAccount(w http.ResponseWriter, r *http.Request) {
	accountIdStr := r.PathValue("accountId")
	accountID, err := strconv.ParseInt(accountIdStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid account ID", http.StatusBadRequest)
		return
	}

	acc, err := h.accountService.GetAccount(accountID)
	if err != nil {
		if errors.Is(err, repository.ErrAccountNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(acc)
}

func (h *RequestHandler) createTransaction(w http.ResponseWriter, r *http.Request) {
	var req CreateTransactionRequest

	if err := Decode(r, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Amount.IsZero() {
		http.Error(w, "amount cannot be zero", http.StatusBadRequest)
		return
	}

	txn, err := h.transactionService.CreateTransaction(req.AccountID, req.OperationTypeID, req.Amount)

	if err != nil {
		if errors.Is(err, repository.ErrAccountNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(txn)
}

func Decode(r *http.Request, v any) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return nil
}
