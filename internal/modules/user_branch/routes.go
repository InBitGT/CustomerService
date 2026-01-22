package user_branch

import (
	"CustomerService/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupUserBranchRoutes(api *mux.Router, handler *Handler) {
	ub := api.PathPrefix("/user-branches").Subrouter()

	protected := ub.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("", handler.Create).Methods("POST")
	protected.HandleFunc("", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.Update).Methods("PUT")
	protected.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
