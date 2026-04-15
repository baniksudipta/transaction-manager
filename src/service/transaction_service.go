package service

import (
	"errors"
	"time"

	"github.com/baniksudipta/transaction-manager/src/entities"
	"github.com/baniksudipta/transaction-manager/src/repository"
	"github.com/shopspring/decimal"
)

type TransactionService interface {
	CreateTransaction(accountID int64, operationTypeID entities.OperationType, amount decimal.Decimal) (entities.Transaction, error)
	GetTransaction(transactionID int64) (entities.Transaction, error)
}

type transactionServiceImpl struct {
	repo        repository.TransactionRepository
	accountRepo repository.AccountRepository
}

func NewTransactionService(repo repository.TransactionRepository, accountRepo repository.AccountRepository) TransactionService {
	return &transactionServiceImpl{repo: repo, accountRepo: accountRepo}
}

func (t *transactionServiceImpl) CreateTransaction(accountID int64, operationTypeID entities.OperationType,
	amount decimal.Decimal) (entities.Transaction, error) {

	if _, err := t.accountRepo.FindByID(accountID); err != nil {
		return entities.Transaction{}, err
	}

	if operationTypeID == entities.Invalid {
		return entities.Transaction{}, errors.New("invalid operation type")
	}

	if operationTypeID.IsDebit() {
		if amount.IsPositive() {
			amount = amount.Neg()
		}
	} else if operationTypeID.IsCredit() {
		if amount.IsNegative() {
			amount = amount.Abs()
		}
	}

	newTransaction := entities.Transaction{
		AccountID:       accountID,
		OperationTypeID: operationTypeID,
		Amount:          amount,
		EventDate:       time.Now(),
	}

	return t.repo.Save(newTransaction)
}

func (t *transactionServiceImpl) GetTransaction(transactionID int64) (entities.Transaction, error) {
	return t.repo.FindByID(transactionID)
}
