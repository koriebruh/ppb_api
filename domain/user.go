package domain

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	User     string `gorm:"unique;not null" json:"user"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Role     string `gorm:"type:ENUM('admin', 'staff', 'konsumen');not null;default:'konsumen'" json:"role"`
}
