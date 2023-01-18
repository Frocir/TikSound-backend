package v1

import "TikSound-backend/store"

type CommentService struct {
	store.Factory
}
type CommentSrv interface {
}

var _ CommentSrv = (*CommentService)(nil)

func newComment(s *service) CommentSrv {
	return &CommentService{s.store}
}
