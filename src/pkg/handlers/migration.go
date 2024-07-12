package handlers

import (
	"github.com/bactruongvan17/taskhub-userservice/src/pkg/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MigrationHandler struct {
	db *gorm.DB
}

func NewMigrationHandler(db *gorm.DB) *MigrationHandler {
	return &MigrationHandler{
		db: db,
	}
}

func (h *MigrationHandler) Migrate(ctx *gin.Context) {
	_ = h.db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	models := []interface{}{
		model.User{},
	}

	for _, m := range models {
		err := h.db.AutoMigrate(m)
		if err != nil {
			_ = ctx.Error(err)
			return
		}
	}
}
