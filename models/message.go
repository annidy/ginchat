package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	FromId   uint
	TargetId uint
	Type     int // 0: p2p, 1: group
	Media    int // 0: text, 1: image
	Content  string
}

func (m *Message) TableName() string {
	return "message"
}
