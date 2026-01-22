package server

import (
	"CustomerService/internal/modules/address"
	"CustomerService/internal/modules/branch"
	"CustomerService/internal/modules/user_branch"
)

func (h *Handlers) GetAddressHandler() *address.Handler {
	return h.Address
}

func (h *Handlers) GetBranchHandler() *branch.Handler {
	return h.Branch
}

func (h *Handlers) GetUserBranchHandler() *user_branch.Handler {
	return h.UserBranch
}
