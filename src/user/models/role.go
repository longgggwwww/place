package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name      string         `json:"name" binding:"required"`
	Level     *uint          `json:"level" binding:"required"`
	Perms     []Perm         `json:"permissions" gorm:"many2many:roles_perms" binding:"omitempty"`
}

func (Role) TableName() string {
	return "roles"
}
