package models

import (
	"fmt"
	"ginchat/utils"

	"gorm.io/gorm"
)

type Community struct {
	gorm.Model
	Name    string
	OwnerId uint
	Icon    string
	Desc    string
	Memo    string
	Cate    int
}

func (table *Community) TableName() string {
	return "community"
}

func CreateCommunity(community *Community) error {
	if len(community.Name) < 3 {
		return fmt.Errorf("community name too short")
	}
	tx := utils.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := utils.Db.Create(community).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := AddContact(community.OwnerId, community.ID, GROUP); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func LoadCommunity(ownerId uint) ([]Community, error) {
	communities := make([]Community, 0)
	err := utils.Db.Where("owner_id = ?", ownerId).Find(&communities).Error
	return communities, err
}
