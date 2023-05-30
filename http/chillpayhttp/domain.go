package chillpayhttp

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/sirupsen/logrus"
)

type ErrorMessageResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type GeneratePaylinkRequest struct {
}

type PaylinkGenerateRequest struct {
	ProductImage       string `json:"productImage"`
	ProductName        string `json:"productName"`
	ProductDescription string `json:"productDescription"`
	PaymentLimit       int    `json:"paymentLimit"`
	StartDate          string `json:"startDate"`
	ExpiredDate        string `json:"expiredDate"`
	Currency           string `json:"currency"`
	Amount             int    `json:"amount"`
	Checksum           string `json:"checksum"`
}

func (h *PaylinkGenerateRequest) GenerateChecksum(md5key string) {
	str := fmt.Sprintf("%s%s%d%s%s%s%d%s",
		h.ProductName,
		h.ProductDescription,
		h.PaymentLimit,
		h.StartDate,
		h.ExpiredDate,
		h.Currency,
		h.Amount,
		md5key,
	)

	logrus.Info("sum:", str)

	hash := md5.Sum([]byte(str))
	h.Checksum = hex.EncodeToString(hash[:])
}

type TransctionIDRequest struct {
	TransactionId string `json:"TransactionId"`
	Checksum      string `json:"Checksum"`
}

func (h *TransctionIDRequest) GenerateChecksum(md5key string) {
	str := fmt.Sprintf("%s%s",
		h.TransactionId,
		md5key,
	)

	hash := md5.Sum([]byte(str))
	h.Checksum = hex.EncodeToString(hash[:])
}

type PaylinkDetailsRequest struct {
	PayLinkId string `json:"PayLinkId"`
	Checksum  string `json:"Checksum"`
}

func (h *PaylinkDetailsRequest) GenerateChecksum(md5key string) {
	str := fmt.Sprintf("%s%s",
		h.PayLinkId,
		md5key,
	)

	hash := md5.Sum([]byte(str))
	h.Checksum = hex.EncodeToString(hash[:])
}

type PaylinktransactionRequest struct {
	TransactionId int64  `json:"TransactionId"`
	Checksum      string `json:"Checksum"`
}

func (h *PaylinktransactionRequest) GenerateChecksum(md5key string) {
	str := fmt.Sprintf("%d%s",
		h.TransactionId,
		md5key,
	)

	hash := md5.Sum([]byte(str))
	h.Checksum = hex.EncodeToString(hash[:])
}

type PaylinkGenerateResponse struct {
	Status  int                         `json:"status"`
	Message string                      `json:"message"`
	Data    PaylinkGenerateResponseData `json:"data"`
}

type PaylinkGenerateResponseData struct {
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

type DetailByTransctionResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		PayLinkID           int     `json:"payLinkId"`
		ProductName         string  `json:"productName"`
		ProductDescription  string  `json:"productDescription"`
		TransactionID       int     `json:"transactionId"`
		TransactionDate     string  `json:"transactionDate"`
		PaymentDate         string  `json:"paymentDate"`
		CustomerName        string  `json:"customerName"`
		CustomerPhoneNumber string  `json:"customerPhoneNumber"`
		ChannelName         string  `json:"channelName"`
		Amount              float64 `json:"amount"`
		Fee                 float64 `json:"fee"`
		Discount            float64 `json:"discount"`
		NetAmount           float64 `json:"netAmount"`
		Currency            string  `json:"currency"`
		PaymentStatus       string  `json:"paymentStatus"`
	} `json:"data"`
}
