package user_branch

import "time"

type UserBranch struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;column:id_user_branch"`
	UserID    uint      `json:"user_id" gorm:"column:user_id;not null;index"`
	BranchID  uint      `json:"branch_id" gorm:"column:branch_id;not null;index"`
	Status    bool      `json:"status" gorm:"type:boolean;default:true"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at;type:timestamptz;autoUpdateTime"`
}

func (UserBranch) TableName() string { return "user_branch" }
