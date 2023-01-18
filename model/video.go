package model

type Video struct {
	Uuid         uint64 `json:"uuid"`
	Author       string `json:"author"`
	PlayUrl      string `json:"playUrl"`
	CoverUrl     string `json:"coverUrl"`
	FavCount     uint16 `json:"favCount"`
	CommentCount uint16 `json:"commentCount"`
}
