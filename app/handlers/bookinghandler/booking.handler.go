package bookinghandler

import (
	"cadigo-api/app/interface/bookinginterface"
	"cadigo-api/app/interface/caddyinterface"
	"cadigo-api/app/interface/coursegolfinterface"
	"cadigo-api/app/interface/customerinterface"
	"cadigo-api/app/interface/paymentinterface"
	"cadigo-api/app/modela"
	"cadigo-api/graph/modelgraph"
	"context"
	"fmt"
	"time"

	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	servBooking    bookinginterface.BookingService
	servCustomer   customerinterface.CustomerService
	servCourseGolf coursegolfinterface.CourseGolfService
	servCaddy      caddyinterface.CaddyService
	servPayment    paymentinterface.PaymentService
}

func NewHandler(servBooking bookinginterface.BookingService,
	servCustomer customerinterface.CustomerService,
	servCourseGolf coursegolfinterface.CourseGolfService,
	servCaddy caddyinterface.CaddyService,
	servPayment paymentinterface.PaymentService) *Handler {
	return &Handler{
		servBooking:    servBooking,
		servCustomer:   servCustomer,
		servCourseGolf: servCourseGolf,
		servCaddy:      servCaddy,
		servPayment:    servPayment,
	}
}

// Customer is the resolver for the customer field.
func (r *Handler) Customer(ctx context.Context, obj *modelgraph.Booking) (*modelgraph.Customer, error) {
	if obj != nil {
		d, err := r.servCustomer.GetByID(ctx, obj.CustomerID)
		if err != nil {
			return nil, nil
		}

		g := d.ToGraph()

		return &g, nil
	}

	return nil, nil
}

// CourseGolf is the resolver for the courseGolf field.
func (r *Handler) CourseGolf(ctx context.Context, obj *modelgraph.Booking) (*modelgraph.CourseGolf, error) {
	if obj != nil {
		d, err := r.servCourseGolf.GetByID(ctx, obj.CourseGolfID)
		if err != nil {
			return nil, nil
		}

		g := d.ToGraph()

		return &g, nil
	}

	return nil, nil
}

// Caddy is the resolver for the caddy field.
func (r *Handler) Caddy(ctx context.Context, obj *modelgraph.Booking) (*modelgraph.Caddy, error) {
	if obj != nil {
		d, err := r.servCaddy.GetByID(ctx, obj.CaddyID)
		if err != nil {
			return nil, nil
		}

		g := d.ToGraph()

		return &g, nil
	}

	return nil, nil
}

// Booking is the resolver for the booking field.
func (r *Handler) Booking(ctx context.Context, input modelgraph.BookingInput) (booking *modelgraph.Booking, err error) {
	// call chillpay
	courseGolf, err := r.servCourseGolf.GetByID(ctx, input.CourseGolfID)
	if err != nil {
		return nil, err
	}

	startDate := time.Now()
	expiredDate := startDate.Add(time.Minute * 30)
	// save payment
	paymentReq := &modela.PaymentRequest{
		ProductName:        courseGolf.Name,
		ProductImage:       "",
		ProductDescription: fmt.Sprintf("%v - %v #(%v)", input.TimeStart.Format("2006-01-02"), input.TimeEnd.Format("2006-01-02"), input.CaddyID),
		PaymentLimit:       int(1),
		StartDate:          startDate.Format("02/01/2006 15:04:05"),
		ExpiredDate:        expiredDate.Format("02/01/2006 15:04:05"),
		Currency:           "THB",
		Amount:             int(input.TotalNet * 100),
	}
	paymentRes, err := r.servPayment.Create(ctx, paymentReq)
	if err != nil {
		logrus.Info(err)
		return nil, fmt.Errorf("payment error")
	}

	record := modela.Booking{}
	err = copier.CopyWithOption(&record, &input, copier.Option{IgnoreEmpty: true})
	if err != nil {
		return nil, err
	}

	record.PaymentID = paymentRes.ID

	// Insert
	res, err := r.servBooking.Create(ctx, &record)
	if err != nil {
		return nil, err
	}

	c := res.ToGraph()

	return &c, nil
}

// GetBooking is the resolver for the getBooking field.
func (r *Handler) GetBooking(ctx context.Context, input modelgraph.GetBookingInput) (*modelgraph.Booking, error) {
	if input.ID != "" {
		r, err := r.servBooking.GetByID(ctx, input.ID)
		if err != nil {
			return nil, err
		}

		g := r.ToGraph()

		return &g, nil
	} else {
		return nil, nil
	}
}

func (r *Handler) Payment(ctx context.Context, obj *modelgraph.Booking) (*modelgraph.Payment, error) {
	if obj.PaymentID != nil {
		fmt.Println("AAA")
		r, err := r.servPayment.GetByID(ctx, *obj.PaymentID)
		if err != nil {
			return nil, err
		}

		g := r.ToGraph()

		return &g, nil
	} else {
		return nil, nil
	}
}

// GetBookings is the resolver for the getBookings field.
func (r *Handler) GetBookings(ctx context.Context, input modelgraph.BookingsInput) (data *modelgraph.BookingData, err error) {
	defaultPagination := new(modela.Pagination).Init()

	err = copier.Copy(&defaultPagination, &input.Pagination)
	if err != nil {
		return nil, err
	}

	res, total, err := r.servBooking.GetAll(ctx, defaultPagination)
	if err != nil {
		return nil, err
	}

	bookingData := []*modelgraph.Booking{}
	for _, v := range res {
		g := v.ToGraph()
		bookingData = append(bookingData, &g)
	}

	data = &modelgraph.BookingData{
		Data: bookingData,
		Pagination: &modelgraph.PaginationType{
			Page:  int(defaultPagination.Page),
			Limit: int(defaultPagination.Limit),
			Total: int(total),
		},
	}

	return data, nil
}
