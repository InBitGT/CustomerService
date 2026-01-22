package branch

type Service interface {
	Create(b *Branch) (*Branch, error)
	GetByID(id uint, tenantID uint) (*Branch, error)
	GetAll(tenantID uint) ([]Branch, error)
	Update(id uint, tenantID uint, b *Branch) (*Branch, error)
	Delete(id uint, tenantID uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Create(b *Branch) (*Branch, error) {
	if err := s.repo.Create(b); err != nil {
		return nil, err
	}
	return b, nil
}

func (s *service) GetByID(id uint, tenantID uint) (*Branch, error) {
	return s.repo.GetByID(id, tenantID)
}

func (s *service) GetAll(tenantID uint) ([]Branch, error) {
	return s.repo.GetAll(tenantID)
}

func (s *service) Update(id uint, tenantID uint, b *Branch) (*Branch, error) {
	if err := s.repo.Update(id, tenantID, b); err != nil {
		return nil, err
	}
	return s.repo.GetByID(id, tenantID)
}

func (s *service) Delete(id uint, tenantID uint) error {
	return s.repo.Delete(id, tenantID)
}
