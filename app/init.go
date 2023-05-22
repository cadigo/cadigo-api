package app

import (
	"cadigo-api/app/handlers/bookinghandler"
	"cadigo-api/app/handlers/caddyhandler"
	"cadigo-api/app/handlers/coursegolfhandler"
	"cadigo-api/app/handlers/customerhandler"
	"cadigo-api/app/handlers/paymenthandler"
	"cadigo-api/app/injectors"
	"cadigo-api/app/services/bookingservice"
	"cadigo-api/app/services/caddyservice"
	"cadigo-api/app/services/coursegolfservice"
	"cadigo-api/app/services/customerservice"
	"cadigo-api/app/services/paymentservice"
	"cadigo-api/db/mongodb/repositories/bookingrepository"
	"cadigo-api/db/mongodb/repositories/caddyrepository"
	"cadigo-api/db/mongodb/repositories/coursegolfrepository"
	"cadigo-api/db/mongodb/repositories/customerrepository"
	"cadigo-api/db/mongodb/repositories/paymentrepository"
)

func caddyHandlerInit() *caddyhandler.Handler {
	mongodbConnector, err := injectors.ProvideMongoDBConnector(&generalConfig)
	if err != nil {
		panic(err)
	}
	baseMongoRepo := injectors.ProvideBaseMongoRepo(&generalConfig, mongodbConnector)

	caddyRepo := caddyrepository.NewRepository(baseMongoRepo)
	caddyServ := caddyservice.NewService(caddyRepo)
	return caddyhandler.NewHandler(caddyServ)
}

func bookingHandlerInit() *bookinghandler.Handler {
	mongodbConnector, err := injectors.ProvideMongoDBConnector(&generalConfig)
	if err != nil {
		panic(err)
	}
	baseMongoRepo := injectors.ProvideBaseMongoRepo(&generalConfig, mongodbConnector)

	bookingRepo := bookingrepository.NewRepository(baseMongoRepo)
	bookingServ := bookingservice.NewService(bookingRepo)

	courseGolfRepo := coursegolfrepository.NewRepository(baseMongoRepo)
	courseGolfServ := coursegolfservice.NewService(courseGolfRepo)

	customerRepo := customerrepository.NewRepository(baseMongoRepo)
	customerServ := customerservice.NewService(customerRepo)

	caddyRepo := caddyrepository.NewRepository(baseMongoRepo)
	caddyServ := caddyservice.NewService(caddyRepo)

	return bookinghandler.NewHandler(bookingServ, customerServ, courseGolfServ, caddyServ)
}

func courseGolfHandlerInit() *coursegolfhandler.Handler {
	mongodbConnector, err := injectors.ProvideMongoDBConnector(&generalConfig)
	if err != nil {
		panic(err)
	}
	baseMongoRepo := injectors.ProvideBaseMongoRepo(&generalConfig, mongodbConnector)

	courseGolfRepo := coursegolfrepository.NewRepository(baseMongoRepo)
	courseGolfServ := coursegolfservice.NewService(courseGolfRepo)
	return coursegolfhandler.NewHandler(courseGolfServ)
}

func customerHandlerInit() *customerhandler.Handler {
	mongodbConnector, err := injectors.ProvideMongoDBConnector(&generalConfig)
	if err != nil {
		panic(err)
	}
	baseMongoRepo := injectors.ProvideBaseMongoRepo(&generalConfig, mongodbConnector)

	customerRepo := customerrepository.NewRepository(baseMongoRepo)
	customerServ := customerservice.NewService(customerRepo)
	return customerhandler.NewHandler(customerServ)
}

func paymentHandlerInit() *paymenthandler.Handler {
	mongodbConnector, err := injectors.ProvideMongoDBConnector(&generalConfig)
	if err != nil {
		panic(err)
	}
	baseMongoRepo := injectors.ProvideBaseMongoRepo(&generalConfig, mongodbConnector)

	paymentRepo := paymentrepository.NewRepository(baseMongoRepo)
	paymentServ := paymentservice.NewService(paymentRepo)
	return paymenthandler.NewHandler(paymentServ)
}
