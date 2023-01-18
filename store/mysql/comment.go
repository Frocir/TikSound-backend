package mysql

import (
	"TikSound-backend/model"
	"context"
	"gorm.io/gorm"
)

type comment struct {
	db *gorm.DB
}

func newComment(ds *dataStore) *comment {
	return &comment{ds.db}
}
func (c *comment) Create(ctx context.Context, cmt *model.Comment) error {

	return nil
}
func (c *comment) Update(ctx context.Context, cmt *model.Comment) error {

	return nil
}
func (c *comment) Query(ctx context.Context, cmt *model.Comment) error {

	return nil
}
