package dto

type AddToCartRequest struct {
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	Jumlah    int  `json:"jumlah"`
}

type GetCartItemsRequest struct {
	UserID uint `json:"user_id"`
}

type AddShippingRequest struct {
	UserID      uint    `json:"user_id"`
	KotaAsal    string  `json:"kota_asal"`
	KotaTujuan  string  `json:"kota_tujuan"`
	BiayaOngkir float64 `json:"biaya_ongkir"`
	Weight      float64 `json:"weight"`
}

type CheckoutRequest struct {
	UserID uint `json:"user_id"`
	IsPaid bool `json:"is_paid"`
}

type RemoveFromCartRequest struct {
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
}
