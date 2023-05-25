package modela

import (
	"cadigo-api/graph/modelgraph"

	"github.com/jinzhu/copier"
)

type Payment struct {
	ID                 string  `json:"id,omitempty" copier:"Id"`
	TransactionID      *int    `json:"transactionId,omitempty" copier:"TransactionId"`
	Amount             *int    `json:"amount,omitempty" copier:"Amount"`
	OrderNo            *string `json:"orderNo,omitempty" copier:"OrderNo"`
	CustomerID         *string `json:"customerId,omitempty" copier:"CustomerId"`
	BankCode           *string `json:"bankCode,omitempty" copier:"BankCode"`
	PaymentDate        *string `json:"paymentDate,omitempty" copier:"PaymentDate"`
	PaymentStatus      *int    `json:"paymentStatus,omitempty" copier:"PaymentStatus"`
	BankRefCode        *string `json:"bankRefCode,omitempty" copier:"BankRefCode"`
	CurrentDate        *string `json:"currentDate,omitempty" copier:"CurrentDate"`
	CurrentTime        *string `json:"currentTime,omitempty" copier:"CurrentTime"`
	PaymentDescription *string `json:"paymentDescription,omitempty" copier:"PaymentDescription"`
	CreditCardToken    *string `json:"creditCardToken,omitempty" copier:"CreditCardToken"`
	Currency           *string `json:"currency,omitempty" copier:"Currency"`
	CustomerName       *string `json:"customerName,omitempty" copier:"CustomerName"`
	CheckSum           *string `json:"checkSum,omitempty" copier:"CheckSum"`
}

func (this Payment) Parse(graph modelgraph.PaymentInput) error {
	return copier.Copy(&this, &graph)
}

func (this Payment) ToGraph() modelgraph.Payment {
	g := modelgraph.Payment{}
	copier.Copy(&g, &this)
	return g
}
