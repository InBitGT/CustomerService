package branch

import "gorm.io/gorm"

type Repository interface {
	Create(b *Branch) error
	GetByID(id uint, tenantID uint) (*Branch, error)
	GetAll(tenantID uint) ([]Branch, error)
	Update(id uint, tenantID uint, b *Branch) error
	Delete(id uint, tenantID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(b *Branch) error {
	return r.db.Create(b).Error
}

func (r *repository) GetByID(id uint, tenantID uint) (*Branch, error) {
	var b Branch
	if err := r.db.Where("id_branch = ? AND tenant_id = ?", id, tenantID).First(&b).Error; err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *repository) GetAll(tenantID uint) ([]Branch, error) {
	var list []Branch
	if err := r.db.Where("tenant_id = ?", tenantID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *repository) Update(id uint, tenantID uint, b *Branch) error {
	return r.db.Model(&Branch{}).
		Where("id_branch = ? AND tenant_id = ?", id, tenantID).
		Updates(b).Error
}

func (r *repository) Delete(id uint, tenantID uint) error {
	return r.db.Where("id_branch = ? AND tenant_id = ?", id, tenantID).Delete(&Branch{}).Error
}
