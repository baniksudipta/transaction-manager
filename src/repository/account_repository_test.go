package repository

import (
	"testing"

	"github.com/baniksudipta/transaction-manager/src/entities"
)

func TestInMemoryAccountRepository(t *testing.T) {
	repo := NewInMemoryAccountRepository()

	acc := entities.Account{
		DocumentNumber: "12345678900",
	}

	// Test Save
	savedAcc, err := repo.Save(acc)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if savedAcc.AccountID <= 0 {
		t.Errorf("expected account ID to be generated and positive")
	}

	// Test FindByID (Success)
	foundAcc, err := repo.FindByID(savedAcc.AccountID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if foundAcc.DocumentNumber != acc.DocumentNumber {
		t.Errorf("expected document number %s, got %s", acc.DocumentNumber, foundAcc.DocumentNumber)
	}

	// Test FindByDocumentNumber (Success)
	foundAccByDoc, err := repo.FindByDocumentNumber(acc.DocumentNumber)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if foundAccByDoc.AccountID != savedAcc.AccountID {
		t.Errorf("expected account ID %d, got %d", savedAcc.AccountID, foundAccByDoc.AccountID)
	}

	// Test FindByID (Not Found)
	if _, err := repo.FindByID(999); err != ErrAccountNotFound {
		t.Errorf("expected ErrAccountNotFound, got %v", err)
	}

	// Test FindByDocumentNumber (Not Found)
	if _, err := repo.FindByDocumentNumber("0000"); err != ErrAccountNotFound {
		t.Errorf("expected ErrAccountNotFound, got %v", err)
	}
}
