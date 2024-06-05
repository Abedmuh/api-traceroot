package users

import "github.com/golang-jwt/jwt/v5"

//main model
type User struct {
	Id         string `json:"id"`
	Email      string `json:"email"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Password   string `json:"password"`
	Perusahaan string `json:"perusahaan"`
	No_tlpn    string `json:"no_tlpn"`
	CreatedAt  string `json:"created_at"`
}

// request
type Users struct {
	Email      string `json:"email" validate:"required,email"`
	First_name string `json:"first_name" validate:"required,min=5,max=50"`
	Last_name  string `json:"last_name" validate:"required,min=5,max=50"`
	Password   string `json:"password" validate:"required,min=5,max=15"`
	Perusahaan string `json:"perusahaan" validate:"required"`
	Jabatan    string `json:"jabatan" validate:"required"`
	No_telpn   string `json:"no_telpn" validate:"required"`
}
type ReqUserLog struct {
	Email    string `json:"email" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=15"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// response
type ResUser struct {
	Email       string `json:"email"`
	AccessToken string `json:"accessToken"`
}
