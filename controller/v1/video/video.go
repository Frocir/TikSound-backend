package video

import (
	v1 "TikSound-backend/service/v1"
	"TikSound-backend/store"
)

type VideoController struct {
	srv v1.Service
}

func NewCommentController(s store.Factory) *VideoController {
	return &VideoController{
		srv: v1.NewService(s),
	}
}
