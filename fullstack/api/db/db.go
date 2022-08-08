package db

import (
	"fmt"
	"log"

	"github.com/denizedizcan/final-project-denizedizcan-propertyfinder/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// init the db
func Init() *gorm.DB {
	dbUrl := "postgres://postgres:mysecretpassword@localhost:5432/postgres"

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to %s database", postgres.Open(dbUrl))
		log.Fatalln(err)
	}
	db.AutoMigrate(
		&models.Product{},
		&models.User{},
		&models.BasketItems{},
		&models.Price{},
		&models.Stock{},
		&models.Order{},
		&models.OrderItems{},
		&models.Basket{},
	)
	return db
}
