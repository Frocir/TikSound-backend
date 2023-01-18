package user

import (
	"TikSound-backend/service/v1"
	"TikSound-backend/store"
)

type UserController struct {
	srv v1.Service
}

func NewUserController(s store.Factory) *UserController {
	return &UserController{
		srv: v1.NewService(s),
	}
}
