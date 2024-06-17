package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *GamePlayer) DressRelicAvatarCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.DressRelicAvatarCsReq, payloadMsg)
	req := msg.(*proto.DressRelicAvatarCsReq)
	g.DressRelicAvatarPlayerSyncScNotify(req.GetDressAvatarId(), req.GetSwitchList())
	g.Send(cmd.DressRelicAvatarScRsp, nil)
}

func (g *GamePlayer) DressRelicAvatarPlayerSyncScNotify(equipAvatarId uint32, paramList []*proto.DressRelicParam) {
	if paramList == nil {
		return
	}
	notify := &proto.PlayerSyncScNotify{
		AvatarSync: &proto.AvatarSync{AvatarList: make([]*proto.Avatar, 0)},
		RelicList:  make([]*proto.Relic, 0),
	}

	equipAvatarDb := g.GetAvatarBinById(equipAvatarId)
	for _, relic := range paramList {
		relicDb := g.GetRelicById(relic.RelicUniqueId)
		if relicDb == nil {
			continue
		}
		baseAvatarDb := g.GetAvatarBinById(relicDb.BaseAvatarId)
		relicDb.BaseAvatarId = equipAvatarId
		if equipAvatarDb != nil {
			oldRelicDb := g.GetAvatarEquipRelic(equipAvatarId, relic.RelicType)
			if oldRelicDb != nil {
				oldRelicDb.BaseAvatarId = 0
				notify.RelicList = append(notify.RelicList, g.GetProtoRelicById(oldRelicDb.UniqueId))
			}
			g.SetAvatarEquipRelic(equipAvatarId, relic.RelicType, relic.RelicUniqueId)
			notify.AvatarSync.AvatarList = append(notify.AvatarSync.AvatarList, g.GetProtoAvatarById(equipAvatarId))
		}
		if baseAvatarDb != nil {
			g.SetAvatarEquipRelic(baseAvatarDb.AvatarId, relic.RelicType, 0)
			notify.AvatarSync.AvatarList = append(notify.AvatarSync.AvatarList, g.GetProtoAvatarById(baseAvatarDb.AvatarId))
		}
		notify.RelicList = append(notify.RelicList, g.GetProtoRelicById(relic.RelicUniqueId))
	}
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
	dbRelic := g.GetRelicById(req.RelicUniqueId)
	if dbRelic == nil {
		rsp := &proto.ExpUpRelicScRsp{}
		g.Send(cmd.ExpUpRelicScRsp, rsp)
		return
	}
	oldLevel = dbRelic.Level
	// 获取需要升级圣遗物的配置信息
	relicConf := gdconf.GetRelicById(dbRelic.Tid)
	if relicConf == nil {
		rsp := &proto.ExpUpRelicScRsp{}
		g.Send(cmd.ExpUpRelicScRsp, rsp)
		return
	}

	// 遍历用来升级的材料
	for _, pileList := range req.GetCostData().ItemList {
		// 如果没有则退出
		if pileList.GetPileItem() == nil {
			continue
		}
		pileItem = append(pileItem, &Material{
			Tid: pileList.GetPileItem().ItemId,
			Num: pileList.GetPileItem().ItemNum,
		})
		// 获取材料配置
		pileconf := gdconf.GetRelicById(pileList.GetPileItem().ItemId)
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
	for _, relic := range req.GetCostData().ItemList {
		// 如果没有则退出
		if relic.GetRelicUniqueId() == 0 {
			continue
		}
		relicList = append(relicList, relic.GetRelicUniqueId())
		// 获取光锥配置
		relicconfig := gdconf.GetRelicById(g.GetProtoRelicById(relic.GetRelicUniqueId()).Tid)
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
