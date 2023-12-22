package internal

import (
	"math/rand"
	"time"

	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	"github.com/gucooing/hkrpg-go/robot/pkg/gdconf"
	pb "google.golang.org/protobuf/proto"
)

func (r *RoBot) GetCurSceneInfoScRsp(payloadMsg pb.Message) {
	rsp := payloadMsg.(*proto.GetCurSceneInfoScRsp)

	r.Game.EntryId = rsp.Scene.EntryId

	for _, entityGroup := range rsp.Scene.EntityGroupList {
		for _, entity := range entityGroup.EntityList {
			if entity.EntityId == rsp.Scene.LeaderEntityId {
				r.Game.Rot = &Vector{
					X: entity.Motion.Pos.X,
					Y: entity.Motion.Pos.Y,
					Z: entity.Motion.Pos.Z,
				}
				r.Game.Pos = &Vector{
					X: entity.Motion.Rot.X,
					Y: entity.Motion.Rot.Y,
					Z: entity.Motion.Rot.Z,
				}
				break
			}
		}
	}

	go r.EnterSceneCsReq()
}

func (r *RoBot) EnterSceneCsReq() {
	entryIdList := gdconf.GetEntryIdList()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		if r.KcpAddr == "" {
			return
		}
		entryId := rand.Intn(len(entryIdList)-1) + 1
		rep := &proto.EnterSceneCsReq{
			EntryId: entryIdList[entryId],
		}

		r.send(cmd.EnterSceneCsReq, rep)

		time.Sleep(3 * time.Second)
	}

}
