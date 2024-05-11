package player

import (
	"github.com/gucooing/hkrpg-go/gameserver/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) GmGive(payloadMsg pb.Message) {
	req := payloadMsg.(*spb.GmGive)
	if req.PlayerUid == 0 {
		return
	}
	itemConf := gdconf.GetItemConfigMap()
	if req.GiveAll {
		var pileItem []*Material
		// add avatar
		for _, avatar := range itemConf.Avatar {
			// 过滤非live角色
			if avatar.ID/1000 != 1 {
				continue
			}
			g.AddAvatar(avatar.ID)
		}
		// add playerIcon
		var playerIconList []uint32
		for _, playerIcon := range itemConf.AvatarPlayerIcon {
			playerIconList = append(playerIconList, playerIcon.ID)
		}
		g.GetItem().HeadIcon = playerIconList
		// add rank
		for _, rank := range itemConf.AvatarRank {
			pileItem = append(pileItem, &Material{
				Tid: rank.ID,
				Num: 6,
			})
		}
		// add equipment
		for _, equipment := range itemConf.Equipment {
			g.AddEquipment(equipment.ID)
		}
		// add item
		for _, item := range itemConf.Item {
			pileItem = append(pileItem, &Material{
				Tid: item.ID,
				Num: 99999,
			})
		}
		// add relic
		for _, relic := range itemConf.Relic {
			g.AddRelic(relic.ID)
		}
		g.AddMaterial(pileItem)
		// g.ScenePlaneEventScNotify(pileItem)
	} else {
		var pileItem []*Material
		for _, item := range itemConf.Item {
			if item.ID == req.ItemId {
				pileItem = append(pileItem, &Material{
					Tid: item.ID,
					Num: req.ItemCount,
				})
				g.AddMaterial(pileItem)
				return
			}
		}
		for _, avatar := range itemConf.Avatar {
			if avatar.ID == req.ItemId {
				g.AddAvatar(avatar.ID)
				return
			}
		}
		for _, avatar := range itemConf.AvatarRank {
			if avatar.ID == req.ItemId {
				pileItem = append(pileItem, &Material{
					Tid: avatar.ID,
					Num: req.ItemCount,
				})
				g.AddMaterial(pileItem)
				return
			}
		}
		for _, avatar := range itemConf.AvatarPlayerIcon {
			if avatar.ID == req.ItemId {
				g.AddHeadIcon(avatar.ID)
				return
			}
		}
		for _, equipment := range itemConf.Equipment {
			if equipment.ID == req.ItemId {
				g.AddEquipment(equipment.ID)
				return
			}
		}
		for _, relic := range itemConf.Relic {
			if relic.ID == req.ItemId {
				g.AddRelic(relic.ID)
				return
			}
		}
	}
}

func (g *GamePlayer) RelicScenePlaneEventScNotify(uniqueId uint32) {
	relicItme := g.GetProtoRelicById(uniqueId)
	// 通知客户端增加了物品
	notify := &proto.ScenePlaneEventScNotify{
		GetItemList: &proto.ItemList{
			ItemList: make([]*proto.Item, 0),
		},
	}
	item := &proto.Item{
		ItemId:      relicItme.Tid,
		Level:       relicItme.Level,
		Num:         1,
		MainAffixId: relicItme.MainAffixId,
		Rank:        0,
		Promotion:   0,
		UniqueId:    relicItme.UniqueId,
	}
	notify.GetItemList.ItemList = append(notify.GetItemList.ItemList, item)
	g.Send(cmd.ScenePlaneEventScNotify, notify)
}

func (g *GamePlayer) GmWorldLevel(payloadMsg pb.Message) {
	req := payloadMsg.(*spb.GmWorldLevel)

	g.PlayerPb.WorldLevel = req.WorldLevel

	// 账号状态通知
	g.PlayerPlayerSyncScNotify()
}

func (g *GamePlayer) DelItem(payloadMsg pb.Message) {
	g.PlayerPb.Item = &spb.Item{
		RelicMap:     make(map[uint32]*spb.Relic),
		EquipmentMap: make(map[uint32]*spb.Equipment),
		MaterialMap:  make(map[uint32]uint32),
		HeadIcon:     make([]uint32, 0),
	}
	g.PlayerPb.Item.MaterialMap[11] = 240
}
