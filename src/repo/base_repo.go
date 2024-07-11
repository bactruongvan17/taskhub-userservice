package repo

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"time"

	"gorm.io/gorm"
)

const (
	StateNew byte = iota + 1 // starts from 1
	StateDoing
	StateDone

	generalQueryTimeout         = 60 * time.Second
	generalQueryTimeout2Minutes = 120 * time.Second
	defaultPageSize             = 30
	maxPageSize                 = 1000
)

type RepoPG struct {
	db *gorm.DB
}

func NewPGRepo(db *gorm.DB) PGInterface {
	return &RepoPG{db: db}
}

type PGInterface interface {
	DBWithTimeout(ctx context.Context) (*gorm.DB, context.CancelFunc)
	DB() (db *gorm.DB)
	Transaction(ctx context.Context, f func(rp PGInterface) error) error
}

func (r *RepoPG) Transaction(ctx context.Context, f func(rp PGInterface) error) (err error) {
	// log := logger.WithCtx(ctx, "RepoPG.Transaction")
	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()
	// create new instance to run the transaction
	repo := *r
	tx = tx.Begin()
	repo.db = tx
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			err = errors.New(fmt.Sprint(r))
			// log.WithError(err).Error("error_500: Panic when run Transaction")
			debug.PrintStack()
			return
		}
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	err = f(&repo)
	if err != nil {
		// log.WithError(err).Error("error_500: Error when run Transaction")
		return err
	}
	return nil
}

func (r *RepoPG) DB() *gorm.DB {
	return r.db
}

func (r *RepoPG) DBWithTimeout(ctx context.Context) (*gorm.DB, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(ctx, generalQueryTimeout)
	return r.db.WithContext(ctx), cancel
}
