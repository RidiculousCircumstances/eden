package domain

import "time"

type Profile struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	SourceID   uint      `gorm:"not null;uniqueIndex:idx_source_external"`
	CityID     uint      `gorm:"index"`
	Name       string    `gorm:"size:255"`
	URL        string    `gorm:"size:512;not null"`
	ExternalID string    `gorm:"size:255;not null;uniqueIndex:idx_source_external"`
	Gender     int       `gorm:"default:0"`
	BirthDate  string    `gorm:"size:255"`
	PersonID   uint      `gorm:"index"`
	CreatedAt  time.Time `gorm:"autoCreateTime;index"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	Photos     []*Photo  `gorm:"foreignKey:ProfileID;references:ID"`
}
