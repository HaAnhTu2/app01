package reponsitory

import (
	"app02/models"
	"app02/models/req"
	"context"
)

type UserRepo interface {
	CreateUser(context context.Context, user models.User) (models.User, error)
	CheckLogin(context context.Context, loginReq req.ReqIn) (models.User, error)
}
