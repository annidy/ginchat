package models

import (
	"fmt"
	"ginchat/utils"

	"gorm.io/gorm"
)

type Community struct {
	gorm.Model
	Name    string `gorm:"primarykey"`
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

func LoadCommunity(userId uint) ([]Community, error) {
	contacts := make([]Contacts, 0)
	if err := utils.Db.Where("owner_id = ?", userId).Find(&contacts).Error; err != nil {
		return nil, err
	}
	communityIDs := make([]uint, 0)
	for _, v := range contacts {
		communityIDs = append(communityIDs, v.TargetId)
	}
	communities := make([]Community, 0)
	if err := utils.Db.Where("id in (?)", communityIDs).Find(&communities).Error; err != nil {
		return nil, err
	}
	return communities, nil
}

func LoadOwnCommunity(ownerId uint) ([]Community, error) {
	communities := make([]Community, 0)
	err := utils.Db.Where("owner_id = ?", ownerId).Find(&communities).Error
	return communities, err
}

func JoinGroup(userId uint, name string) error {
	comm := Community{}
	if err := utils.Db.Where("name = ?", name).First(&comm).Error; err != nil {
		return err
	}
	fmt.Println("JoinGroup", comm)

	if err := utils.Db.Where("owner_id = ? AND target_id = ? AND type = 2", userId, comm.ID).First(&Contacts{}).Error; err == nil {
		return fmt.Errorf("user has joined this community")
	}

	return AddContact(userId, comm.ID, GROUP)
}
