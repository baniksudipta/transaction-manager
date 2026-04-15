package handler

import (
	"github.com/baniksudipta/transaction-manager/src/entities"
	"github.com/shopspring/decimal"
)

type CreateTransactionRequest struct {
	AccountID       int64                  `json:"account_id"`
	OperationTypeID entities.OperationType `json:"operation_type_id"`
	Amount          decimal.Decimal        `json:"amount"`
}
