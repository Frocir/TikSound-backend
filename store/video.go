package store

import (
	"TikSound-backend/model"
	"context"
)

type VideoStore interface {
	Create(context.Context, *model.Video) error
	Update(context.Context, *model.Video) error
	Query(context.Context, *model.Video) error
}
