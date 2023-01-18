package store

import (
	"TikSound-backend/model"
	"context"
)

type CommentStore interface {
	Create(context.Context, *model.Comment) error
	Update(context.Context, *model.Comment) error
	Query(context.Context, *model.Comment) error
}
