package domain

import (
	"gorm.io/gorm"
)

type Shipping struct {
	gorm.Model
	CartID      uint    `gorm:"not null"`
	KotaAsal    string  `gorm:"not null"`
	KotaTujuan  string  `gorm:"not null"`
	BiayaOngkir float64 `gorm:"not null"` // Ongkir yang dihitung dari API eksternal
	Weight      float64 `gorm:"not null"` // Berat total barang
	Cart        Cart    `gorm:"foreignKey:CartID"`
}
