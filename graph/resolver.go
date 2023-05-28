package graph

import (
	"cadigo-api/app/handlers/bookinghandler"
	"cadigo-api/app/handlers/caddyhandler"
	"cadigo-api/app/handlers/chathandler"
	"cadigo-api/app/handlers/coursegolfhandler"
	"cadigo-api/app/handlers/customerhandler"
	"cadigo-api/app/handlers/paymenthandler"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CaddyHandler      *caddyhandler.Handler
	BookingHandler    *bookinghandler.Handler
	CoursegolfHandler *coursegolfhandler.Handler
	CustomerHandler   *customerhandler.Handler
	PaymentHandler    *paymenthandler.Handler
	ChatHandler       *chathandler.Handler
}
