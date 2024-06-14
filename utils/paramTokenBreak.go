package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetTokenEmail(ctx *gin.Context) (string, error) {
	reqMail, exists := ctx.Get("email")
	if !exists {
		return "", errors.New("email not found in context")
	}

	email, ok := reqMail.(string)
	if !ok {
		return "", errors.New("email ID is not a string")
	}

	return email, nil
}