package domain

import "time"

type Photo struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	ProfileID uint      `gorm:"index"`
	IndexID   uint32    `gorm:"index;unique"`
	URL       string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Faces     []*Face   `gorm:"foreignKey:PhotoID;references:ID"`
}

type Face struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	PhotoID uint `gorm:"index"`
	Age     int
	Sex     int
	Bbox    string `gorm:"size:512"`
}
