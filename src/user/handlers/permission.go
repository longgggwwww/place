package handlers

import (
	"log"
	"net/http"
	"strconv"
	"user/common"
	"user/dtos"
	"user/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Group struct {
	Repo   *gorm.DB
	Logger *log.Logger
}

func (g *Group) Create(c *gin.Context) {
	// Check input
	var dto dtos.CreateGroup
	if err := c.ShouldBind(&dto); err != nil {
		c.JSON(http.StatusBadRequest, common.Error{
			Msg:  "Invalid input",
			Errs: err.Error(),
		})
		return
	}

	// Create record
	if err := g.Repo.Create(&dto).Error; err != nil {
		c.JSON(http.StatusBadRequest, common.Error{
			Msg:  "Create item failed",
			Errs: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, dto.ID)
}

func (g *Group) Find(c *gin.Context) {
	var groups []models.PermGrp
	if err := g.Repo.Preload("Perms").Find(&groups).Error; err != nil {
		c.JSON(http.StatusNotFound, common.Error{
			Msg:  "Not found",
			Errs: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, groups)
}

func (g *Group) FindById(c *gin.Context) {
	// Validate ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.Error{
			Msg:  "Invalid ID",
			Errs: err.Error(),
		})
		return
	}

	// Find record
	var group models.PermGrp
	if err := g.Repo.Preload("Perms").First(&group, id).Error; err != nil {
		c.JSON(http.StatusNotFound, common.Error{
			Msg:  "Not found",
			Errs: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, group)
}

func (g *Group) Update(c *gin.Context) {
	// Validate ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.Error{
			Msg:  "Invalid ID",
			Errs: err.Error(),
		})
		return
	}

	// Validate input
	var dto dtos.UpdateGroup
	if err := c.ShouldBind(&dto); err != nil {
		c.JSON(http.StatusBadRequest, common.Error{
			Msg:  "Invalid input",
			Errs: err.Error(),
		})
		return
	}

	// Update record
	if err := g.Repo.Where("id = ?", id).Updates(&dto).Error; err != nil {
		c.JSON(http.StatusBadRequest, common.Error{
			Msg:  "Update item failed",
			Errs: err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}

func (g *Group) Delete(c *gin.Context) {
	// Check ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.Error{
			Msg:  "Invalid ID",
			Errs: err.Error(),
		})
		return
	}

	// Delete record
	if err := g.Repo.Delete(&models.PermGrp{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, common.Error{
			Msg:  "Delete item failed",
			Errs: err.Error(),
		})
		return
	}
	c.Status(http.StatusNoContent)
}

func (g *Group) DeleteMany(c *gin.Context) {
	// Check input
	var dto dtos.DeleteGroup
	if err := c.ShouldBind(&dto); err != nil {
		c.JSON(http.StatusBadRequest, common.Error{
			Msg:  "Invalid input",
			Errs: err.Error(),
		})
		return
	}

	// Delete records
	if err := g.Repo.Delete(&models.PermGrp{}, "id IN ?", dto.IDs).Error; err != nil {
		c.JSON(http.StatusBadRequest, common.Error{
			Msg:  "Delete item failed",
			Errs: err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}

type Perm struct {
	Repo   *gorm.DB
	Logger *log.Logger
}

func (p *Perm) Create(c *gin.Context) {
	// Check input
	var dto dtos.CreatePerm
	if err := c.ShouldBind(&dto); err != nil {
		c.JSON(http.StatusBadRequest, common.Error{
			Msg:  "Invalid input",
			Errs: err.Error(),
		})
		return
	}

	// Create record
	if err := p.Repo.Create(&dto).Error; err != nil {
		c.JSON(http.StatusBadRequest, common.Error{
			Msg:  "Create item failed",
			Errs: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, dto.ID)
}

func (p *Perm) Find(c *gin.Context) {
	var perm []models.Perm
	if err := p.Repo.Preload("Group").Find(&perm).Error; err != nil {
		c.JSON(http.StatusNotFound, common.Error{
			Msg:  "Not Found",
			Errs: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, perm)
}
