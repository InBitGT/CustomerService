package address

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(a *Address) error
	GetByID(id uint) (*Address, error)
	GetAll() ([]Address, error)
	Update(id uint, a *Address) error
	Delete(id uint) error     // soft delete (status=false)
	HardDelete(id uint) error // físico
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(a *Address) error {
	return r.db.Create(a).Error
}

func (r *repository) GetByID(id uint) (*Address, error) {
	var a Address
	if err := r.db.
		Where("id_address = ? AND status = true", id).
		First(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *repository) GetAll() ([]Address, error) {
	var list []Address
	if err := r.db.
		Where("status = true").
		Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *repository) Update(id uint, a *Address) error {
	updates := map[string]interface{}{
		"line1":       a.Line1,
		"line2":       a.Line2,
		"city":        a.City,
		"state":       a.State,
		"country":     a.Country,
		"postal_code": a.PostalCode,
		"status":      a.Status,
	}

	res := r.db.Model(&Address{}).
		Where("id_address = ?", id).
		Updates(updates)

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// ✅ Soft delete: status=false
func (r *repository) Delete(id uint) error {
	res := r.db.Model(&Address{}).
		Where("id_address = ? AND status = true", id).
		Updates(map[string]interface{}{
			"status": false,
		})

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// ✅ Hard delete: físico
func (r *repository) HardDelete(id uint) error {
	return r.db.Unscoped().Where("id_address = ?", id).Delete(&Address{}).Error
}
