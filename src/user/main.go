package main

import (
	"log"
	"user/handlers"
	"user/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	l := log.Default()

	dsn := "host=localhost user=root password=1234 dbname=user port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		l.Fatal("Failed to connect Postgres")
	}

	// Migration
	if err := db.AutoMigrate(&models.PermGrp{}, &models.Perm{}, &models.Role{}); err != nil {
		l.Fatal(err.Error())
	}

	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		perm := v1.Group("permissions")
		{
			// api/v1/permissions/group
			groupH := handlers.Group{
				Repo:   db,
				Logger: l,
			}
			perm.POST("group", groupH.Create)
			perm.GET("group", groupH.Find)
			perm.GET("group/:id", groupH.FindById)
			perm.PATCH("group/:id", groupH.Update)
			perm.DELETE("group", groupH.DeleteMany)
			perm.DELETE("group/:id", groupH.Delete)

			// api/v1/permissions
			permH := handlers.Perm{
				Repo:   db,
				Logger: l,
			}
			perm.POST("", permH.Create)
			perm.GET("", permH.Find)
		}
	}

	r.Run() // :8080
}
