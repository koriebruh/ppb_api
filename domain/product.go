package domain

type Product struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string  `gorm:"not null" json:"name"`
	Description string  `json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
	ImageUrl    string  `gorm:"text" json:"imageUrl"`
}
