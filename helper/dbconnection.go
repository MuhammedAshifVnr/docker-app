package helper

import (
	"ecom/database"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func EnvLoader() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("-------------Faild to load env file------------ ")
	}
}

func DbConnect() {

	dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(".....................Connection faild..................")
	}

	DB = db

	DB.AutoMigrate(&database.Wallet{}, &database.OrderItems{}, &database.Order{}, &database.User{}, &database.Admin{}, &database.Category{}, &database.Otp{})
	DB.AutoMigrate(&database.Offers{}, &database.Transactions{}, &database.Whislist{}, &database.Coupon{}, &database.Cart{}, &database.Address{}, &database.Product{}, &database.Review{})
	fmt.Println("SUCCESSFULLY connected to DATABASE")

}
