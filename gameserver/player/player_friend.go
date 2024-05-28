package player

import (
	"time"

	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *GamePlayer) HandleGetFriendLoginInfoCsReq(payloadMsg []byte) {
	rsp := &proto.GetFriendLoginInfoScRsp{
		FriendUidList: g.GetFriendList(),
		Retcode:       0,
	}
	g.Send(cmd.GetFriendLoginInfoScRsp, rsp)
}

func (g *GamePlayer) GetFriendListInfoCsReq(payloadMsg []byte) {
	rsp := new(proto.GetFriendListInfoScRsp)
	rsp.FriendList = make([]*proto.FriendSimpleInfo, 0)
	for _, uid := range g.GetFriendList() {
		simpleInfo := g.GetPlayerSimpleInfo(uid)
		if simpleInfo == nil {
			continue
		}
		rsp.FriendList = append(rsp.FriendList, &proto.FriendSimpleInfo{PlayerInfo: simpleInfo})
	}
	g.Send(cmd.GetFriendListInfoScRsp, rsp)
}

func (g *GamePlayer) GetPlayerDetailInfoCsReq(payloadMsg []byte) {
	msg := g.DecodePayloadToProto(cmd.GetPlayerDetailInfoCsReq, payloadMsg)
	req := msg.(*proto.GetPlayerDetailInfoCsReq)
	g.Send(cmd.GetPlayerDetailInfoScRsp, &proto.GetPlayerDetailInfoScRsp{
		DetailInfo: g.GetPlayerDetailInfo(req.Uid),
		Retcode:    0,
	})
}

func (g *GamePlayer) GetFriendApplyListInfoCsReq(payloadMsg []byte) {
	receiveApplyList := g.GetRecvApplyFriend()
	rsp := &proto.GetFriendApplyListInfoScRsp{
		SendApplyList:    make([]uint32, 0),
		Retcode:          0,
		ReceiveApplyList: make([]*proto.FriendApplyInfo, 0),
	}
	for _, receiveApply := range receiveApplyList {
		bin := g.GetFriendApplyInfo(receiveApply)
		if bin == nil {
			continue
		}
		rsp.SendApplyList = append(rsp.SendApplyList, receiveApply.ApplyUid)
		rsp.ReceiveApplyList = append(rsp.ReceiveApplyList, bin)
	}
	g.Send(cmd.GetFriendApplyListInfoScRsp, rsp)
}

func (g *GamePlayer) GetChatFriendHistoryCsReq(payloadMsg []byte) {
	g.Send(cmd.GetChatFriendHistoryScRsp, &proto.GetChatFriendHistoryScRsp{
		FriendHistoryInfo: []*proto.FriendHistoryInfo{
			{
				LastSendTime: time.Now().Unix(),
				ContactId:    0,
			},
		},
		Retcode: 0,
	})
}