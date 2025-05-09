package models

import (
	"gorm.io/gorm"
	"time"
)

type CommonField struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt" gorm:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func GormPaginate(page, pageSize int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 1000:
			pageSize = 1000
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}
