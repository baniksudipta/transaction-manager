package repository

import (
	"testing"
	"time"

	"github.com/baniksudipta/transaction-manager/src/entities"
	"github.com/shopspring/decimal"
)

func TestInMemoryTransactionRepository(t *testing.T) {
	repo := NewInMemoryTransactionRepository()

	txn := entities.Transaction{
		AccountID:       1,
		OperationTypeID: entities.NormalPurchase,
		Amount:          decimal.NewFromFloat(-50.5),
		EventDate:       time.Now(),
	}

	// Test Save
	savedTxn, err := repo.Save(txn)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if savedTxn.TransactionID <= 0 {
		t.Errorf("expected transaction ID to be generated and positive")
	}

	// Test FindByID (Success)
	foundTxn, err := repo.FindByID(savedTxn.TransactionID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if foundTxn.AccountID != txn.AccountID {
		t.Errorf("expected account ID %d, got %d", txn.AccountID, foundTxn.AccountID)
	}

	// Test FindByID (Not Found)
	if _, err := repo.FindByID(999); err != ErrTransactionNotFound {
		t.Errorf("expected ErrTransactionNotFound, got %v", err)
	}
}
