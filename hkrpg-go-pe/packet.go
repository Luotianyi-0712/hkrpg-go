package hkrpg_go_pe

import (
	"time"

	"github.com/gucooing/hkrpg-go/gameserver/player"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	pb "google.golang.org/protobuf/proto"
)

type packetFunc func(h *HkRpgGoServer, p *PlayerGame, protoMsg pb.Message)

var packetCaptureMap = map[uint16]packetFunc{ // 抽包
	cmd.PlayerLogoutCsReq: PlayerLogoutCsReq, // 下线请求
	cmd.SendMsgCsReq:      SendMsgCsReq,      // 发送聊天消息
	cmd.ApplyFriendCsReq:  ApplyFriendCsReq,  // 发起好友申请
}

func (h *HkRpgGoServer) packetCapture(p *PlayerGame, cmdId uint16, protoMsg pb.Message) {
	handelFunc, ok := packetCaptureMap[cmdId]
	if !ok {
		p.sendGameMsg(player.Client, cmdId, protoMsg)
		return
	}
	handelFunc(h, p, protoMsg)
}

func PlayerLogoutCsReq(h *HkRpgGoServer, p *PlayerGame, payloadMsg pb.Message) {
	h.DelPlayer(p.S.Uid)
}

func SendMsgCsReq(h *HkRpgGoServer, p *PlayerGame, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SendMsgCsReq)

	targetList := req.TargetList
	targetList = append(targetList, p.S.Uid)
	for _, targetUid := range targetList {
		target := h.GetPlayer(targetUid)
		if target == nil {
			continue
		}
		notify := &proto.RevcMsgScNotify{
			SourceUid:   p.S.Uid,
			MessageText: req.MessageText,
			ExtraId:     req.ExtraId,
			MessageType: req.MessageType,
			TargetUid:   targetUid,
			IGNEAJDPAPE: req.IGNEAJDPAPE,
			ChatType:    req.ChatType,
		}
		target.toSession(player.Msg{
			CmdId:     cmd.RevcMsgScNotify,
			PlayerMsg: notify,
		})
	}

	rsp := &proto.SendMsgScRsp{
		EndTime: uint64(time.Now().Unix()),
		Retcode: 0,
	}
	p.toSession(player.Msg{
		CmdId:     cmd.SendMsgScRsp,
		PlayerMsg: rsp,
	})
}

func ApplyFriendCsReq(h *HkRpgGoServer, p *PlayerGame, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.ApplyFriendCsReq)
	rsp := &proto.ApplyFriendScRsp{
		Uid:     req.Uid,
		Retcode: 0,
	}
	p.toSession(player.Msg{
		CmdId:     cmd.ApplyFriendScRsp,
		PlayerMsg: rsp,
	})
}
