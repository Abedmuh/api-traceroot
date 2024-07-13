package icmp

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
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

	servers = []Server{
		{
			Name:     "jakarta",
			Host:  	  viper.GetString("SSH_JAKARTA_HOST"),
			Username: viper.GetString("SSH_JAKARTA_USERNAME"),
			Password: viper.GetString("SSH_JAKARTA_PASSWORD"),
		},
		{
			Name:     "bandung",
			Host:     "190.xxx.xx.xx",
			Username: viper.GetString("SSH_BANDUNG_USERNAME"),
			Password: viper.GetString("SSH_BANDUNG_PASSWORD"),
		},
	}
}