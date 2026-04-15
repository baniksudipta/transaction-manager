package service

import (
	"testing"

	"github.com/baniksudipta/transaction-manager/src/repository"
)

func TestAccountService_CreateAccount(t *testing.T) {
	repo := repository.NewInMemoryAccountRepository()
	s := NewAccountService(repo)
	acc, err := s.CreateAccount("12345678900")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if acc.DocumentNumber != "12345678900" {
		t.Errorf("expected document number to be 12345678900, got %s", acc.DocumentNumber)
	}

	// Test existing account
	acc2, err := s.CreateAccount("12345678900")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if acc2.AccountID != acc.AccountID {
		t.Errorf("expected account ID to be %d, got %d", acc.AccountID, acc2.AccountID)
	}
}

func TestAccountService_GetAccount(t *testing.T) {
	repo := repository.NewInMemoryAccountRepository()
	s := NewAccountService(repo)
	acc, _ := s.CreateAccount("12345678900")

	foundAcc, err := s.GetAccount(acc.AccountID)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if foundAcc.AccountID != acc.AccountID {
		t.Errorf("expected account ID to be %d, got %d", acc.AccountID, foundAcc.AccountID)
	}

	// Test not found
	_, err = s.GetAccount(999)
	if err == nil {
		t.Error("expected error, got nil")
	}
}
