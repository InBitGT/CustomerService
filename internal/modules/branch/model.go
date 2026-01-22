package branch

import "time"

type Branch struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement;column:id_branch"`
	Name        string    `json:"name" gorm:"type:varchar(100);not null"`
	AddressID   uint      `json:"address_id" gorm:"column:address_id"`
	Description string    `json:"description" gorm:"type:text"`
	TenantID    uint      `json:"tenant_id" gorm:"column:tenant_id;not null;index"`
	Status      bool      `json:"status" gorm:"type:boolean;default:true"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt   time.Time `json:"update_at" gorm:"column:update_at;type:timestamptz;autoUpdateTime"`
}

func (Branch) TableName() string { return "branch" }
