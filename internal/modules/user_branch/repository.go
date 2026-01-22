package user_branch

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(tenantID uint, ub *UserBranch) error
	GetByID(tenantID uint, id uint) (*UserBranch, error)
	GetAll(tenantID uint) ([]UserBranch, error)
	Update(tenantID uint, id uint, ub *UserBranch) error
	Delete(tenantID uint, id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

// Nota: scoping por tenant se hace validando que el branch pertenezca al tenant
func (r *repository) Create(tenantID uint, ub *UserBranch) error {
	// Valida branch pertenece al tenant (tabla branch ya existe en este microservicio)
	var cnt int64
	if err := r.db.Table("branch").
		Where("id_branch = ? AND tenant_id = ?", ub.BranchID, tenantID).
		Count(&cnt).Error; err != nil {
		return err
	}
	if cnt == 0 {
		return gorm.ErrRecordNotFound
	}

	return r.db.Create(ub).Error
}

func (r *repository) GetByID(tenantID uint, id uint) (*UserBranch, error) {
	var ub UserBranch
	err := r.db.Table("user_branch ub").
		Select("ub.*").
		Joins("JOIN branch b ON b.id_branch = ub.branch_id").
		Where("ub.id_user_branch = ? AND b.tenant_id = ?", id, tenantID).
		First(&ub).Error
	if err != nil {
		return nil, err
	}
	return &ub, nil
}

func (r *repository) GetAll(tenantID uint) ([]UserBranch, error) {
	var list []UserBranch
	err := r.db.Table("user_branch ub").
		Select("ub.*").
		Joins("JOIN branch b ON b.id_branch = ub.branch_id").
		Where("b.tenant_id = ?", tenantID).
		Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *repository) Update(tenantID uint, id uint, ub *UserBranch) error {
	// Si están cambiando BranchID, validar que pertenezca al tenant
	if ub.BranchID != 0 {
		var cnt int64
		if err := r.db.Table("branch").
			Where("id_branch = ? AND tenant_id = ?", ub.BranchID, tenantID).
			Count(&cnt).Error; err != nil {
			return err
		}
		if cnt == 0 {
			return gorm.ErrRecordNotFound
		}
	}

	// Update scoping: solo update si pertenece al tenant via join
	res := r.db.Table("user_branch ub").
		Joins("JOIN branch b ON b.id_branch = ub.branch_id").
		Where("ub.id_user_branch = ? AND b.tenant_id = ?", id, tenantID).
		Updates(map[string]interface{}{
			"user_id":   ub.UserID,
			"branch_id": ub.BranchID,
			"status":    ub.Status,
		})

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *repository) Delete(tenantID uint, id uint) error {
	res := r.db.Table("user_branch ub").
		Joins("JOIN branch b ON b.id_branch = ub.branch_id").
		Where("ub.id_user_branch = ? AND b.tenant_id = ?", id, tenantID).
		Delete(&UserBranch{})

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
