package domain

import "time"

type Photo struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	ProfileID uint      `gorm:"index"`
	IndexID   uint      `gorm:"index;unique"`
	URL       string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
