package modeld

import (
	"cadigo-api/app/modela"
	"time"
)

type Payment struct {
	BaseBSONModel      `bson:",inline"`
	TransactionID      int    `json:"TransactionId"`
	Amount             int    `json:"Amount"`
	OrderNo            string `json:"OrderNo"`
	PaymentID          string `json:"PaymentId"`
	BankCode           string `json:"BankCode"`
	PaymentDate        string `json:"PaymentDate"`
	PaymentStatus      int    `json:"PaymentStatus"`
	BankRefCode        string `json:"BankRefCode"`
	CurrentDate        string `json:"CurrentDate"`
	CurrentTime        string `json:"CurrentTime"`
	PaymentDescription string `json:"PaymentDescription"`
	CreditCardToken    string `json:"CreditCardToken"`
	Currency           string `json:"Currency"`
	PaymentName        string `json:"PaymentName"`
	CheckSum           string `json:"CheckSum"`
}

func (this Payment) Init() Payment {
	this.BaseBSONModel.CreatedAt = time.Now()
	this.BaseBSONModel.UpdatedAt = time.Now()

	return this
}

func (this Payment) Update() Payment {
	this.BaseBSONModel.UpdatedAt = time.Now()

	return this
}

func (this Payment) ToPayment() modela.Payment {
	return modela.Payment{
		ID:                 this.RawID.Hex(),
		TransactionID:      &this.TransactionID,
		Amount:             &this.Amount,
		OrderNo:            &this.OrderNo,
		BankCode:           &this.BankCode,
		PaymentDate:        &this.PaymentDate,
		PaymentStatus:      &this.PaymentStatus,
		BankRefCode:        &this.BankRefCode,
		CurrentDate:        &this.Currency,
		CurrentTime:        &this.CurrentTime,
		PaymentDescription: &this.PaymentDescription,
		CreditCardToken:    &this.CreditCardToken,
		Currency:           &this.Currency,
		CheckSum:           &this.CheckSum,
		// CustomerName: this.Cu,
		// CustomerID: this.cu,
		// CustomerID: ,
	}
}
