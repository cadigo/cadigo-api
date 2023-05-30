package paymentservice

import (
	"cadigo-api/app/interface/paymentinterface"
	"cadigo-api/app/modela"
	"cadigo-api/http/chillpayhttp"
	"context"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

type Service struct {
	repo     paymentinterface.PaymentRepository
	chillpay chillpayhttp.ChillpayHTTP
}

func NewService(repo paymentinterface.PaymentRepository, chillpay chillpayhttp.ChillpayHTTP) paymentinterface.PaymentService {
	return &Service{
		repo:     repo,
		chillpay: chillpay,
	}
}

func (serv *Service) Create(ctx context.Context, record *modela.PaymentRequest) (*modela.Payment, error) {
	var (
		chillpay chillpayhttp.PaylinkGenerateRequest
	)

	chillpay = record.ToChillpay()

	res, err := serv.chillpay.GetPaylinkGenerate(&chillpay)

	if err != nil {
		logrus.Info(err)
		return nil, fmt.Errorf("pay link error")
	}

	k, err := json.Marshal(res)
	fmt.Println(string(k))

	a, err := new(modela.Payment).Init(*res)
	if err != nil {
		logrus.Info(err)
		return nil, fmt.Errorf("create payment error")
	}

	return serv.repo.Create(ctx, a)
}

func (serv *Service) GetAll(ctx context.Context, pagination modela.Pagination) (result []*modela.Payment, total int64, err error) {
	return serv.repo.GetAll(ctx, pagination)
}

func (serv *Service) Update(ctx context.Context, argID string, record *modela.Payment) (*modela.Payment, error) {
	return serv.repo.Update(ctx, argID, record)
}

func (serv *Service) GetByID(ctx context.Context, id string) (result *modela.Payment, err error) {
	return serv.repo.GetByID(ctx, id)
}
