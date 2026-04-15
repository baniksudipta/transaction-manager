package main

import (
	"log"
	"net/http"

	"github.com/baniksudipta/transaction-manager/src/handler"
	"github.com/baniksudipta/transaction-manager/src/repository"
	"github.com/baniksudipta/transaction-manager/src/service"
)

func main() {
	accountRepo := repository.NewInMemoryAccountRepository()
	transactionRepo := repository.NewInMemoryTransactionRepository()

	accountService := service.NewAccountService(accountRepo)
	transactionService := service.NewTransactionService(transactionRepo, accountRepo)

	requestMapping := handler.GetRequestMapping(accountService, transactionService)
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", requestMapping)
}
