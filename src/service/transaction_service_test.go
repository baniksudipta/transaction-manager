package service

import (
	"testing"

	"github.com/baniksudipta/transaction-manager/src/entities"
	"github.com/baniksudipta/transaction-manager/src/repository"
	"github.com/shopspring/decimal"
)

func TestTransactionService_CreateTransaction(t *testing.T) {
	repo := repository.NewInMemoryTransactionRepository()
	accountRepo := repository.NewInMemoryAccountRepository()
	s := NewTransactionService(repo, accountRepo)

	accountRepo.Save(entities.Account{DocumentNumber: "12345678900"})

	// Test debit operation
	txn, err := s.CreateTransaction(1, entities.NormalPurchase, decimal.NewFromFloat(50.5))
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if txn.Amount.IsPositive() {
		t.Errorf("expected amount to be negative for debit, got %v", txn.Amount)
	}
	if txn.Amount.String() != "-50.5" {
		t.Errorf("expected amount to be -50.5, got %v", txn.Amount)
	}
	if txn.EventDate.IsZero() {
		t.Error("expected event date to be set")
	}

	// Test credit operation
	txn2, err := s.CreateTransaction(1, entities.CreditVoucher, decimal.NewFromFloat(-100.0))
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if txn2.Amount.IsNegative() {
		t.Errorf("expected amount to be positive for credit, got %v", txn2.Amount)
	}
	if txn2.Amount.String() != "100" {
		t.Errorf("expected amount to be 100, got %v", txn2.Amount)
	}

	// Test account not found
	_, err = s.CreateTransaction(999, entities.NormalPurchase, decimal.NewFromFloat(50.5))
	if err != repository.ErrAccountNotFound {
		t.Errorf("expected ErrAccountNotFound, got %v", err)
	}
}

func TestTransactionService_GetTransaction(t *testing.T) {
	repo := repository.NewInMemoryTransactionRepository()
	accountRepo := repository.NewInMemoryAccountRepository()
	s := NewTransactionService(repo, accountRepo)

	accountRepo.Save(entities.Account{DocumentNumber: "12345678900"})

	txn, _ := s.CreateTransaction(1, entities.NormalPurchase, decimal.NewFromFloat(10.0))

	foundTxn, err := s.GetTransaction(txn.TransactionID)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if foundTxn.TransactionID != txn.TransactionID {
		t.Errorf("expected transaction ID %d, got %d", txn.TransactionID, foundTxn.TransactionID)
	}
}
