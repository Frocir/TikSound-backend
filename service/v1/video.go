package v1

import "TikSound-backend/store"

type videoService struct {
	store.Factory
}
type VideoSrv interface {
}

var _ VideoSrv = (*videoService)(nil)

func newVideoService(s *service) VideoSrv {
	return &videoService{s.store}
}
