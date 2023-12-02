package models

import (
	"time"

	"gorm.io/gorm"
)

// PERMISSION GROUP
type PermGrp struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name      string         `json:"name"`
	Perms     []*Perm        `json:"permissions" gorm:"foreignKey:GroupID"`
}

func (PermGrp) TableName() string {
	return "permission-groups"
}

// PERMISSION
type Perm struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name      string         `json:"name" gorm:"unique_index"`
	Desc      string         `json:"description" gorm:"column:description"`
	GroupID   uint           `json:"group_id"`
	Group     PermGrp        `json:"group"`
}

func (Perm) TableName() string {
	return "permissions"
}
