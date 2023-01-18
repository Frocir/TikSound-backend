package mysql

import (
	"TikSound-backend/model"
	"context"
	"gorm.io/gorm"
)

type video struct {
	db *gorm.DB
}

func newVideo(ds *dataStore) *video {
	return &video{ds.db}
}

func (v *video) Create(ctx context.Context, user *model.Video) error {
	return nil
}

func (v *video) Update(ctx context.Context, user *model.Video) error {

	return nil
}

func (v *video) Query(ctx context.Context, user *model.Video) error {

	return nil
}
