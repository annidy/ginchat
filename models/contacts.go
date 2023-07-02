package models

import "gorm.io/gorm"

const (
	P2P = iota
	GROUP
)

type Contacts struct {
	gorm.Model
	OwnerId  uint
	TargetId uint
	Type     int
}

func (table *Contacts) TableName() string {
	return "contacts"
}
