package repo

import (
	"context"

	"github.com/bactruongvan17/taskhub-userservice/src/pkg/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *RepoPG) CreateUser(ctx context.Context, model *model.User, tx *gorm.DB) error {
	var cancel context.CancelFunc
	if tx == nil {
		tx, cancel = r.DBWithTimeout(ctx)
		defer cancel()
	}

	if err := tx.Debug().Create(model).Error; err != nil {
		return err
	}

	return nil
}

func (r *RepoPG) GetUserByEmail(ctx context.Context, email string, tx *gorm.DB) (rs *model.User, err error) {
	var cancel context.CancelFunc
	if tx == nil {
		tx, cancel = r.DBWithTimeout(ctx)
		defer cancel()
	}

	if err = tx.Where("email = ?", email).First(&rs).Error; err != nil {
		return nil, err
	}

	return rs, nil
}

func (r *RepoPG) GetUserById(ctx context.Context, userId uuid.UUID, tx *gorm.DB) (rs *model.User, err error) {
	var cancel context.CancelFunc
	if tx == nil {
		tx, cancel = r.DBWithTimeout(ctx)
		defer cancel()
	}

	if err = tx.First(&rs, userId).Error; err != nil {
		return nil, err
	}

	return rs, nil
}
