package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromId   uint
	TargetId uint
	Type     int
	Content  string
}

func (m *Message) TableName() string {
	return "message"
}
