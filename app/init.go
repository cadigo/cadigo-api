package app

import (
	"cadigo-api/app/handlers/caddyhandler"
	"cadigo-api/app/handlers/coursegolfhandler"
	"cadigo-api/app/handlers/customerhandler"
	"cadigo-api/app/handlers/paymenthandler"
	"cadigo-api/app/injectors"
	"cadigo-api/app/services/caddyservice"
	"cadigo-api/db/mongodb/repositories/caddyrepository"
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

func bookingHandlerInit() *caddyhandler.Handler {
	mongodbConnector, err := injectors.ProvideMongoDBConnector(&generalConfig)
	if err != nil {
		panic(err)
	}
	baseMongoRepo := injectors.ProvideBaseMongoRepo(&generalConfig, mongodbConnector)

	caddyRepo := caddyrepository.NewRepository(baseMongoRepo)
	caddyServ := caddyservice.NewService(caddyRepo)
	return caddyhandler.NewHandler(caddyServ)
}

func courseGolfHandlerInit() *coursegolfhandler.Handler {
	mongodbConnector, err := injectors.ProvideMongoDBConnector(&generalConfig)
	if err != nil {
		panic(err)
	}
	baseMongoRepo := injectors.ProvideBaseMongoRepo(&generalConfig, mongodbConnector)

	caddyRepo := caddyrepository.NewRepository(baseMongoRepo)
	caddyServ := caddyservice.NewService(caddyRepo)
	return caddyhandler.NewHandler(caddyServ)
}

func customerHandlerInit() *customerhandler.Handler {
	mongodbConnector, err := injectors.ProvideMongoDBConnector(&generalConfig)
	if err != nil {
		panic(err)
	}
	baseMongoRepo := injectors.ProvideBaseMongoRepo(&generalConfig, mongodbConnector)

	caddyRepo := caddyrepository.NewRepository(baseMongoRepo)
	caddyServ := caddyservice.NewService(caddyRepo)
	return caddyhandler.NewHandler(caddyServ)
}

func paymentHandlerInit() *paymenthandler.Handler {
	mongodbConnector, err := injectors.ProvideMongoDBConnector(&generalConfig)
	if err != nil {
		panic(err)
	}
	baseMongoRepo := injectors.ProvideBaseMongoRepo(&generalConfig, mongodbConnector)

	caddyRepo := caddyrepository.NewRepository(baseMongoRepo)
	caddyServ := caddyservice.NewService(caddyRepo)
	return caddyhandler.NewHandler(caddyServ)
}
