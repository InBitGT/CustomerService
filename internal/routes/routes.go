package routes

import (
	"CustomerService/internal/middleware"
	"CustomerService/internal/modules/address"
	"CustomerService/internal/modules/branch"
	"CustomerService/internal/modules/user_branch"

	"github.com/gorilla/mux"
)

type RouteHandlers interface {
	GetAddressHandler() *address.Handler
	GetBranchHandler() *branch.Handler
	GetUserBranchHandler() *user_branch.Handler
}

func SetupRoutes(router *mux.Router, handlers RouteHandlers) {
	router.Use(middleware.ContentTypeJSON)
	router.Use(middleware.Logger)
	router.Use(middleware.Recovery)

	api := router.PathPrefix("/api").Subrouter()

	address.SetupAddressRoutes(api, handlers.GetAddressHandler())
	branch.SetupBranchRoutes(api, handlers.GetBranchHandler())
	user_branch.SetupUserBranchRoutes(api, handlers.GetUserBranchHandler())
}
