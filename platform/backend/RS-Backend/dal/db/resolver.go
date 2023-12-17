package db

import (
	"context"

	"gorm.io/gorm"
)

type defaultWriteDBResolver struct {
	IDB
	useRead bool
}

func NewDefaultWriteDBResolver(db IDB) IDBResolver {
	return &defaultWriteDBResolver{IDB: db, useRead: false}
}

func (r *defaultWriteDBResolver) UseWrite() IDB {
	if !r.useRead {
		return r
	}
	return &defaultWriteDBResolver{IDB: r.IDB, useRead: false}
}

func (r *defaultWriteDBResolver) UseRead() IDB {
	if r.useRead {
		return r
	}
	return &defaultWriteDBResolver{IDB: r.IDB, useRead: true}
}

func (r *defaultWriteDBResolver) Read(ctx context.Context) *gorm.DB {
	if r.useRead {
		return r.IDB.Read(ctx)
	}
	return r.IDB.Write(ctx)
}
