package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

// 添加物品
func (g *GamePlayer) GmGive(payloadMsg pb.Message) {
	req := payloadMsg.(*spb.GmGive)
	allSync := &AllPlayerSync{
		IsBasic:       true,
		AvatarList:    make([]uint32, 0),
		MaterialList:  make([]uint32, 0),
		EquipmentList: make([]uint32, 0),
		RelicList:     make([]uint32, 0),
	}

	if req.GiveAll {
		g.AllGive(allSync)
	} else {
		allSync.MaterialList = append(allSync.MaterialList, req.ItemId)
		g.AddItem([]*Material{{
			Tid: req.ItemId,
			Num: req.ItemCount,
		}})
	}
	// 同步通知
	g.AllPlayerSyncScNotify(allSync)
}

func (g *GamePlayer) AllGive(allSync *AllPlayerSync) {
	var pileItem []*Material
	itemConf := gdconf.GetItemConfigMap()
	avatarConf := gdconf.GetAvatarDataMap()
	// add avatar
	for _, avatar := range avatarConf {
		x := avatar.AvatarId / 1000
		if x != 1 && x != 8 {
			continue
		}
		allSync.AvatarList = append(allSync.AvatarList, avatar.AvatarId)
		g.AddAvatar(avatar.AvatarId, proto.AddAvatarSrcState_ADD_AVATAR_SRC_GACHA)
	}
	// add playerIcon
	var playerIconList []uint32
	for _, playerIcon := range itemConf.AvatarPlayerIcon {
		playerIconList = append(playerIconList, playerIcon.ID)
	}
	g.GetItem().HeadIcon = playerIconList
	// add rank
	for _, rank := range itemConf.AvatarRank {
		allSync.MaterialList = append(allSync.MaterialList, rank.ID)
		pileItem = append(pileItem, &Material{
			Tid: rank.ID,
			Num: 6,
		})
	}
	// add equipment
	for _, equipment := range itemConf.Equipment {
		uniqueId := g.AddEquipment(equipment.ID)
		allSync.EquipmentList = append(allSync.EquipmentList, uniqueId)
	}
	// add item
	for _, item := range itemConf.Item {
		allSync.MaterialList = append(allSync.MaterialList, item.ID)
		pileItem = append(pileItem, &Material{
			Tid: item.ID,
			Num: 999999999,
		})
	}
	// add relic
	for _, relic := range itemConf.Relic {
		uniqueId := g.AddRelic(relic.ID)
		allSync.RelicList = append(allSync.RelicList, uniqueId)
	}
	// add bt relic
	for _, relic := range itemConf.Relic {
		uniqueId := g.AddBtRelic(relic.ID)
		allSync.RelicList = append(allSync.RelicList, uniqueId)
	}
	g.AddItem(pileItem)
}

// 设置世界等级
func (g *GamePlayer) GmWorldLevel(payloadMsg pb.Message) {
	req := payloadMsg.(*spb.GmWorldLevel)
	g.SetWorldLevel(req.WorldLevel)
	// 账号状态通知
	g.PlayerPlayerSyncScNotify()
}

// 清空背包
func (g *GamePlayer) DelItem(payloadMsg pb.Message) {
	db := g.GetItem()
	db = &spb.Item{
		RelicMap:     make(map[uint32]*spb.Relic),
		EquipmentMap: make(map[uint32]*spb.Equipment),
		MaterialMap:  make(map[uint32]uint32),
		HeadIcon:     make([]uint32, 0),
	}
	db.MaterialMap[11] = 240
}

// 角色一键满级
func (g *GamePlayer) GmMaxCurAvatar(payloadMsg pb.Message) {
	req := payloadMsg.(*spb.MaxCurAvatar)
	allSync := &AllPlayerSync{AvatarList: make([]uint32, 0)}
	if req.All {
		bin := g.GetAvatar()
		if bin == nil {
			return
		}
		for _, db := range bin.AvatarList {
			g.SetAvatarMaxByDb(db)
			allSync.AvatarList = append(allSync.AvatarList, db.AvatarId)
		}
	} else {
		var db *spb.AvatarBin
		db = g.GetAvatarBinById(req.AvatarId)
		if db == nil {
			db = g.GetCurAvatar()
		}
		allSync.AvatarList = append(allSync.AvatarList, db.AvatarId)
		g.SetAvatarMaxByDb(db)
	}
	g.AllPlayerSyncScNotify(allSync)
}

func (g *GamePlayer) SetAvatarMaxByDb(db *spb.AvatarBin) {
	if db == nil {
		return
	}
	db.Level = 80          // 80级
	db.PromoteLevel = 6    // 突破等级
	db.Hp = 10000          // 满血
	db.SpBar.CurSp = 10000 // 满能量
	for _, info := range db.MultiPathAvatarInfoList {
		info.Rank = 6                              // 六命
		for _, skill := range info.SkilltreeList { // 技能满级
			conf := gdconf.GetAvatarSkilltreeBySkillId(skill.PointId, 1)
			if conf == nil {
				continue
			}
			skill.Level = conf.MaxLevel
		}
	}
}

func (g *GamePlayer) RecoverLine() {
	db := g.GetCurLineUp()
	for _, a := range db.AvatarIdList {
		bin := g.GetAvatarById(a.AvatarId)
		if bin != nil {
			bin.Hp = 10000
			bin.SpBar.CurSp = 10000
		}
	}
	g.SyncLineupNotify(db)
}

func (g *GamePlayer) GmMission(req *spb.GmMission) {
	if req.FinishAll {
		g.FinishAllMission()
		g.FinishAllTutorial()
		return
	}
}

func (g *GamePlayer) FinishAllMission() {
	db := g.GetMainMission()
	db.SubMissionList = make(map[uint32]*spb.MissionInfo)
	db.MainMissionList = make(map[uint32]*spb.MissionInfo)
	for id, info := range gdconf.GetSubMainMission() {
		if db.FinishSubMissionList == nil {
			db.FinishSubMissionList = make(map[uint32]*spb.MissionInfo)
		}
		db.FinishSubMissionList[id] = &spb.MissionInfo{
			MissionId: id,
			Progress:  info.Progress,
			Status:    spb.MissionStatus_MISSION_FINISH,
		}
	}
	for id := range gdconf.GetGoppMainMission() {
		if db.FinishMainMissionList == nil {
			db.FinishMainMissionList = make(map[uint32]*spb.MissionInfo)
		}
		db.FinishMainMissionList[id] = &spb.MissionInfo{
			MissionId: id,
			Progress:  1,
			Status:    spb.MissionStatus_MISSION_FINISH,
		}
	}
}

func (g *GamePlayer) FinishAllTutorial() {
	tDb := g.GetTutorial()
	for id := range gdconf.GetTutorialData() {
		tDb[id] = &spb.TutorialInfo{
			Id:     id,
			Status: spb.TutorialStatus_TUTORIAL_FINISH,
		}
	}
	gDb := g.GetTutorialGuide()
	for id := range gdconf.GetTutorialGuideGroup() {
		gDb[id] = &spb.TutorialInfo{
			Id:     id,
			Status: spb.TutorialStatus_TUTORIAL_FINISH,
		}
	}
}
