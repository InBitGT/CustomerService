package address

type Service interface {
	Create(a *Address) (*Address, error)
	GetByID(id uint) (*Address, error)
	GetAll() ([]Address, error)
	Update(id uint, a *Address) (*Address, error)
	Delete(id uint) error
	HardDelete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Create(a *Address) (*Address, error) {
	if err := s.repo.Create(a); err != nil {
		return nil, err
	}
	return a, nil
}

func (s *service) GetByID(id uint) (*Address, error) {
	return s.repo.GetByID(id)
}

func (s *service) GetAll() ([]Address, error) {
	return s.repo.GetAll()
}

func (s *service) Update(id uint, a *Address) (*Address, error) {
	if err := s.repo.Update(id, a); err != nil {
		return nil, err
	}
	return s.repo.GetByID(id)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *service) HardDelete(id uint) error {
	return s.repo.HardDelete(id)
}
