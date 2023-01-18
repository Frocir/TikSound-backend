package mysql

import (
	"TikSound-backend/model"
	"context"
	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func newUser(ds *dataStore) *user {
	return &user{ds.db}
}

func (u *user) Create(ctx context.Context, user *model.User) error {

	return nil
}
func (u *user) Update(ctx context.Context, user *model.User) error {

	return nil
}
func (u *user) Query(ctx context.Context, user *model.User) error {

	return nil
}
