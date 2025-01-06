package conf

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
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

	//if err = DB.AutoMigrate(
	//	&domain.Product{},
	//	&domain.User{},
	//	&domain.Cart{},
	//	&domain.CartItem{},
	//	&domain.Shipping{},
	//); err != nil {
	//	log.Fatalf("Error fail migrate: %v", err)
	//}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get *sql.DB from GORM: %v", err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(15 * time.Minute)

	log.Println("connection aws establish")

	return DB
}
