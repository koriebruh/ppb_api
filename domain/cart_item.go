package domain

import (
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	CartID    uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Jumlah    int     `gorm:"not null"`
	Harga     float64 `gorm:"not null"`
	Subtotal  float64 `gorm:"not null"` // Harga x Jumlah
	Product   Product `gorm:"foreignKey:ProductID"`
	Cart      Cart    `gorm:"foreignKey:CartID"`
}
