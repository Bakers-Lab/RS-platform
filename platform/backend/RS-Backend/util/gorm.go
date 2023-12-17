package util

import (
	"errors"

	"gorm.io/gorm"
)

func GormRealError(tx *gorm.DB) error {
	if tx.Error == nil {
		return nil
	}
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return tx.Error
}
