package v1

import "TikSound-backend/store"

type UserService struct {
	store.Factory
}
type UserSrv interface {
}

var _ UserSrv = (*UserService)(nil)

func newUserService(srv *service) UserSrv {
	return &UserService{srv.store}
}
