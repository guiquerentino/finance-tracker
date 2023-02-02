package structs

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	TransactionID   string
	Value           float64
	TransactionName string
	TransactionDate time.Time
}

func NewTransaction(value float64, transactionName string) *Transaction {
	return &Transaction{TransactionID: uuid.NewString(), Value: value, TransactionName: transactionName, TransactionDate: time.Now()}
}
