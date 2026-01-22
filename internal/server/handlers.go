package server

import (
	"CustomerService/internal/modules/address"
	"CustomerService/internal/modules/branch"
	"CustomerService/internal/modules/user_branch"

	"gorm.io/gorm"
)

func NewHandlers(db *gorm.DB) *Handlers {
	addrRepo := address.NewRepository(db)
	addrService := address.NewService(addrRepo)
	addrHandler := address.NewHandler(addrService)

	branchRepo := branch.NewRepository(db)
	branchService := branch.NewService(branchRepo)
	branchHandler := branch.NewHandler(branchService)

	ubRepo := user_branch.NewRepository(db)
	ubService := user_branch.NewService(ubRepo)
	ubHandler := user_branch.NewHandler(ubService)

	return &Handlers{
		Address:    addrHandler,
		Branch:     branchHandler,
		UserBranch: ubHandler,
	}
}
