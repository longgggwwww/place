package dtos

import (
	"time"
	"user/models"

	"gorm.io/gorm"
)

// ------ PERMISSION GROUP ----------

type CreateGroup struct {
	ID        uint           `json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Name      string         `json:"name" binding:"required"`
	Perms     []CreatePerm   `json:"permissions" gorm:"foreignKey:GroupID"`
}

func (CreateGroup) TableName() string {
	return models.PermGrp{}.TableName()
}

type UpdateGroup struct {
	Name  string       `json:"name"`
	Perms []CreatePerm `json:"permissions" gorm:"foreignKey:GroupID"`
}

func (UpdateGroup) TableName() string {
	return models.PermGrp{}.TableName()
}

type DeleteGroup struct {
	IDs []uint `json:"ids" binding:"required"`
}

// ------ PERMISSION ----------

type CreatePerm struct {
	ID          uint           `json:"-"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-"`
	Name        string         `json:"name" binding:"required"`
	Description string         `json:"description"`
	GroupID     uint           `json:"group_id" binding:"required"`
	Group       *CreateGroup   `json:"group"`
}

func (CreatePerm) TableName() string {
	return models.Perm{}.TableName()
}
