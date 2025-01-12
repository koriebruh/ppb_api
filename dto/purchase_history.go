package dto

import "time"

type PurchaseHistory struct {
	Name     string    `json:"name"`
	Jumlah   int       `json:"jumlah"`
	Price    float64   `json:"price"`
	Subtotal float64   `json:"subtotal"`
	BuyTime  time.Time `json:"buy_time"`
}
