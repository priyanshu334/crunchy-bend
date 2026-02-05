package image

import "time"

type Image struct {
	ID        uint   `gorm:"primaryKey"`
	URL       string `gorm:"not null"`
	OwnerType string `gorm:"index;not null"`
	OwnerID   uint   `gorm:"index;not null"`
	CreatedAt time.Time
}
