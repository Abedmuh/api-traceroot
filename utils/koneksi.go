package utils

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBConnection() (*gorm.DB, error) {

	viper.SetConfigFile(".env")
	viper.AddConfigPath("../")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("Error reading .env file: %v\n", err)
			log.Println("Switching to environment variables...")
			viper.AutomaticEnv()
		} else {
			log.Fatalf("file found but error: %v\n", err)
		}
	}

	dbUser := viper.GetString("DB_USERNAME")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbName := viper.GetString("DB_NAME")
	dbParams := viper.GetString("DB_PARAMS")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbParams)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Ping the database to ensure connection is established
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}