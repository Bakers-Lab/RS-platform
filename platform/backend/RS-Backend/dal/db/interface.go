package db

import (
	"context"

	"gorm.io/gorm"
)

type IDB interface {
	Read(context.Context) *gorm.DB
	Write(context.Context) *gorm.DB
	Transaction(context.Context, func(*gorm.DB) error) error
	Close()
}

type IDBResolver interface {
	IDB
}
