package server

import (
	"CustomerService/internal/modules/address"
	"CustomerService/internal/modules/branch"
	"CustomerService/internal/modules/user_branch"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	Router *mux.Router
	DB     *gorm.DB
}

type Handlers struct {
	Address    *address.Handler
	Branch     *branch.Handler
	UserBranch *user_branch.Handler
}
