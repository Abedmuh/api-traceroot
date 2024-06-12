package utils

import "github.com/go-playground/validator/v10"


func CommandValidator(fl validator.FieldLevel) bool {
	var allowedCommands = []string{
		"ping", 
		"traceroute",
	}

	command := fl.Field().String()
	for _, allowedCommand := range allowedCommands {
		if command == allowedCommand {
			return true
		}
	}
	return false
}