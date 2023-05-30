package modela

import (
	"cadigo-api/graph/modelgraph"
	"cadigo-api/http/chillpayhttp"

	"github.com/jinzhu/copier"
)

type PaymentRequest struct {
	ProductImage       string `json:"ProductImage"`
	ProductName        string `json:"ProductName"`
	ProductDescription string `json:"ProductDescription"`
	PaymentLimit       int    `json:"PaymentLimit"`
	StartDate          string `json:"StartDate"`
	ExpiredDate        string `json:"ExpiredDate"`
	Currency           string `json:"Currency"`
	Amount             int    `json:"Amount"`
}

func (this PaymentRequest) ToChillpay() chillpayhttp.PaylinkGenerateRequest {
	p := chillpayhttp.PaylinkGenerateRequest{}
	copier.Copy(&p, &this)
	return p
}

func (this PaymentRequest) Init(res *Booking) (*PaymentRequest, error) {
	err := copier.Copy(&this, res)
	if err != nil {
		return nil, err
	}

	return &this, nil
}

// func (this PaylinkGenerateResponse) ToModelA() modela.Payment {
// 	m := modela.Payment{}
// 	copier.Copy(&m, &this)
// 	return m
// }

type Payment struct {
	ID                 string  `json:"id" copier:"ID"`
	PayLinkID          int     `json:"payLinkId" copier:"PayLinkID"`
	ProductImage       string  `json:"productImage" copier:"ProductImage"`
	ProductName        string  `json:"productName" copier:"ProductName"`
	ProductDescription string  `json:"productDescription" copier:"ProductDescription"`
	Amount             float64 `json:"amount" copier:"Amount"`
	Currency           string  `json:"currency" copier:"Currency"`
	CreatedDate        string  `json:"createdDate" copier:"CreatedDate"`
	StartDate          string  `json:"startDate" copier:"StartDate"`
	ExpiredDate        string  `json:"expiredDate" copier:"ExpiredDate"`
	PaymentLimit       int     `json:"paymentLimit" copier:"PaymentLimit"`
	Status             string  `json:"status" copier:"Status"`
	PayLinkToken       string  `json:"payLinkToken" copier:"PayLinkToken"`
	PaymentURL         string  `json:"paymentUrl" copier:"PaymentURL"`
	QrImage            string  `json:"qrImage" copier:"QrImage"`
}

func (this Payment) Init(res chillpayhttp.PaylinkGenerateResponse) (*Payment, error) {
	err := copier.Copy(&this, &res.Data)
	if err != nil {
		return nil, err
	}

	return &this, nil
}

func (this Payment) ToGraph() modelgraph.Payment {

	a := int(this.Amount)

	return modelgraph.Payment{
		ID:                 this.ID,
		PayLinkID:          &this.PayLinkID,
		ProductImage:       &this.ProductImage,
		ProductName:        &this.ProductName,
		ProductDescription: &this.ProductDescription,
		Amount:             &a,
		Currency:           &this.Currency,
		CreatedDate:        &this.CreatedDate,
		StartDate:          &this.StartDate,
		ExpiredDate:        &this.ExpiredDate,
		PaymentLimit:       &this.PaymentLimit,
		Status:             &this.Status,
		PayLinkToken:       &this.PayLinkToken,
		PaymentURL:         &this.PaymentURL,
		QRImage:            &this.QrImage,
	}
}

// func (this Payment) ToChillpay() chillpayhttp.PaylinkGenerateRequest {
// 	p := chillpayhttp.PaylinkGenerateRequest{}
// 	copier.Copy(&p, &this)
// 	return p
// }
