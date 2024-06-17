package db

import (
	"github.com/gucooing/hkrpg-go/pkg/database"
)

// 使用账号id拉取数据
func (s *Store) QueryAccountUidByFieldPlayer(uid uint32) *database.PlayerData {
	var playerData *database.PlayerData
	s.PlayerDataMysql.Model(&database.PlayerData{}).Where("uid = ?", uid).First(&playerData)
	return playerData
}

// 添加新账号数据
func (s *Store) AddDatePlayerFieldByFieldName(player *database.PlayerData) error {
	if err := s.PlayerDataMysql.Create(player).Error; err != nil {
		return err
	}
	return nil
}

// 更新账号
func (s *Store) UpdatePlayer(player *database.PlayerData) error {
	if player.Uid == 0 {
		return nil
	}
	if err := s.PlayerDataMysql.Model(&database.PlayerData{}).Where("uid = ?", player.Uid).Updates(player).Error; err == nil {
		return nil
	} else {
		return err
	}
}

// 拉取全部邮件
func (s *Store) GetAllMail() []*database.Mail {
	var mailMap []*database.Mail
	s.ServerConf.Find(&mailMap)
	return mailMap
}

// 拉取全部模拟宇宙
func (s *Store) GetAllRogue() []*database.RogueConf {
	var rogueMap []*database.RogueConf
	s.ServerConf.Find(&rogueMap)
	return rogueMap
}

// 拉取地图文件
func (s *Store) GetBlockData(uid, entryId uint32) *database.BlockData {
	var blockData *database.BlockData
	s.PlayerDataMysql.Where(&database.BlockData{Uid: uid, EntryId: entryId}).First(&blockData)
	return blockData
}

// 更新地图文件
func (s *Store) UpdateBlockData(blockData *database.BlockData) error {
	if blockData.Uid == 0 {
		return nil
	}
	if err := s.PlayerDataMysql.Save(blockData).Error; err == nil {
		return nil
	} else {
		return err
	}
}
