package repository

import (
	"errors"
	"sync"
	"sync/atomic"

	"github.com/baniksudipta/transaction-manager/src/entities"
)

var ErrAccountNotFound = errors.New("account not found")

type AccountRepository interface {
	Save(account entities.Account) (entities.Account, error)
	FindByID(accountID int64) (entities.Account, error)
	FindByDocumentNumber(documentNumber string) (entities.Account, error)
}

type InMemoryAccountRepository struct {
	accountStore              sync.Map
	documentNumberToAccountID sync.Map
	accountIdNum              atomic.Int64
}

func NewInMemoryAccountRepository() *InMemoryAccountRepository {
	return &InMemoryAccountRepository{}
}

func (r *InMemoryAccountRepository) Save(account entities.Account) (entities.Account, error) {
	newAccountID := r.accountIdNum.Add(1)
	account.AccountID = newAccountID

	r.accountStore.Store(newAccountID, account)
	r.documentNumberToAccountID.Store(account.DocumentNumber, account)

	return account, nil
}

func (r *InMemoryAccountRepository) FindByID(accountID int64) (entities.Account, error) {
	if existing, ok := r.accountStore.Load(accountID); ok {
		return existing.(entities.Account), nil
	}
	return entities.Account{}, ErrAccountNotFound
}

func (r *InMemoryAccountRepository) FindByDocumentNumber(documentNumber string) (entities.Account, error) {
	if existing, ok := r.documentNumberToAccountID.Load(documentNumber); ok {
		return existing.(entities.Account), nil
	}
	return entities.Account{}, ErrAccountNotFound
}
