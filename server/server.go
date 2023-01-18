package server

import (
	"TikSound-backend/config"
	"fmt"
	"net/http"
)

type Server struct {
	Srv  *http.Server // 服务
	Name string       // 服务名
	Addr string       // 监听端口
}

func NewServer(cfg *config.TkConfig) (*Server, error) {
	r, err := InitRouter(cfg)
	if err != nil {
		return nil, err
	}
	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Server.Addr),
		Handler: r,
	}
	myServer := &Server{
		Srv:  s,
		Name: cfg.Server.Name,
		Addr: cfg.Server.Addr,
	}
	return myServer, nil
}

func (s *Server) Run() {
	fmt.Printf("TikSound-backend server is running at port %v\n", s.Addr)
	err := s.Srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
