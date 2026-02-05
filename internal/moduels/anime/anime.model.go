package anime

import "time"

type Anime struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Status      string `gorm:"default:ongoing"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
