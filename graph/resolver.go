package graph

import "cadigo-api/app/handlers/caddyHandler"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CaddyHandler *caddyHandler.Handler
}
