package http

import (
	"mvrp/http/handler"
	"mvrp/http/middleware"
	"mvrp/http/router"
)

func Init() {
	/*
		Initialize the http layer.
	*/

	// Initialize the middleware.
	middleware.Init()

	// Initialize the handlers.
	handler.Init()

	// Initialize the router.
	// NOTE: The router must be initialized last.
	router.Init()
}
