package domain

import (
	"gorm.io/gorm"
	"time"
)

type Cart struct {
	gorm.Model
	UserID    uint       `gorm:"not null"`
	Status    string     `gorm:"not null"` // Status bisa "aktif" atau "selesai"
	Tanggal   time.Time  `gorm:"not null"`
	User      User       `gorm:"foreignKey:UserID"`
	CartItems []CartItem `gorm:"foreignKey:CartID"`
	Shipping  *Shipping  `gorm:"foreignKey:CartID"`
	Total     float64    `gorm:"not null"`
}
