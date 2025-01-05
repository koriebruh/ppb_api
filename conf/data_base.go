package conf

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"koriebruh/uas-ppb/domain"
	"log"
	"log/slog"
)

func NewConnection() *gorm.DB {
	config := NewConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DataBase.User,
		config.DataBase.Pass,
		config.DataBase.Host,
		config.DataBase.Port,
		config.DataBase.Name,
	)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	if err = DB.AutoMigrate(
		&domain.Product{},
		&domain.User{},
		&domain.Cart{},
		&domain.CartItem{},
		&domain.Shipping{},
	); err != nil {
		log.Fatalf("Error fail migrate: %v", err)
	}

	slog.Info("connection aws establish")

	return DB
}
