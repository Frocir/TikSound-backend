package model

import "time"

type Comment struct {
	Uuid       uint64    `json:"uuid"`
	Content    string    `json:"content"`
	OwnerId    uint64    `json:"ownerId"`
	CreateDate time.Time `json:"createDate"`
}
