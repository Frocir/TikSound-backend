package store

import (
	"TikSound-backend/model"
	"context"
)

type UserStore interface {
	Create(context.Context, *model.User) error
	Update(context.Context, *model.User) error
	Query(context.Context, *model.User) error
}
