package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func GetDBConnection() (*sql.DB, error) {

	viper.AutomaticEnv()
	// if err := viper.ReadInConfig(); err != nil {
	// 	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
	// 		log.Printf("Error reading .env file: %v\n", err)
	// 		log.Println("Switching to environment variables...")
	// 	} else {
	// 		log.Fatalf("file found but error: %v\n", err)
	// 	}
	// }

	// dbUser := viper.GetString("DB_USERNAME")
	// dbPassword := viper.GetString("DB_PASSWORD")
	// dbHost := viper.GetString("DB_HOST")
	// dbPort := viper.GetString("DB_PORT")
	// dbName := viper.GetString("DB_NAME")
	// dbParams := viper.GetString("DB_PARAMS")

	// connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
	// 	dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", "postgres://postgres:pass@localhost:5432/abhvps?sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %s", err)
	} else {
		fmt.Println("Connected to database")
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %s", err)
	}

	return db, nil
}
