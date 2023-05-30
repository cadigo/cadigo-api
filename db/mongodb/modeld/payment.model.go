package modeld

import (
	"cadigo-api/app/modela"
	"time"
)

type Payment struct {
	BaseBSONModel      `bson:",inline"`
	PayLinkID          int     `bson:"payLinkId"`
	ProductImage       string  `bson:"productImage"`
	ProductName        string  `bson:"productName"`
	ProductDescription string  `bson:"productDescription"`
	Amount             float64 `bson:"amount"`
	Currency           string  `bson:"currency"`
	CreatedDate        string  `bson:"createdDate"`
	StartDate          string  `bson:"startDate"`
	ExpiredDate        string  `bson:"expiredDate"`
	PaymentLimit       int     `bson:"paymentLimit"`
	Status             string  `bson:"status"`
	PayLinkToken       string  `bson:"payLinkToken"`
	PaymentURL         string  `bson:"paymentUrl"`
	QrImage            string  `bson:"qrImage"`
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
		PayLinkID:          this.PayLinkID,
		ProductImage:       this.ProductImage,
		ProductName:        this.ProductName,
		ProductDescription: this.ProductDescription,
		Amount:             this.Amount,
		Currency:           this.Currency,
		CreatedDate:        this.CreatedDate,
		StartDate:          this.StartDate,
		ExpiredDate:        this.ExpiredDate,
		PaymentLimit:       this.PaymentLimit,
		Status:             this.Status,
		PayLinkToken:       this.PayLinkToken,
		PaymentURL:         this.PaymentURL,
		QrImage:            this.QrImage,
	}
}
