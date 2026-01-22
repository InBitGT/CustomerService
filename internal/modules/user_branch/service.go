package user_branch

type Service interface {
	Create(tenantID uint, ub *UserBranch) (*UserBranch, error)
	GetByID(tenantID uint, id uint) (*UserBranch, error)
	GetAll(tenantID uint) ([]UserBranch, error)
	Update(tenantID uint, id uint, ub *UserBranch) (*UserBranch, error)
	Delete(tenantID uint, id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Create(tenantID uint, ub *UserBranch) (*UserBranch, error) {
	if err := s.repo.Create(tenantID, ub); err != nil {
		return nil, err
	}
	return ub, nil
}

func (s *service) GetByID(tenantID uint, id uint) (*UserBranch, error) {
	return s.repo.GetByID(tenantID, id)
}

func (s *service) GetAll(tenantID uint) ([]UserBranch, error) {
	return s.repo.GetAll(tenantID)
}

func (s *service) Update(tenantID uint, id uint, ub *UserBranch) (*UserBranch, error) {
	if err := s.repo.Update(tenantID, id, ub); err != nil {
		return nil, err
	}
	return s.repo.GetByID(tenantID, id)
}

func (s *service) Delete(tenantID uint, id uint) error {
	return s.repo.Delete(tenantID, id)
}
