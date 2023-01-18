package v1

import (
	"TikSound-backend/store"
)

type Service interface {
	Users() UserSrv
	Videos() VideoSrv
	Comments() CommentSrv
}

type service struct {
	store store.Factory
}

func (s *service) Users() UserSrv {
	return newUserService(s)
}

func (s *service) Videos() VideoSrv {
	return newVideoService(s)
}

func (s *service) Comments() CommentSrv {
	return newComment(s)
}

func NewService(store store.Factory) Service {
	return &service{
		store: store,
	}
}
