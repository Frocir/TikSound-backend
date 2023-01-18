package model

type User struct {
	*UserInfo
}

type UserInfo struct {
	Uuid     uint64 `json:"uuid"`
	NickName string `json:"nickName"`
	Mail     string `json:"mail"`
	Passwd   string `json:"passwd"`
	Token    string `json:"token"`
}
type Follower struct {
}
