package player

import (
	"strconv"

	"github.com/gucooing/hkrpg-go/gameserver/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *GamePlayer) DressRelicAvatarCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.DressRelicAvatarCsReq, payloadMsg)
	req := msg.(*proto.DressRelicAvatarCsReq)

	g.DressRelicAvatarPlayerSyncScNotify(req.BaseAvatarId, req.ParamList)
	rsp := new(proto.GetChallengeScRsp)
	// TODO 是的，没错，还是同样的原因
	g.Send(cmd.DressAvatarScRsp, rsp)
}

func (g *GamePlayer) DressRelicAvatarPlayerSyncScNotify(avatarId uint32, paramList []*proto.RelicParam) {
	notify := &proto.PlayerSyncScNotify{
		AvatarSync: &proto.AvatarSync{AvatarList: make([]*proto.Avatar, 0)},
		RelicList:  make([]*proto.Relic, 0),
	}
	avatardb := g.PlayerPb.Avatar.Avatar[avatarId]
	// 是否已被装备
	for _, relic := range paramList {
		relicdb := g.GetItem().RelicMap[relic.RelicUniqueId]
		if relicdb == nil {
			return
		}
		if relicdb.BaseAvatarId != 0 {
			// 进入交换
			avatarDbs := g.PlayerPb.Avatar.Avatar[relicdb.BaseAvatarId]
			if avatardb.EquipRelic[relic.Slot] == 0 {
				delete(avatarDbs.EquipRelic, relic.Slot)
			} else {
				avatarDbs.EquipRelic[relic.Slot] = avatardb.EquipRelic[relic.Slot]
				g.GetItem().RelicMap[avatardb.EquipRelic[relic.Slot]].BaseAvatarId = avatarDbs.AvatarId

				relicList := g.GetProtoRelicById(avatardb.EquipRelic[relic.Slot])
				notify.RelicList = append(notify.RelicList, relicList)
			}
			avatar := g.GetProtoAvatarById(avatarDbs.AvatarId)
			notify.AvatarSync.AvatarList = append(notify.AvatarSync.AvatarList, avatar)
		}

		if avatardb.EquipRelic[relic.Slot] != 0 {
			oldEquipRelic := g.GetItem().RelicMap[avatardb.EquipRelic[relic.Slot]]
			if oldEquipRelic != nil {
				if oldEquipRelic.BaseAvatarId == avatarId {
					oldEquipRelic.BaseAvatarId = 0
					relicList := g.GetProtoRelicById(avatardb.EquipRelic[relic.Slot])
					notify.RelicList = append(notify.RelicList, relicList)
				}
			}
		}
		relicdb.BaseAvatarId = avatarId
		if avatardb.EquipRelic == nil {
			avatardb.EquipRelic = make(map[uint32]uint32)
		}
		avatardb.EquipRelic[relic.Slot] = relic.RelicUniqueId

		relicList := g.GetProtoRelicById(relic.RelicUniqueId)
		notify.RelicList = append(notify.RelicList, relicList)
	}

	avatar := g.GetProtoAvatarById(avatarId)
	notify.AvatarSync.AvatarList = append(notify.AvatarSync.AvatarList, avatar)

	g.Send(cmd.PlayerSyncScNotify, notify)
}

func (g *GamePlayer) ExpUpRelicCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.ExpUpRelicCsReq, payloadMsg)
	req := msg.(*proto.ExpUpRelicCsReq)
	if req.RelicUniqueId == 0 {
		rsp := &proto.ExpUpRelicScRsp{}
		g.Send(cmd.ExpUpRelicScRsp, rsp)
		return
	}

	var relicList []uint32   // 需要删除的relicList
	var pileItem []*Material // 需要删除的升级材料
	var delScoin uint32      // 扣除的信用点
	var addExp uint32        // 增加的经验
	var oldLevel uint32      // 升级前等级

	// 从背包获取需要升级的圣遗物
	dbRelic := g.getRelicDbById(req.RelicUniqueId)
	if dbRelic == nil {
		rsp := &proto.ExpUpRelicScRsp{}
		g.Send(cmd.ExpUpRelicScRsp, rsp)
		return
	}
	oldLevel = dbRelic.Level
	// 获取需要升级圣遗物的配置信息
	relicConf := gdconf.GetRelicById(strconv.Itoa(int(dbRelic.Tid)))
	if relicConf == nil {
		rsp := &proto.ExpUpRelicScRsp{}
		g.Send(cmd.ExpUpRelicScRsp, rsp)
		return
	}

	// 遍历用来升级的材料
	for _, pileList := range req.ItemCostList.ItemList {
		// 如果没有则退出
		if pileList.GetPileItem() == nil {
			continue
		}
		pile := new(Material)
		pile.Tid = pileList.GetPileItem().ItemId
		pile.Num = pileList.GetPileItem().ItemNum

		pileItem = append(pileItem, pile)
		// 获取材料配置
		pileconf := gdconf.GetRelicById(strconv.Itoa(int(pileList.GetPileItem().ItemId)))
		if pileconf == nil {
			rsp := &proto.ExpUpRelicScRsp{}
			g.Send(cmd.ExpUpRelicScRsp, rsp)
			return
		}
		// 获取要扣多少信用点
		delScoin += pileconf.CoinCost * pileList.GetPileItem().ItemNum
		// 获取能添加多少经验
		addExp += pileconf.ExpProvide * pileList.GetPileItem().ItemNum
	}

	// 遍历用来升级的光锥
	for _, relic := range req.ItemCostList.ItemList {
		// 如果没有则退出
		if relic.GetRelicUniqueId() == 0 {
			continue
		}
		relicList = append(relicList, relic.GetRelicUniqueId())
		// 获取光锥配置
		relicconfig := gdconf.GetRelicById(strconv.Itoa(int(g.GetProtoRelicById(relic.GetRelicUniqueId()).Tid)))
		if relicconfig == nil {
			rsp := &proto.ExpUpRelicScRsp{}
			g.Send(cmd.ExpUpRelicScRsp, rsp)
			return
		}
		// 获取要扣多少信用点
		delScoin += relicconfig.CoinCost
		// 获取能添加多少经验
		addExp += relicconfig.ExpProvide
	}

	// 计算添加后有多少经验
	exp := addExp + dbRelic.Exp

	// 获取能升级到的等级和升级后经验
	level, exp := gdconf.GetRelicExpByLevel(relicConf.ExpType, exp, dbRelic.Level, dbRelic.Tid)
	if level == 0 && exp == 0 {
		rsp := &proto.ExpUpRelicScRsp{}
		g.Send(cmd.ExpUpRelicScRsp, rsp)
		return
	}

	// 添加副属性
	addSubAffixes := 0
	for i := oldLevel; i <= level; i++ {
		if i%3 == 0 {
			addSubAffixes++
		}
	}
	if oldLevel%3 == 0 {
		addSubAffixes--
	}
	g.addRelicAffix(&addRelicAffix{
		addSubAffixes:     addSubAffixes, // int((level - oldLevel + 2) / 3),
		mainAffixProperty: dbRelic.MainAffixProperty,
		subAffixGroup:     relicConf.SubAffixGroup,
		relicAffix:        dbRelic.RelicAffix,
	})
	// 扣除本次升级需要的信用点
	pileItem = append(pileItem, &Material{
		Tid: 2,
		Num: delScoin,
	})
	// 更新需要升级的圣遗物状态
	dbRelic.Level = level
	dbRelic.Exp = exp

	// 删除用来升级的材料
	if len(pileItem) != 0 {
		g.DelMaterial(pileItem)
	}
	if len(relicList) != 0 {
		// 删除用来升级的圣遗物
		g.DelRelicPlayerSyncScNotify(relicList)
	}
	// 通知角色还有多少信用点
	g.PlayerPlayerSyncScNotify()
	// 通知升级后圣遗物消息
	g.RelicPlayerSyncScNotify(req.RelicUniqueId)
	rsp := &proto.ExpUpRelicScRsp{}
	g.Send(cmd.ExpUpRelicScRsp, rsp)
}

func (g *GamePlayer) DelRelicPlayerSyncScNotify(relicList []uint32) {
	for _, relic := range relicList {
		delete(g.GetItem().RelicMap, relic)
	}

	notify := &proto.PlayerSyncScNotify{DelRelicList: relicList}
	g.Send(cmd.PlayerSyncScNotify, notify)
}
