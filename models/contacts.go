package models

import (
	"ginchat/utils"

	"gorm.io/gorm"
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

func SearchFriends(userId uint) []UserBasic {
	return Search(userId, Friend)
}

func Search(userId uint, ctype int) []UserBasic {
	contacts := make([]Contacts, 0)
	utils.Db.Where("owner_id = ? and type = ?", userId, ctype).Find(&contacts)

	uids := make([]uint, 0)
	for _, v := range contacts {
		uids = append(uids, v.TargetId)
	}
	users := make([]UserBasic, 0)
	utils.Db.Where("id in (?)", uids).Find(&users)
	return users
}

func AddFriend(userId uint, targetId uint) error {
	return AddContact(userId, targetId, Friend)
}

func AddContact(userId uint, targetId uint, contactType int) error {
	tx := utils.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := utils.Db.Create(&Contacts{
		OwnerId:  userId,
		TargetId: targetId,
		Type:     contactType,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if contactType == Friend {
		if err := utils.Db.Create(&Contacts{
			OwnerId:  targetId,
			TargetId: userId,
			Type:     contactType,
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func ContactsOfCommunity(communityId uint) []Contacts {
	contacts := make([]Contacts, 0)
	utils.Db.Where("target_id = ? and type = ?", communityId, GROUP).Find(&contacts)
	return contacts
}
