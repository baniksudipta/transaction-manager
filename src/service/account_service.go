package service

import (
	"github.com/baniksudipta/transaction-manager/src/entities"
	"github.com/baniksudipta/transaction-manager/src/repository"
)

type AccountService interface {
	CreateAccount(documentNumber string) (entities.Account, error)
	GetAccount(accountID int64) (entities.Account, error)
}

type accountServiceImpl struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) AccountService {
	return &accountServiceImpl{repo: repo}
}

func (a *accountServiceImpl) CreateAccount(documentNumber string) (entities.Account, error) {
	if existing, err := a.repo.FindByDocumentNumber(documentNumber); err == nil {
		return existing, nil
	}

	newAccount := entities.Account{
		DocumentNumber: documentNumber,
	}

	return a.repo.Save(newAccount)
}

func (a *accountServiceImpl) GetAccount(accountID int64) (entities.Account, error) {
	return a.repo.FindByID(accountID)
}
