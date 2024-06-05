package users

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// interface
type UserSvcInter interface {
	CreateUser(req Users, tx *gorm.DB, ctx *gin.Context) (ResUser, error)
	LoginUser(req ReqUserLog, tx *gorm.DB, ctx *gin.Context) (ResUser, error)
}

// implementations
type UserSvcImpl struct {
}

func NewUserService() UserSvcInter {
	return &UserSvcImpl{}
}

func (u *UserSvcImpl) CreateUser(req Users, tx *gorm.DB, ctx *gin.Context) (ResUser, error) {
	// Check if a user with the same email already exists
	var existingUser User
	if err := tx.Where("email =?", req.Email).First(&existingUser).Error; err == nil {
		return ResUser{}, errors.New("user with the same email already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return ResUser{}, err
	}
	// Check if a user with the same phone number already exists
	var existingPhoneNumber User
	if err := tx.Where("no_telpn = ?", req.No_telpn).First(&existingPhoneNumber).Error; err == nil {
		return ResUser{}, errors.New("user with the same phone number already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return ResUser{}, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return ResUser{}, err
	}
	req.Password = string(hashedPassword)
	if err := tx.Create(&req).Error; err != nil {
		return ResUser{}, err
	}
	return ResUser{}, nil
}

func (u *UserSvcImpl) LoginUser(req ReqUserLog, tx *gorm.DB, ctx *gin.Context) (ResUser, error) {
	// Check if a user with the given email exists
	var user User
	if err := tx.Where("email = ?", req.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ResUser{}, errors.New("email or password not found")
		}
		return ResUser{}, err
	}

	// Verify the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return ResUser{}, errors.New("email or password mismatch")
	}

	// Generate JWT token
	timeExp := viper.GetDuration("JWT_TIME_EXP")
	secretKey := viper.GetString("JWT_SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": req.Email,
		"exp":   time.Now().Add(time.Duration(timeExp) * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("Error signing")
		return ResUser{}, err
	}

	// Create the res
	res := ResUser{
		Email:       user.Email,
		AccessToken: tokenString,
	}

	return res, nil
}
