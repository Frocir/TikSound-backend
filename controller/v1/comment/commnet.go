package comment

import (
	"TikSound-backend/service/v1"
	"TikSound-backend/store"
)

type CommentController struct {
	srv v1.Service
}

func NewCommentController(s store.Factory) *CommentController {
	return &CommentController{
		srv: v1.NewService(s),
	}
}
