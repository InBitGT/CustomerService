package address

import (
	"CustomerService/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupAddressRoutes(api *mux.Router, handler *Handler) {
	addr := api.PathPrefix("/addresses").Subrouter()

	// Internos (sin JWT, con internal key)
	internal := addr.PathPrefix("/internal").Subrouter()
	internal.Use(middleware.InternalKeyMiddleware)
	internal.HandleFunc("", handler.CreateInternal).Methods("POST")
	internal.HandleFunc("/{id}", handler.DeleteInternal).Methods("DELETE")

	// Protegidos (JWT)
	protected := addr.NewRoute().Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("", handler.Create).Methods("POST")
	protected.HandleFunc("", handler.GetAll).Methods("GET")
	protected.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	protected.HandleFunc("/{id}", handler.Update).Methods("PUT")
	protected.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}
