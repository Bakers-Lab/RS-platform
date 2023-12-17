package db

import (
	"RS-Backend/util"
	"context"
	"errors"
	"runtime/debug"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	dbImpl struct {
		db *gorm.DB
	}
	DBConf struct {
	}
)

func (d *dbImpl) Read(ctx context.Context) *gorm.DB {
	return d.db.Session(&gorm.Session{QueryFields: true, PrepareStmt: true, Context: ctx})
}

func (d *dbImpl) Write(ctx context.Context) *gorm.DB {
	return d.db.Session(&gorm.Session{QueryFields: true, Context: ctx})
}

func (d *dbImpl) Transaction(ctx context.Context, f func(tx *gorm.DB) error) error {
	return d.Write(ctx).Transaction(func(tx *gorm.DB) (e error) {
		defer func() {
			if r := recover(); r != nil {
				logrus.Error("DB Transaction executed with error, err = %v, stack=%v", r, string(debug.Stack()))
				e = errors.New("DB error")
			}
		}()
		return f(tx)
	})
}

func (d *dbImpl) Close() {
	db, _ := d.db.DB()
	if db != nil {
		_ = db.Close()
	}
}

func NewDB(dsn string) IDB {
	dial := postgres.Open(dsn)
	util.PanicIf(dial == nil, "dial is nil")
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	util.PanicIf(err != nil, "Failed to connect to the database")
	return &dbImpl{db: db}
}
