package player

import (
	"encoding/base64"
	"strconv"
	"time"

	"github.com/gucooing/gunet"
	"github.com/gucooing/hkrpg-go/gameserver/db"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

var SNOWFLAKE *alg.SnowflakeWorker // 雪花唯一id生成器

type GamePlayer struct {
	Uid           uint32
	AccountId     uint32
	GameAppId     uint32
	GateAppId     uint32
	IsJumpMission bool
	DB            *gorm.DB
	IsPE          bool
	// 玩家数据
	Platform       spb.PlatformType        // 登录设备
	OnlineData     *OnlineData             // 玩家在线数据
	BasicBin       *spb.PlayerBasicCompBin // 玩家pb数据
	RouteManager   *RouteManager           // 路由
	SendChan       chan Msg                // 发送消息通道
	LastUpDataTime int64                   // 最近一次的活跃时间
}

type RecvMsg struct {
	CmdId     uint16
	PlayerMsg []byte
}

type Msg struct {
	AppId     uint32 // gs appid
	CmdId     uint16
	PlayerMsg pb.Message
}

func getCurTime() uint64 {
	return uint64(time.Now().UnixMilli())
}

// 拉取账户数据
func (g *GamePlayer) GetPlayerDateByDb() {
	var err error
	dbPlayer := database.QueryAccountUidByFieldPlayer(g.DB, g.Uid)
	if dbPlayer == nil || dbPlayer.BinData == nil {
		dbPlayer = new(database.PlayerData)
		logger.Info("新账号登录，进入初始化流程")
		g.BasicBin = g.NewBasicBin()
		// 初始化完毕保存账号数据
		dbPlayer.Uid = g.Uid
		dbPlayer.Level = g.GetLevel()
		dbPlayer.Exp = g.GetMaterialById(Exp)
		dbPlayer.Nickname = g.GetNickname()
		dbPlayer.BinData, err = pb.Marshal(g.BasicBin)
		dbPlayer.DataVersion = g.GetDataVersion()
		if err != nil {
			logger.Error("pb marshal error: %v", err)
		}

		if g.IsJumpMission {
			g.FinishAllMission()
			g.FinishAllTutorial()
		}

		err = database.AddDatePlayerFieldByFieldName(g.DB, dbPlayer)
		if err != nil {
			logger.Error("账号数据储存失败,err:%s", err.Error())
		}
	} else {
		g.BasicBin = new(spb.PlayerBasicCompBin)
		err = pb.Unmarshal(dbPlayer.BinData, g.BasicBin)
		if err != nil {
			logger.Error("unmarshal proto data err: %v", err)
			g.BasicBin = g.NewBasicBin()
		}
	}
	if g.GetIsProficientPlayer() { // 是否是老玩家

	} else {

	}
}

func (g *GamePlayer) UpPlayerDate(status spb.PlayerStatusType) bool {
	var err error
	if !g.IsPE {
		redisDb, ok := db.GetDb().GetPlayerStatus(strconv.Itoa(int(g.AccountId)))
		if !ok {
			return false
		}
		statu := new(spb.PlayerStatusRedisData)
		err = pb.Unmarshal(redisDb, statu)
		if err != nil {
			logger.Error("PlayerStatusRedisData Unmarshal error")
			db.GetDb().DistUnlockPlayerStatus(strconv.Itoa(int(g.AccountId)))
			return false
		}
		if statu.GameserverId != g.GameAppId && statu.DataVersion != g.GetDataVersion() {
			// 脏数据
			logger.Info("[UID:%v]数据过期，已丢弃", g.Uid)
			return false
		}
	}
	//  确认写入，更新数据版本
	g.AddDataVersion()
	dbDate := new(database.PlayerData)
	dbDate.Uid = g.Uid
	dbDate.Level = g.GetLevel()
	dbDate.Exp = g.GetMaterialById(Exp)
	dbDate.Nickname = g.GetNickname()
	dbDate.BinData, err = pb.Marshal(g.BasicBin)
	dbDate.DataVersion = g.GetDataVersion()
	if err != nil {
		logger.Error("pb marshal error: %v", err)
		return false
	}
	if err = database.UpdatePlayer(g.DB, dbDate); err != nil {
		logger.Error("Update Player error")
		return false
	}
	if !g.SetPlayerPlayerBasicBriefData(status) {
		logger.Error("[UID:%v]玩家简要信息保存失败", g.Uid)
	}
	// 保存地图数据
	for _, block := range g.GetAllBlockMap() {
		g.UpdateBlock(block)
	}

	return true
}

func (g *GamePlayer) SetPlayerPlayerBasicBriefData(status spb.PlayerStatusType) bool {
	playerBasicBrief := &spb.PlayerBasicBriefData{
		Nickname:          g.GetNickname(),
		Level:             g.GetLevel(),
		WorldLevel:        g.GetWorldLevel(),
		LastLoginTime:     time.Now().Unix(),
		HeadImageAvatarId: g.GetHeadIcon(),
		Exp:               g.GetMaterialById(Exp),
		PlatformType:      g.Platform,
		Uid:               g.Uid,
		Status:            status,
		Signature:         g.GetSignature(),
	}

	bin, err := pb.Marshal(playerBasicBrief)
	if err != nil {
		logger.Error("pb marshal error: %v", err)
		return false
	}
	if g.IsPE {
		player := &database.PlayerBasic{
			Uid:     g.Uid,
			BinData: bin,
		}
		return database.UpdatePlayerBasic(g.DB, player)
	} else {
		return db.GetDb().SetPlayerPlayerBasicBriefData(g.Uid, bin)
	}
}

func (g *GamePlayer) Send(cmdId uint16, playerMsg pb.Message) {
	// 打印需要的数据包
	go LogMsgSeed(cmdId, playerMsg)
	rspMsg := new(alg.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = playerMsg
	tcpMsg := alg.EncodeProtoToPayload(rspMsg)
	binMsg := alg.EncodePayloadToBin(tcpMsg, nil)

	// NewM
	if cmdId == cmd.GetTutorialGuideScRsp {
		newMsg := alg.EncodePayloadToBin(&alg.PackMsg{CmdId: NewM, ProtoData: gunet.GetGunetTcpConn()}, nil)
		gtgMsg := &spb.GameToGateMsgNotify{
			Uid: g.Uid,
			Msg: newMsg,
		}
		g.SendChan <- Msg{
			AppId:     g.GateAppId,
			CmdId:     cmd.GameToGateMsgNotify,
			PlayerMsg: gtgMsg,
		}
	}

	var msg Msg
	gtgMsg := &spb.GameToGateMsgNotify{
		Uid: g.Uid,
		Msg: binMsg,
	}
	if g.IsPE {
		msg = Msg{
			CmdId:     cmdId,
			PlayerMsg: playerMsg,
		}
	} else {
		msg = Msg{
			AppId:     g.GateAppId,
			CmdId:     cmd.GameToGateMsgNotify,
			PlayerMsg: gtgMsg,
		}
	}

	g.SendChan <- msg
}

func (g *GamePlayer) DecodePayloadToProto(cmdId uint16, msg []byte) (protoObj pb.Message) {
	protoObj = cmd.GetSharedCmdProtoMap().GetProtoObjCacheByCmdId(cmdId)
	if protoObj == nil {
		logger.Debug("get new proto object is nil")
		return nil
	}
	err := pb.Unmarshal(msg, protoObj)
	if err != nil {
		logger.Error("unmarshal proto data err: %v", err)
		return nil
	}
	return protoObj
}

var blacklist = []uint16{cmd.SceneEntityMoveScRsp, cmd.SceneEntityMoveCsReq, cmd.PlayerHeartBeatScRsp, cmd.PlayerHeartBeatCsReq} // 黑名单
func IsValid(cmdid uint16) bool {
	for _, value := range blacklist {
		if cmdid == value {
			return false
		}
	}
	return true
}

// 异步打印数据包
func LogMsgSeed(cmdId uint16, playerMsg pb.Message) {
	if IsValid(cmdId) {
		data := protojson.Format(playerMsg)
		logger.Debug("S --> C : NAME: %s KcpMsg: \n%s\n", cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(cmdId), data)
	}
}

func LogMsgRecv(cmdId uint16, payloadMsg []byte) {
	if IsValid(cmdId) {
		protoObj := cmd.GetSharedCmdProtoMap().GetProtoObjCacheByCmdId(cmdId)
		if protoObj == nil {
			logger.Debug("get new proto object is nil")
			return
		}
		err := pb.Unmarshal(payloadMsg, protoObj)
		if err != nil {
			logger.Error("unmarshal proto data NAME: %s  err: %v || b64:%s", cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(cmdId), err, base64.StdEncoding.EncodeToString(payloadMsg))
			return
		}
		data := protojson.Format(protoObj)
		logger.Debug("C --> S : NAME: %s KcpMsg: \n%s\n", cmd.GetSharedCmdProtoMap().GetCmdNameByCmdId(cmdId), data)
	}
}
