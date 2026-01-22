package branch

import (
	"CustomerService/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupBranchRoutes(api *mux.Router, handler *Handler) {
	br := api.PathPrefix("/branches").Subrouter()

	protected := br.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("", handler.Create).Methods("POST")
	protected.HandleFunc("", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.Update).Methods("PUT")
	protected.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
