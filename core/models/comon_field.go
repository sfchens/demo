package models

import (
	"gorm.io/gorm"
	"time"
)

type CommonField struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
