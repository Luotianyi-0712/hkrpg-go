package player

import (
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

const MaxLineupList = 20 // 设置最大普通队伍数

func (g *GamePlayer) GetLineUp() *spb.LineUp {
	db := g.GetPlayerPb()
	if db.LineUp == nil {
		db.LineUp = &spb.LineUp{
			MainLineUp:     0,
			LineUpList:     make(map[uint32]*spb.Line),
			BattleLineList: make(map[uint32]*spb.BattleLine),
		}
		db.LineUp.LineUpList[0] = &spb.Line{
			Name:         "hkrpg",
			AvatarIdList: make(map[uint32]*spb.LineAvatarList),
			LeaderSlot:   0,
		}
		db.LineUp.LineUpList[0].AvatarIdList[0] = &spb.LineAvatarList{AvatarId: uint32(g.GetAvatar().CurMainAvatar), Slot: 0}
	}
	return db.LineUp
}

func (g *GamePlayer) GetLineUpById(index uint32) *spb.Line {
	db := g.GetLineUp()
	if db.LineUpList == nil {
		db.LineUpList = make(map[uint32]*spb.Line)
		db.LineUpList[0] = &spb.Line{
			Name:         "hkrpg",
			AvatarIdList: make(map[uint32]*spb.LineAvatarList),
			LeaderSlot:   0,
		}
		db.LineUpList[0].AvatarIdList[0] = &spb.LineAvatarList{AvatarId: uint32(g.GetAvatar().CurMainAvatar), Slot: 0}
	}
	if db.LineUpList[index] == nil {
		db.LineUpList[index] = &spb.Line{
			Name:         "",
			AvatarIdList: make(map[uint32]*spb.LineAvatarList),
			LeaderSlot:   0,
		}
	}
	return db.LineUpList[index]
}

func (g *GamePlayer) GetBattleLineUpById(index uint32) *spb.BattleLine {
	db := g.GetLineUp()
	if db.BattleLineList == nil {
		db.BattleLineList = make(map[uint32]*spb.BattleLine)
	}
	if db.BattleLineList[index] == nil {
		db.BattleLineList[index] = &spb.BattleLine{
			AvatarIdList:    make(map[uint32]*spb.LineAvatarList),
			LeaderSlot:      0,
			ExtraLineupType: spb.ExtraLineupType(index),
		}
	}
	return db.BattleLineList[index]
}

func (g *GamePlayer) GetCurLineUp() *spb.Line {
	db := g.GetLineUp()
	return db.LineUpList[db.MainLineUp]
}

func (g *GamePlayer) GetSceneAvatarId() uint32 {
	db := g.GetCurLineUp()
	if db.AvatarIdList[db.LeaderSlot] == nil {
		logger.Error("获取的角色不存在于队伍中")
		return 8001
	} else {
		return db.AvatarIdList[db.LeaderSlot].AvatarId
	}
}

// 队伍更新
func (g *GamePlayer) UnDbLineUp(index uint32, Slot uint32, avatarId uint32) {
	db := g.GetLineUpById(index)
	db.AvatarIdList[Slot] = &spb.LineAvatarList{AvatarId: avatarId, Slot: Slot}
	db.LeaderSlot = Slot
}

// 交换角色
func (g *GamePlayer) SwapLineup(index, src_slot, dst_slot uint32) {
	db := g.GetLineUpById(index)
	src := db.AvatarIdList[src_slot]
	dst := db.AvatarIdList[dst_slot]
	src.Slot = dst.Slot
	dst.Slot = src.Slot
	db.AvatarIdList[src_slot] = dst
	db.AvatarIdList[dst_slot] = src
}

/*****************************************功能方法****************************/

func (g *GamePlayer) GetLineUpPb(id uint32) *proto.LineupInfo {
	db := g.GetLineUpById(id)
	lineupList := &proto.LineupInfo{
		IsVirtual:  false,
		LeaderSlot: db.LeaderSlot,
		AvatarList: make([]*proto.LineupAvatar, 0),
		Index:      id,
		MaxMp:      5,
		Mp:         5,
		Name:       db.Name,
		PlaneId:    0,
	}
	for slot, lineAvatar := range db.AvatarIdList {
		if lineAvatar == nil || lineAvatar.AvatarId == 0 {
			continue
		}
		avatarBin := g.GetAvatarBinById(lineAvatar.AvatarId)
		if avatarBin == nil {
			lineupAvatar := &proto.LineupAvatar{
				AvatarType: proto.AvatarType_AVATAR_TRIAL_TYPE,
				Slot:       slot,
				Satiety:    0,
				Hp:         10000,
				Id:         lineAvatar.AvatarId,
				SpBar: &proto.SpBarInfo{
					CurSp: 6000,
					MaxSp: 10000,
				},
			}
			lineupList.AvatarList = append(lineupList.AvatarList, lineupAvatar)
		} else {
			lineupAvatar := &proto.LineupAvatar{
				AvatarType: proto.AvatarType(avatarBin.AvatarType),
				Slot:       slot,
				Satiety:    0,
				Hp:         avatarBin.Hp,
				Id:         lineAvatar.AvatarId,
				SpBar: &proto.SpBarInfo{
					CurSp: avatarBin.SpBar.CurSp,
					MaxSp: avatarBin.SpBar.MaxSp,
				},
			}
			lineupList.AvatarList = append(lineupList.AvatarList, lineupAvatar)
		}
	}
	return lineupList
}
