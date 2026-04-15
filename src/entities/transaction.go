package entities

import (
	"time"

	"github.com/shopspring/decimal"
)

type OperationType int64

const (
	Invalid                  OperationType = 0
	NormalPurchase           OperationType = 1
	PurchaseWithInstallments OperationType = 2
	Withdrawal               OperationType = 3
	CreditVoucher            OperationType = 4
)

func (o OperationType) IsCredit() bool {
	return o == CreditVoucher
}

func (o OperationType) IsDebit() bool {
	return !o.IsCredit()
}

func (o OperationType) String() string {
	switch o {
	case NormalPurchase:
		return "Normal Purchase"
	case PurchaseWithInstallments:
		return "Purchase with installments"
	case Withdrawal:
		return "Withdrawal"
	case CreditVoucher:
		return "Credit Voucher"
	default:
		return "Invalid"
	}
}

type Transaction struct {
	TransactionID   int64           `json:"transaction_id"`
	AccountID       int64           `json:"account_id"`
	OperationTypeID OperationType   `json:"operation_type_id"`
	Amount          decimal.Decimal `json:"amount"`
	EventDate       time.Time       `json:"event_date"`
}
