package models

import "gorm.io/gorm"

type GroupBasic struct {
	gorm.Model
	Name    string
	OwnerId uint
	Avatar  string
	Desc    string
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}
