package structs

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	transactionID   string
	value           float64
	transactionName string
	transactionDate time.Time
}

func NewTransaction(value float64, transactionName string) *Transaction {
	return &Transaction{transactionID: uuid.NewString(), value: value, transactionName: transactionName, transactionDate: time.Now()}
}

func (t *Transaction) SetValue(value float64) {
	t.value = value
}
func (t *Transaction) SetTransactionName(transactionName string) {
	t.transactionName = transactionName
}

func (t *Transaction) SetTransactionDate(transactionDate time.Time) {
	t.transactionDate = transactionDate
}
func (t *Transaction) SetTransactionID(transactionID string) {
	t.transactionID = transactionID
}

func (t *Transaction) GetTransactionId() string {
	return t.transactionID
}

func (t *Transaction) GetValue() float64 {
	return t.value
}

func (t *Transaction) GetTransactionName() string {
	return t.transactionName
}

func (t *Transaction) GetTransactionDate() time.Time {
	return t.transactionDate
}
