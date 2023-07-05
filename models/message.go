package models

import (
	"gorm.io/gorm"
)

const (
	_ = iota
	Friend
	GROUP
)

type Message struct {
	gorm.Model
	FromId   uint
	TargetId uint
	Type     int // 0: friend, 1: group, 3: 心跳
	Media    int // 0: text, 1: image
	Content  string
}

func (m *Message) TableName() string {
	return "message"
}
