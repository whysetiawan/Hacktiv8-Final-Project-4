package config

import (
	"final-project-4/httpserver/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, error) {
	dbCredential := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", os.Getenv("PGHOST"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"), os.Getenv("PGPORT"))

	db, err := gorm.Open(postgres.Open(dbCredential), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&models.UserModel{},
		&models.CategoryModel{},
		&models.ProductModel{},
		&models.TransactionHistoryModel{},
	)
	return db, err
}
