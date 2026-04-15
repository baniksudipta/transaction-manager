package repository

import (
	"errors"
	"sync"
	"sync/atomic"

	"github.com/baniksudipta/transaction-manager/src/entities"
)

var ErrTransactionNotFound = errors.New("transaction not found")

type TransactionRepository interface {
	Save(transaction entities.Transaction) (entities.Transaction, error)
	FindByID(transactionID int64) (entities.Transaction, error)
}

type InMemoryTransactionRepository struct {
	transactionStore sync.Map
	transactionIdNum atomic.Int64
}

func NewInMemoryTransactionRepository() *InMemoryTransactionRepository {
	return &InMemoryTransactionRepository{}
}

func (r *InMemoryTransactionRepository) Save(transaction entities.Transaction) (entities.Transaction, error) {
	newTransactionID := r.transactionIdNum.Add(1)
	transaction.TransactionID = newTransactionID

	r.transactionStore.Store(newTransactionID, transaction)
	return transaction, nil
}

func (r *InMemoryTransactionRepository) FindByID(transactionID int64) (entities.Transaction, error) {
	if existing, ok := r.transactionStore.Load(transactionID); ok {
		return existing.(entities.Transaction), nil
	}
	return entities.Transaction{}, ErrTransactionNotFound
}
