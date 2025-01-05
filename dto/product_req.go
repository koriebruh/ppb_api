package dto

type ProdukRequest struct {
	Name        string  `validate:"required" json:"name"`
	Description string  `validate:"required" json:"description"`
	Price       float64 `validate:"required" json:"price"`
	ImageUrl    string  `validate:"required" json:"imageUrl"`
}
