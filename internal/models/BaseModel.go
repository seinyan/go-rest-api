package models

import (
	"time"
)

type BaseModel struct {
	//*gorm.Model
	ID        uint64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
