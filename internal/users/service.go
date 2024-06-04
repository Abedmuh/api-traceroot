package users

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// interface
type UserSvcInter interface {
	CreateUser(req ReqUserReg, tx *sql.DB, ctx *gin.Context) (ResUser, error)
	LoginUser(useDb User, req ReqUserLog, tx *sql.DB, ctx *gin.Context) (ResUser, error)

	// checking
	CheckUserReg(req string, tx *sql.DB, ctx *gin.Context) error
	CheckUserLog(req string, tx *sql.DB, ctx *gin.Context) (User, error)
}

// implementations
type UserSvcImpl struct {
}

func NewUserService() UserSvcInter {
	return &UserSvcImpl{}
}

func (u *UserSvcImpl) CreateUser(req ReqUserReg, tx *sql.DB, ctx *gin.Context) (ResUser, error) {
	return ResUser{}, nil
}

func (u *UserSvcImpl) LoginUser(useDb User, req ReqUserLog, tx *sql.DB, ctx *gin.Context) (ResUser, error) {
	return ResUser{}, nil
}

func (u *UserSvcImpl) CheckUserReg(req string, tx *sql.DB, ctx *gin.Context) error {
	return nil
}

func (u *UserSvcImpl) CheckUserLog(req string, tx *sql.DB, ctx *gin.Context) (User, error) {
	return User{}, nil
}
