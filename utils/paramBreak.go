package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ParamBreakUser(ctx *gin.Context) (string, error) {
	user, err := ctx.Get("email")
	if !err {
		return "", fmt.Errorf( "unauthorized cant find user" )
	}
	reqUser, err := user.(string)
	if !err {
		return "", fmt.Errorf( "unauthorized cant find user" )
	}

	return reqUser, nil
}