package modela

import (
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
	ID                 string  `json:"Id"`
	PayLinkID          int     `json:"payLinkId"`
	ProductImage       string  `json:"productImage"`
	ProductName        string  `json:"productName"`
	ProductDescription string  `json:"productDescription"`
	Amount             float64 `json:"amount"`
	Currency           string  `json:"currency"`
	CreatedDate        string  `json:"createdDate"`
	StartDate          string  `json:"startDate"`
	ExpiredDate        string  `json:"expiredDate"`
	PaymentLimit       int     `json:"paymentLimit"`
	Status             string  `json:"status"`
	PayLinkToken       string  `json:"payLinkToken"`
	PaymentURL         string  `json:"paymentUrl"`
	QrImage            string  `json:"qrImage"`
}

func (this Payment) Init(res chillpayhttp.PaylinkGenerateResponse) (*Payment, error) {
	err := copier.Copy(&this, &res.Data)
	if err != nil {
		return nil, err
	}

	return &this, nil
}

// func (this Payment) ToGraph() modelgraph.Payment {
// 	g := modelgraph.Payment{}
// 	copier.Copy(&g, &this)
// 	return g
// }

// func (this Payment) ToChillpay() chillpayhttp.PaylinkGenerateRequest {
// 	p := chillpayhttp.PaylinkGenerateRequest{}
// 	copier.Copy(&p, &this)
// 	return p
// }
