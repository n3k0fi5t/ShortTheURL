package shorter

import "time"

type ShortUrl struct {
	ID        uint `gorm:"primarykey"`
	URL       string
	CreateAt  time.Time
	UpdatedAt time.Time
}
