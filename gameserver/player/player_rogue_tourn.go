package player

import (
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

func (g *GamePlayer) RogueTournQueryCsReq(payloadMsg pb.Message) {
	rsp := &proto.RogueTournQueryScRsp{
		Retcode:           0,
		RogueTournCurInfo: g.GetRogueTournCurInfo(),
		RogueGetInfo: &proto.RogueTournInfo{
			RogueTournSaveList:       make([]*proto.RogueTournSaveList, 0),
			RogueTournAreaInfo:       g.GetRogueTournAreaInfo(),
			PermanentInfo:            g.GetInspirationCircuitInfo(),
			RogueSeasonInfo:          g.GetRogueTournSeasonInfo(),
			ExtraScoreInfo:           g.GetExtraScoreInfo(),
			RogueTournExpInfo:        g.GetRogueTournExpInfo(),
			RogueTournHandbook:       g.GetRogueTournHandbookInfo(),
			RogueTournDifficultyInfo: g.GetRogueTournDifficultyInfo(),
		},
	}
	g.Send(cmd.RogueTournQueryScRsp, rsp)
}

func (g *GamePlayer) RogueTournGetPermanentTalentInfoCsReq(payloadMsg pb.Message) {
	rsp := &proto.RogueTournGetPermanentTalentInfoScRsp{
		PermanentInfo: g.GetInspirationCircuitInfo(),
		Retcode:       0,
	}
	g.Send(cmd.RogueTournGetPermanentTalentInfoScRsp, rsp)
}

func (g *GamePlayer) RogueTournGetMiscRealTimeDataCsReq(payloadMsg pb.Message) {
	rsp := &proto.RogueTournGetMiscRealTimeDataScRsp{}
	g.Send(cmd.RogueTournGetMiscRealTimeDataScRsp, rsp)
}

func (g *GamePlayer) RogueTournStartCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.RogueTournStartCsReq)
	rsp := new(proto.RogueTournStartScRsp)
	conf := gdconf.GetRogueTournAreaById(req.AreaId)
	if conf == nil {
		g.Send(cmd.RogueTournStartScRsp, rsp)
		return
	}
	// 更新队伍
	g.SetBattleLineUp(RogueTourn, req.BaseAvatarIdList)
	// 更新db
	g.SetBattleStatus(spb.BattleType_Battle_ROGUE_TOURN)
	g.SetMaterialById(Cf, 100) // 将宇宙碎片重置成100个
	g.NewCurRogueTourn(req.AreaId)
	curRoom := g.GetCurRogueTournRoom()

	rsp.RogueTournCurSceneInfo = &proto.RogueTournCurSceneInfo{
		Lineup:     g.GetLineUpPb(g.GetBattleLineUpById(RogueTourn)),
		RotateInfo: g.GetRogueMapRotateInfo(curRoom.RoomId),
		Scene:      g.GetRogueTournScene(curRoom.RoomId),
	}
	rsp.RogueTournCurInfo = g.GetRogueTournCurInfo()

	// 选择三个初始方程
	g.AddRogueBuffNum()
	g.FormulaSyncRogueCommonPendingActionScNotify([]uint32{130204, 130408, 130307})
	rsp.RogueTournCurInfo.RogueTournCurAreaInfo.PendingAction = &proto.RogueCommonPendingAction{
		QueuePosition: g.GetRogueBuffNum(),
		RogueAction: &proto.RogueAction{
			Action: &proto.RogueAction_RogueFormulaSelectInfo{
				RogueFormulaSelectInfo: &proto.RogueFormulaSelectInfo{
					SelectFormulaIdList: []uint32{130204, 130408, 130307},
				},
			},
		},
	}

	g.Send(cmd.RogueTournStartScRsp, rsp)
}

func (g *GamePlayer) RogueTournEnterCsReq(payloadMsg pb.Message) {
	curRoom := g.GetCurRogueTournRoom()
	rsp := &proto.RogueTournEnterScRsp{
		RogueTournCurInfo: g.GetRogueTournCurInfo(),
		RogueTournCurSceneInfo: &proto.RogueTournCurSceneInfo{
			Lineup:     g.GetLineUpPb(g.GetBattleLineUpById(RogueTourn)),
			RotateInfo: g.GetRogueMapRotateInfo(curRoom.RoomId),
			Scene:      g.GetRogueTournScene(curRoom.RoomId),
		},
	}

	g.Send(cmd.RogueTournEnterScRsp, rsp)
}

func (g *GamePlayer) RogueTournSettleCsReq(payloadMsg pb.Message) {
	rsp := &proto.RogueTournSettleScRsp{
		Retcode: 0,
		// IOLFDOIPNKA:            nil,
	}
	db := g.GetRogueTourn()
	db.CurRogueTourn = nil
	g.SetBattleStatus(spb.BattleType_Battle_NONE)
	g.Send(cmd.RogueTournSettleScRsp, rsp)
}

func (g *GamePlayer) RogueTournEnterRoomCsReq(payloadMsg pb.Message) {
	req := payloadMsg.(*proto.RogueTournEnterRoomCsReq)
	g.UpdateRogueTournEnterRoom(req.CurRoomIndex, req.NextRoomType)
	curRoom := g.GetCurRogueTournRoom()
	rsp := &proto.RogueTournEnterRoomScRsp{
		RogueTournCurSceneInfo: &proto.RogueTournCurSceneInfo{
			Lineup:     g.GetLineUpPb(g.GetBattleLineUpById(RogueTourn)),
			RotateInfo: g.GetRogueMapRotateInfo(curRoom.RoomId),
			Scene:      g.GetRogueTournScene(curRoom.RoomId),
		},
		Retcode: 0,
	}

	g.Send(cmd.RogueTournEnterRoomScRsp, rsp)
}

// 房间切换通知
func (g *GamePlayer) RogueTournLevelInfoUpdateScNotify(layerIndex, roomIndex uint32) {
	db := g.GetCurRogueTourn()
	layerInfo := db.CurLayerList[layerIndex]
	roomInfo := layerInfo.RogueTournRoomList[roomIndex]
	notify := &proto.RogueTournLevelInfoUpdateScNotify{
		CurLevelIndex: layerIndex,
		Status:        proto.RogueTournLevelStatus(layerInfo.Status),
		LevelInfoList: make([]*proto.RogueTournLevel, 0),
	}
	notify.LevelInfoList = append(notify.LevelInfoList, &proto.RogueTournLevel{
		LayerId:      layerInfo.LayerId,
		Status:       proto.RogueTournLayerStatus(layerInfo.Status),
		CurRoomIndex: layerInfo.CurRoomIndex,
		TournRoomList: []*proto.RogueTournRoomList{
			{
				Status:    proto.RogueTournRoomStatus(roomInfo.Status),
				RoomIndex: roomInfo.RoomIndex,
				RoomId:    roomInfo.RoomId,
			},
		},
		LevelIndex: layerInfo.LayerIndex,
	})

	g.Send(cmd.RogueTournLevelInfoUpdateScNotify, notify)
}

// 方程选择通知
func (g *GamePlayer) FormulaSyncRogueCommonPendingActionScNotify(formulaList []uint32) {
	notify := &proto.SyncRogueCommonPendingActionScNotify{
		RogueSubMode: 301,
		Action: &proto.RogueCommonPendingAction{
			QueuePosition: g.GetRogueBuffNum(),
			RogueAction: &proto.RogueAction{
				Action: &proto.RogueAction_RogueFormulaSelectInfo{
					RogueFormulaSelectInfo: &proto.RogueFormulaSelectInfo{
						SelectFormulaIdList: formulaList,
					},
				},
			},
		},
	}

	g.Send(cmd.SyncRogueCommonPendingActionScNotify, notify)
}

// 方程选择后通知
func (g *GamePlayer) FormulaSyncRogueCommonActionResultScNotify(formulaId uint32) {
	conf := gdconf.GetRogueTournFormulaById(formulaId)
	if conf == nil {
		return
	}
	notify := &proto.SyncRogueCommonActionResultScNotify{
		ActionResult: make([]*proto.RogueCommonActionResult, 0),
		RogueSubMode: 301,
	}
	notify.ActionResult = append(notify.ActionResult, &proto.RogueCommonActionResult{
		Source: 0,
		RogueAction: &proto.RogueCommonActionResultData{
			ResultData: &proto.RogueCommonActionResultData_GetFormulaList{
				GetFormulaList: &proto.RogueCommonFormula{
					FormulaInfo: &proto.FormulaInfo{
						IsExpand:  false,
						FormulaId: formulaId,
						FormulaBuffTypeList: []*proto.FormulaBuffTypeInfo{
							{
								FormulaBuffNum: conf.MainBuffNum,
								Key:            conf.MainBuffTypeID,
							},
							{
								FormulaBuffNum: conf.SubBuffNum,
								Key:            conf.SubBuffTypeID,
							},
						},
					},
				},
			},
		},
	})

	g.Send(cmd.SyncRogueCommonActionResultScNotify, notify)
}
