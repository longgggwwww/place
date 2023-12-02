package handlers

import (
	"log"
	"net/http"
	"user/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Role struct {
	Repo   *gorm.DB
	Logger *log.Logger
}

func (r *Role) Find(c *gin.Context) {
	var roles []models.Role
	if err := r.Repo.Find(&roles).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Roles Not Found",
		})
		return
	}
	c.JSON(http.StatusOK, roles)
}

func (r *Role) FindById(c *gin.Context) {
	id := c.Param("id")
	var role Role
	if err := r.Repo.Preload("Perms").First(id, &role).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "Role Not Found",
		})
		return
	}
	c.JSON(http.StatusOK, role)
}

func (r *Role) Create(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBind(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":    "Validation Failed",
			"errors": err.Error(),
		})
		return
	}
	if err := r.Repo.Preload("Perms").Create(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":    "Create Role Failed",
			"errors": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, role)
}
