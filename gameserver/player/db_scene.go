package player

import (
	"strconv"
	"strings"
	"time"

	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/pkg/gdconf"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

type EntityAll interface {
	AddEntity(g *GamePlayer, groupID uint32)
}

type SceneMap struct {
	LoadedGroup   map[uint32]*GroupInfo // 已加载场景
	NoLoadedGroup map[uint32]*GroupInfo // 未加载场景
}

type GroupInfo struct {
	EntryId   uint32               // 地图
	PlaneID   uint32               // 地图2
	FloorID   uint32               // 地图3
	GroupID   uint32               // 区域
	EntityMap map[uint32]EntityAll // 场景实体(加载区域才写！！！
}

type Entity struct {
	EntityId uint32 // 实体id
	InstId   uint32 // 配置表中id
	EntryId  uint32 // 地图
	GroupId  uint32 // 区域
	Pos      *proto.Vector
	Rot      *proto.Vector
}

type AvatarEntity struct {
	Entity
	AvatarId uint32 // 角色id
}

type MonsterEntity struct {
	Entity
	EventID uint32 // 怪物id
}

type NpcEntity struct {
	Entity
	NpcId uint32 // ncp id
}

type PropEntity struct {
	Entity
	PropId uint32 // 物品id
}

func NewSceneMap() *SceneMap { // 清空实体列表用的
	db := &SceneMap{
		LoadedGroup:   make(map[uint32]*GroupInfo),
		NoLoadedGroup: make(map[uint32]*GroupInfo),
	}
	return db
}

func (g *GamePlayer) UpSceneMap() {
	db := g.GetOnlineData()
	db.SceneMap = NewSceneMap()
}

func (g *GamePlayer) GetSceneMap() *SceneMap {
	db := g.GetOnlineData()
	if db.SceneMap == nil {
		db.SceneMap = NewSceneMap()
	}
	return db.SceneMap
}

func (g *GamePlayer) GetLoadedGroup() map[uint32]*GroupInfo {
	db := g.GetSceneMap()
	if db.LoadedGroup == nil {
		db.LoadedGroup = make(map[uint32]*GroupInfo)
	}
	return db.LoadedGroup
}

func (g *GamePlayer) GetNoLoadedGroup() map[uint32]*GroupInfo {
	db := g.GetSceneMap()
	if db.NoLoadedGroup == nil {
		db.NoLoadedGroup = make(map[uint32]*GroupInfo)
	}
	return db.NoLoadedGroup
}

func (g *GamePlayer) AddLoadedGroup(entryId, planeID, floorID, groupID uint32) {
	db := g.GetLoadedGroup()
	db[groupID] = &GroupInfo{
		EntryId:   entryId,
		PlaneID:   planeID,
		FloorID:   floorID,
		GroupID:   groupID,
		EntityMap: make(map[uint32]EntityAll),
	}
}

func (g *GamePlayer) AddNoLoadedGroup(entryId, planeID, floorID, groupID uint32) {
	db := g.GetNoLoadedGroup()
	db[groupID] = &GroupInfo{
		EntryId:   entryId,
		PlaneID:   planeID,
		FloorID:   floorID,
		GroupID:   groupID,
		EntityMap: make(map[uint32]EntityAll),
	}
}

func (g *GamePlayer) GetGroupInfoByGroupID(groupID uint32) *GroupInfo {
	db := g.GetLoadedGroup()
	if db[groupID] == nil {
		db[groupID] = &GroupInfo{}
	}
	return db[groupID]
}

func (g *GamePlayer) GetEntity(groupID uint32) map[uint32]EntityAll {
	db := g.GetGroupInfoByGroupID(groupID)
	if db.EntityMap == nil {
		db.EntityMap = make(map[uint32]EntityAll)
	}
	return db.EntityMap
}

func (ae *AvatarEntity) AddEntity(g *GamePlayer, groupID uint32) {
	db := g.GetEntity(groupID)
	db[ae.EntityId] = ae
}

func (me *MonsterEntity) AddEntity(g *GamePlayer, groupID uint32) {
	db := g.GetEntity(groupID)
	db[me.EntityId] = me
}

func (ne *NpcEntity) AddEntity(g *GamePlayer, groupID uint32) {
	db := g.GetEntity(groupID)
	db[ne.EntityId] = ne
}

func (pe *PropEntity) AddEntity(g *GamePlayer, groupID uint32) {
	db := g.GetEntity(groupID)
	db[pe.EntityId] = pe
}

func (g *GamePlayer) AddEntity(groupID uint32, t EntityAll) {
	t.AddEntity(g, groupID)
}

func (g *GamePlayer) GetEntityById(id uint32) EntityAll { // 根据实体id拉取实体
	db := g.GetLoadedGroup()
	for _, info := range db {
		if info.EntityMap != nil {
			for eid := range info.EntityMap {
				if eid == id {
					return info.EntityMap[eid]
				}
			}
		}
	}
	return nil
}

func (g *GamePlayer) GetEntryId(t EntityAll) uint32 {
	if t == nil {
		return 0
	}
	switch t.(type) {
	case *PropEntity:
		return t.(*PropEntity).EntityId
	case *MonsterEntity:
		return t.(*MonsterEntity).EntityId
	case *NpcEntity:
		return t.(*NpcEntity).EntityId
	default:
		return 0
	}
}

func (g *GamePlayer) GetMonsterEntityById(id uint32) *MonsterEntity {
	db := g.GetEntityById(id)
	if db == nil {
		return nil
	}
	switch db.(type) {
	case *MonsterEntity:
		return db.(*MonsterEntity)
	}
	return nil
}

func (g *GamePlayer) GetPropEntityById(id uint32) *PropEntity {
	db := g.GetEntityById(id)
	if db == nil {
		return nil
	}
	switch db.(type) {
	case *PropEntity:
		return db.(*PropEntity)
	}
	return nil
}

func (g *GamePlayer) GetPropEntity(groupId, instId uint32) *PropEntity {
	db := g.GetEntity(groupId)
	for _, entity := range db {
		switch entity.(type) {
		case *PropEntity:
			if entity.(*PropEntity).GroupId == groupId && entity.(*PropEntity).InstId == instId {
				return entity.(*PropEntity)
			}
		}
	}
	return nil
}

func (g *GamePlayer) GetAvatarEntity(id uint32) *AvatarEntity {
	db := g.GetEntityById(id)
	if db == nil {
		return nil
	}
	switch db.(type) {
	case *AvatarEntity:
		return db.(*AvatarEntity)
	}
	return nil
}

func (g *GamePlayer) GetAllPropEntity() []*PropEntity {
	peList := make([]*PropEntity, 0)
	db := g.GetLoadedGroup()
	for _, groupInfo := range db {
		if groupInfo.EntityMap != nil {
			for _, info := range groupInfo.EntityMap {
				switch info.(type) {
				case *PropEntity:
					peList = append(peList, info.(*PropEntity))
				}
			}
		}
	}
	return peList
}

func NewScene() *spb.Scene {
	return &spb.Scene{
		EntryId: 2000101,
	}
}

func (g *GamePlayer) GetScene() *spb.Scene {
	db := g.BasicBin
	if db.Scene == nil {
		db.Scene = NewScene()
	}
	return db.Scene
}

func (g *GamePlayer) GetCurEntryId() uint32 {
	db := g.GetScene()
	return db.EntryId
}

func (g *GamePlayer) SetCurEntryId(id uint32) {
	db := g.GetScene()
	db.EntryId = id
}

func NewPos() *spb.VectorBin {
	return &spb.VectorBin{
		X: 99,
		Y: 62,
		Z: -4800,
	}
}

func NewRot() *spb.VectorBin {
	return &spb.VectorBin{
		X: 0,
		Y: 0,
		Z: 0,
	}
}

func (g *GamePlayer) GetPos() *spb.VectorBin {
	db := g.GetBasicBin()
	if db.Pos == nil {
		db.Pos = NewPos()
	}
	return db.Pos
}

func (g *GamePlayer) GetRot() *spb.VectorBin {
	db := g.GetBasicBin()
	if db.Rot == nil {
		db.Rot = NewRot()
	}
	return db.Rot
}

func (g *GamePlayer) SetPos(x, y, z int32) {
	db := g.GetPos()
	db.X = x
	db.Y = y
	db.Z = z
}

func (g *GamePlayer) SetRot(x, y, z int32) {
	db := g.GetRot()
	db.X = x
	db.Y = y
	db.Z = z
}

func (g *GamePlayer) IfLoadMap(levelGroup *gdconf.GoppLevelGroup) bool {
	if levelGroup.GroupName == "TrainVisitorDemo" {
		return false
	}
	switch levelGroup.Category {
	case "": // 基础
		return g.IfMissionLoadMap(levelGroup, true)
	case "System": // 副本/关卡/等入口
		return g.IfMissionLoadMap(levelGroup, true)
	case "Atmosphere": // 特殊交互物品
		return true
	case "Custom": // 特殊环境场景:模拟宇宙等
		return false
	case "Mission": // 任务
		return g.IfMissionLoadMap(levelGroup, false)
	default:
		logger.Warn("未知的地图类型 Category:%s", levelGroup.Category)
		return true
	}
}

func (g *GamePlayer) IfMissionLoadMap(levelGroup *gdconf.GoppLevelGroup, mainIsLoaded bool) bool {
	finishSubMainMissionList := g.GetFinishSubMainMissionList() // 已完成子任务
	subMainMissionList := g.GetSubMainMissionList()             // 接受的子任务
	mainMissionList := g.GetMainMissionList()                   // 接取的主任务
	finishMainMissionList := g.GetFinishMainMissionList()       // 已完成的主任务
	isLoaded := mainIsLoaded
	if levelGroup.OwnerMainMissionID != 0 {
		if mainMissionList[levelGroup.OwnerMainMissionID] == nil {
			return false
		}
	}
	if levelGroup.LoadCondition == nil &&
		levelGroup.UnloadCondition == nil {
		if levelGroup.Category == "Mission" && levelGroup.OwnerMainMissionID != 0 {
			return true
		}
		return mainIsLoaded
	}
	// 检查强制卸载条件
	// 检查加载条件
	// if levelGroup.GroupId == 66 {
	// 	logger.Info("")
	// 	return true
	// }
	if levelGroup.LoadCondition != nil {
		for _, conditions := range levelGroup.LoadCondition.Conditions {
			if conditions.Phase == "Finish" { // 完成了这个任务
				if finishSubMainMissionList[conditions.ID] != nil || finishMainMissionList[conditions.ID] != nil {
					isLoaded = true
					if levelGroup.LoadCondition.Operation == "Or" {
						break
					}
				} else {
					isLoaded = false
				}
				continue
			}
			if conditions.Type == "SubMission" && conditions.Phase == "" { // 接取了这个子任务
				if subMainMissionList[conditions.ID] != nil {
					isLoaded = true
					if levelGroup.LoadCondition.Operation == "Or" {
						break
					}
				} else {
					isLoaded = false
				}
				continue
			}
			if conditions.Type == "" && conditions.Phase == "" { // 接取主线任务
				if mainMissionList[conditions.ID] != nil {
					isLoaded = true
					if levelGroup.LoadCondition.Operation == "Or" {
						break
					}
				} else {
					isLoaded = false
				}
				continue
			}
		}
	}

	if !isLoaded {
		return isLoaded
	}

	// 检查卸载条件
	if levelGroup.UnloadCondition != nil {
		all := false
		for _, conditions := range levelGroup.UnloadCondition.Conditions {
			if conditions.Phase == "Finish" { // 完成了这个任务
				if finishSubMainMissionList[conditions.ID] != nil || finishMainMissionList[conditions.ID] != nil {
					if levelGroup.UnloadCondition.Operation == "Or" {
						isLoaded = false
						break
					} else {
						all = true
					}
				}
				continue
			}
			if conditions.Phase == "" { // 接取了这个任务
				if subMainMissionList[conditions.ID] != nil || mainMissionList[conditions.ID] != nil {
					if levelGroup.UnloadCondition.Operation == "Or" {
						isLoaded = false
						break
					} else {
						all = true
					}
				}
				continue
			}
		}
		if all {
			isLoaded = false
		}
	}

	return isLoaded
}

// 检查场景上是否有实体需要卸载/加载
func (g *GamePlayer) AutoEntryGroup() {
	loadedGroup := g.GetLoadedGroup()        // 已加载区域
	noLoadedGroup := g.GetNoLoadedGroup()    // 未加载区域
	uninstallGroup := make([]*GroupInfo, 0)  // 卸载场景列表
	loadedGroupList := make([]*GroupInfo, 0) // 加载列表
	// 检查已加载区域是否需要卸载
	for id, info := range loadedGroup {
		group := gdconf.GetServerGroupById(info.PlaneID, info.FloorID, info.GroupID)
		if group == nil {
			continue
		}
		if !g.IfLoadMap(group) {
			// 此处卸载
			uninstallGroup = append(uninstallGroup, info)
			noLoadedGroup[id] = info
			delete(loadedGroup, id)
		}
	}
	// 检查未加载区域是否需要加载
	for id, info := range noLoadedGroup {
		group := gdconf.GetServerGroupById(info.PlaneID, info.FloorID, info.GroupID)
		if group == nil {
			continue
		}
		if g.IfLoadMap(group) {
			// 此处加载
			loadedGroupList = append(loadedGroupList, info)
			loadedGroup[id] = info
			delete(noLoadedGroup, id)
		}
	}

	// 卸载/加载
	g.UpSceneGroupRefreshScNotify(uninstallGroup, loadedGroupList)
}

func NewBlockMap() map[uint32]*spb.BlockBin {
	return map[uint32]*spb.BlockBin{}
}

func (g *GamePlayer) GetBlockMap() map[uint32]*spb.BlockBin {
	db := g.GetOnlineData()
	if db.BlockMap == nil {
		db.BlockMap = NewBlockMap()
	}
	return db.BlockMap
}

func (g *GamePlayer) GetAllBlockMap() map[uint32]*spb.BlockBin {
	db := g.GetOnlineData()
	if db.BlockMap == nil {
		db.BlockMap = NewBlockMap()
	}
	blockMap := make(map[uint32]*spb.BlockBin)
	db.blockMapLock.Lock()
	defer db.blockMapLock.Unlock()
	for k, v := range db.BlockMap {
		blockMap[k] = v
	}
	return blockMap
}

// 从db拉取地图数据
func (g *GamePlayer) GetBlock(entryId uint32) *spb.BlockBin {
	newEntryId := entryId
	if entryId >= 10000000 {
		newEntryId = alg.S2U32(strconv.Itoa(int(entryId))[:7])
	}
	if mapEntrance := gdconf.GetMapEntranceById(newEntryId); mapEntrance == nil {
		newEntryId = entryId
	}
	on := g.GetOnlineData()
	db := g.GetBlockMap()
	on.blockMapLock.Lock()
	defer on.blockMapLock.Unlock()
	if db[newEntryId] == nil {
		bin := database.GetBlockData(g.DB, g.Uid, newEntryId)
		block := new(spb.BlockBin)
		if err := pb.Unmarshal(bin.BinData, block); err != nil {
			logger.Debug("entryId:%v,block error", newEntryId)
		}
		db[newEntryId] = block
	}

	db[newEntryId].EntryId = newEntryId
	return db[newEntryId]
}

// 更新地图数据到数据库
func (g *GamePlayer) UpdateBlock(block *spb.BlockBin) {
	bin, err := pb.Marshal(block)
	if err != nil {
		return
	}
	newEntryId := block.EntryId
	if block.EntryId >= 10000000 {
		newEntryId = alg.S2U32(strconv.Itoa(int(block.EntryId))[:7])
	}
	if mapEntrance := gdconf.GetMapEntranceById(newEntryId); mapEntrance == nil {
		newEntryId = block.EntryId
	}
	blockData := &database.BlockData{
		Uid:         g.Uid,
		EntryId:     newEntryId,
		DataVersion: 0, // TODO
		BinData:     bin,
	}
	if err = database.UpdateBlockData(g.DB, blockData); err != nil {
		logger.Debug("updata block data error:%s", err.Error())
	}
}

func (g *GamePlayer) GetPropState(db *spb.BlockBin, groupId, propId uint32, state string) uint32 {
	if db == nil {
		return gdconf.GetStateValue(state)
	}
	if db.BlockList == nil {
		db.BlockList = make(map[uint32]*spb.BlockList)
	}
	if db.BlockList[groupId] == nil {
		db.BlockList[groupId] = &spb.BlockList{
			PropInfo: make(map[uint32]*spb.PropInfo),
		}
	}
	if db.BlockList[groupId].PropInfo == nil {
		db.BlockList[groupId].PropInfo = make(map[uint32]*spb.PropInfo)
	}
	if db.BlockList[groupId].PropInfo[propId] == nil {
		db.BlockList[groupId].PropInfo[propId] = &spb.PropInfo{
			InstId: propId,
			// PropId:    prop.PropID,
			PropState: gdconf.GetStateValue(state),
		}
	}
	return db.BlockList[groupId].PropInfo[propId].PropState
}

func (g *GamePlayer) UpPropState(db *spb.BlockBin, groupId, propId, state uint32) {
	if db.BlockList == nil {
		db.BlockList = make(map[uint32]*spb.BlockList)
	}
	if db.BlockList[groupId] == nil {
		db.BlockList[groupId] = &spb.BlockList{
			PropInfo: make(map[uint32]*spb.PropInfo),
		}
	}
	if db.BlockList[groupId].PropInfo == nil {
		db.BlockList[groupId].PropInfo = make(map[uint32]*spb.PropInfo)
	}
	if db.BlockList[groupId].PropInfo[propId] == nil {
		db.BlockList[groupId].PropInfo[propId] = &spb.PropInfo{
			InstId:    propId,
			PropState: state,
		}
	} else {
		db.BlockList[groupId].PropInfo[propId].PropState = state
	}
}

func (g *GamePlayer) GetGroupState(db *spb.BlockBin, groupId uint32) uint32 {
	if db.BlockList == nil {
		db.BlockList = make(map[uint32]*spb.BlockList)
	}
	if db.BlockList[groupId] == nil {
		db.BlockList[groupId] = &spb.BlockList{
			PropInfo: make(map[uint32]*spb.PropInfo),
		}
	}
	return db.BlockList[groupId].GroupState
}

func (g *GamePlayer) SetGroupState(db *spb.BlockBin, groupId, groupState uint32) {
	if db.BlockList == nil {
		db.BlockList = make(map[uint32]*spb.BlockList)
	}
	if db.BlockList[groupId] == nil {
		db.BlockList[groupId] = &spb.BlockList{
			PropInfo: make(map[uint32]*spb.PropInfo),
		}
	}
	db.BlockList[groupId].GroupState = groupState
}

func (g *GamePlayer) ObjectCaptureUpPropState(db *spb.BlockBin, groupId, propId, state uint32) {
	if db.BlockList == nil {
		db.BlockList = make(map[uint32]*spb.BlockList)
	}
	if db.BlockList[groupId] == nil {
		db.BlockList[groupId] = &spb.BlockList{
			PropInfo: make(map[uint32]*spb.PropInfo),
		}
	}
	if db.BlockList[groupId].PropInfo == nil {
		db.BlockList[groupId].PropInfo = make(map[uint32]*spb.PropInfo)
	}
	if db.BlockList[groupId].PropInfo[propId] == nil {
		db.BlockList[groupId].PropInfo[propId] = &spb.PropInfo{
			InstId:    propId,
			PropState: state,
		}
	}
}

func (g *GamePlayer) StageObjectCapture(prop *gdconf.PropList, groupId uint32, db *spb.BlockBin) {
	if db == nil {
		return
	}
	if strings.Contains(prop.Name, "Elevator0") {
		g.ObjectCaptureUpPropState(db, groupId, prop.ID, 15)
	}
	if prop.ValueSource != nil {
		for _, v := range prop.ValueSource.Values {
			if v.Key == "IsAutoDoor" {
				g.ObjectCaptureUpPropState(db, groupId, prop.ID, 1)
				break
			}
		}
	}
	if prop.StageObjectCapture != nil { // 特殊处理
		switch prop.StageObjectCapture.BlockAlias {
		case "RogueLobby_01": // 模拟宇宙入口直接开放
			g.ObjectCaptureUpPropState(db, groupId, prop.ID, 1)
			break
		}
	}
	if conf := gdconf.GetSpecialProp(db.EntryId); conf != nil {
		if spGroup := conf.GroupList[groupId]; spGroup != nil {
			if state := spGroup.PropState[prop.ID]; state != "" {
				g.ObjectCaptureUpPropState(db, groupId, prop.ID, gdconf.GetStateValue(state))
			}
		}
	}
}

func floorTentry(floorID uint32) uint32 {
	if floorID < 10000000 {
		return 1000001
	}
	st := strconv.Itoa(int(floorID))
	entryId := alg.S2U32(st[:6] + st[7:8])
	if mapEntrance := gdconf.GetMapEntranceById(entryId); mapEntrance == nil {
		return floorID
	}
	return entryId
}

// 任务设置物品状态
func (g *GamePlayer) SetFloorSavedValue(conf *gdconf.SubMission, finishAction *gdconf.FinishAction) {
	if len(finishAction.FinishActionParaString) < 4 {
		return
	}
	planeId := alg.S2U32(finishAction.FinishActionParaString[0])
	floorId := alg.S2U32(finishAction.FinishActionParaString[1])
	name := finishAction.FinishActionParaString[2]
	state := alg.S2I32(finishAction.FinishActionParaString[3])
	notify := &proto.UpdateFloorSavedValueNotify{
		SavedValue: map[string]int32{name: state},
	}
	g.Send(cmd.UpdateFloorSavedValueNotify, notify)

	db := g.GetBlock(floorTentry(conf.WayPointFloorID))
	groupID, instId := gdconf.GetSavedValue(planeId, floorId, name)
	if groupID == 0 || instId == 0 {
		return // TODO subMission 103030204
	}
	g.UpPropState(db, groupID, instId, uint32(state))
	if enep := g.GetPropEntity(groupID, instId); enep != nil {
		g.PropSceneGroupRefreshScNotify([]uint32{enep.EntityId}, db)
	}
}

/****************************************************功能***************************************************/

func (g *GamePlayer) GetPosPb() *proto.Vector {
	db := g.GetPos()
	return &proto.Vector{
		Y: db.Y,
		X: db.X,
		Z: db.Z,
	}
}

func (g *GamePlayer) GetRotPb() *proto.Vector {
	db := g.GetRot()
	return &proto.Vector{
		Y: db.Y,
		X: db.X,
		Z: db.Z,
	}
}

func (g *GamePlayer) GetSceneAvatarByLineUP(entityGroupList *proto.SceneEntityGroupInfo, lineUp *spb.Line, leaderEntityId uint32, pos, rot *proto.Vector) {
	for sole, lineAvatar := range lineUp.AvatarIdList {
		if lineAvatar.AvatarId == 0 {
			continue
		}
		actor := &proto.SceneActorInfo{
			AvatarType:   proto.AvatarType_AVATAR_FORMAL_TYPE,
			BaseAvatarId: lineAvatar.AvatarId,
			MapLayer:     0,
			Uid:          0,
		}
		if lineAvatar.LineAvatarType == spb.LineAvatarType_LineAvatarType_TRIAL {
			conf := gdconf.GetSpecialAvatarById(lineAvatar.AvatarId)
			if conf == nil {
				continue
			}
			actor = &proto.SceneActorInfo{
				AvatarType:   proto.AvatarType_AVATAR_TRIAL_TYPE,
				BaseAvatarId: conf.AvatarID,
				MapLayer:     0,
				Uid:          0,
			}
		}
		entityList := &proto.SceneEntityInfo{
			Actor: actor,
			Motion: &proto.MotionInfo{
				Pos: pos,
				Rot: rot,
			},
		}
		if sole == lineUp.LeaderSlot {
			entityList.EntityId = leaderEntityId
		} else {
			entityId := g.GetNextGameObjectGuid()
			entityList.EntityId = entityId
		}
		g.AddEntity(0, &AvatarEntity{
			Entity: Entity{
				EntityId: entityList.EntityId,
				GroupId:  0,
				Pos:      pos,
				Rot:      rot,
			},
			AvatarId: lineAvatar.AvatarId,
		})
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
}

func (g *GamePlayer) GetPropByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.GoppLevelGroup, db *spb.BlockBin, entryId uint32) *proto.SceneEntityGroupInfo {
	for _, propList := range sceneGroup.PropList {
		entityId := g.GetNextGameObjectGuid()
		g.StageObjectCapture(propList, sceneGroup.GroupId, db)
		pos := &proto.Vector{
			X: int32(propList.PosX * 1000),
			Y: int32(propList.PosY * 1000),
			Z: int32(propList.PosZ * 1000),
		}
		rot := &proto.Vector{
			X: int32(propList.RotX * 1000),
			Y: int32(propList.RotY * 1000),
			Z: int32(propList.RotZ * 1000),
		}
		entityList := &proto.SceneEntityInfo{
			GroupId:  sceneGroup.GroupId, // 文件名后那个G
			InstId:   propList.ID,        // ID
			EntityId: entityId,
			Motion: &proto.MotionInfo{
				Pos: pos,
				Rot: rot,
			},
			Prop: &proto.ScenePropInfo{
				PropId:    propList.PropID, // PropID
				PropState: g.GetPropState(db, sceneGroup.GroupId, propList.ID, propList.State),
			},
		}
		// 添加物品实体
		g.AddEntity(sceneGroup.GroupId, &PropEntity{
			Entity: Entity{
				EntityId: entityId,
				InstId:   propList.ID,
				EntryId:  entryId,
				GroupId:  sceneGroup.GroupId,
				Pos:      pos,
				Rot:      rot,
			},
			PropId: propList.PropID,
		})
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
	return entityGroupList
}

func (g *GamePlayer) GetNPCMonsterByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.GoppLevelGroup) *proto.SceneEntityGroupInfo {
	for _, monsterList := range sceneGroup.MonsterList {
		entityId := g.GetNextGameObjectGuid()
		pos := &proto.Vector{
			X: int32(monsterList.PosX * 1000),
			Y: int32(monsterList.PosY * 1000),
			Z: int32(monsterList.PosZ * 1000),
		}
		rot := &proto.Vector{
			X: int32(monsterList.RotX * 1000),
			Y: int32(monsterList.RotY * 1000),
			Z: int32(monsterList.RotZ * 1000),
		}
		entityList := &proto.SceneEntityInfo{
			GroupId:  sceneGroup.GroupId,
			InstId:   monsterList.ID,
			EntityId: entityId,
			Motion: &proto.MotionInfo{
				Pos: pos,
				Rot: rot,
			},
			NpcMonster: &proto.SceneNpcMonsterInfo{
				WorldLevel: g.GetWorldLevel(),
				MonsterId:  monsterList.NPCMonsterID,
				EventId:    monsterList.EventID,
			},
		}
		// 添加怪物实体
		g.AddEntity(sceneGroup.GroupId, &MonsterEntity{
			Entity: Entity{
				InstId:   monsterList.ID,
				EntityId: entityId,
				GroupId:  sceneGroup.GroupId,
				Pos:      pos,
				Rot:      rot,
			},
			EventID: monsterList.EventID,
		})
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
	return entityGroupList
}

func (g *GamePlayer) GetNPCByID(entityGroupList *proto.SceneEntityGroupInfo, sceneGroup *gdconf.GoppLevelGroup) *proto.SceneEntityGroupInfo {
	for _, npcList := range sceneGroup.NPCList {
		entityId := g.GetNextGameObjectGuid()
		pos := &proto.Vector{
			X: int32(npcList.PosX * 1000),
			Y: int32(npcList.PosY * 1000),
			Z: int32(npcList.PosZ * 1000),
		}
		rot := &proto.Vector{
			X: int32(npcList.RotX * 1000),
			Y: int32(npcList.RotY * 1000),
			Z: int32(npcList.RotZ * 1000),
		}
		entityList := &proto.SceneEntityInfo{
			GroupId:  sceneGroup.GroupId,
			InstId:   npcList.ID,
			EntityId: entityId,
			Motion: &proto.MotionInfo{
				Pos: pos,
				Rot: rot,
			},
			Npc: &proto.SceneNpcInfo{
				ExtraInfo: nil,
				NpcId:     npcList.NPCID,
			},
		}
		// 添加npc
		g.AddEntity(sceneGroup.GroupId, &NpcEntity{
			Entity: Entity{
				EntityId: entityId,
				GroupId:  sceneGroup.GroupId,
				Pos:      pos,
				Rot:      rot,
				InstId:   npcList.ID,
			},
			NpcId: npcList.NPCID,
		})
		entityGroupList.EntityList = append(entityGroupList.EntityList, entityList)
	}
	return entityGroupList
}

func (g *GamePlayer) GetSceneInfo(entryId uint32, pos, rot *proto.Vector, lineUp *spb.Line) *proto.SceneInfo {
	leaderEntityId := g.GetNextGameObjectGuid()
	mapEntrance := gdconf.GetMapEntranceById(entryId)
	if mapEntrance == nil {
		return nil
	}
	foorMap := gdconf.GetServerGroup(mapEntrance.PlaneID, mapEntrance.FloorID)
	if foorMap == nil {
		return nil
	}
	scene := &proto.SceneInfo{
		WorldId:            gdconf.GetMazePlaneById(mapEntrance.PlaneID).WorldID,
		LeaderEntityId:     leaderEntityId,
		FloorId:            mapEntrance.FloorID,
		GameModeType:       gdconf.GetPlaneType(gdconf.GetMazePlaneById(mapEntrance.PlaneID).PlaneType),
		PlaneId:            mapEntrance.PlaneID,
		EntryId:            entryId,
		EntityGroupList:    make([]*proto.SceneEntityGroupInfo, 0),
		GroupIdList:        make([]uint32, 0),
		LightenSectionList: make([]uint32, 0),
		GroupStateList:     make([]*proto.SceneGroupState, 0),
		SceneMissionInfo:   g.GetMissionStatusBySceneInfo(gdconf.GetGroupById(mapEntrance.PlaneID, mapEntrance.FloorID)),
	}
	for i := uint32(0); i < 100; i++ {
		scene.LightenSectionList = append(scene.LightenSectionList, i)
	}
	// 获取场景实体
	entityGroup := &proto.SceneEntityGroupInfo{
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	// 清理老实体列表
	g.UpSceneMap()
	// 添加队伍角色进实体列表，并设置坐标
	g.GetSceneAvatarByLineUP(entityGroup, lineUp, leaderEntityId, pos, rot)
	blockBin := g.GetBlock(entryId)
	scene.EntityGroupList = append(scene.EntityGroupList, entityGroup)
	for _, levelGroup := range foorMap {
		if !g.IfLoadMap(levelGroup) {
			g.AddNoLoadedGroup(entryId, mapEntrance.PlaneID, mapEntrance.FloorID, levelGroup.GroupId)
			continue
		} else {
			g.AddLoadedGroup(entryId, mapEntrance.PlaneID, mapEntrance.FloorID, levelGroup.GroupId)
		}
		scene.GroupIdList = append(scene.GroupIdList, levelGroup.GroupId)
		entityGroupLists := &proto.SceneEntityGroupInfo{
			GroupId:    levelGroup.GroupId,
			EntityList: make([]*proto.SceneEntityInfo, 0),
			State:      g.GetGroupState(blockBin, levelGroup.GroupId),
		}
		// 添加物品实体
		g.GetPropByID(entityGroupLists, levelGroup, blockBin, entryId)
		// 添加怪物实体
		g.GetNPCMonsterByID(entityGroupLists, levelGroup)
		// 添加NPC实体
		g.GetNPCByID(entityGroupLists, levelGroup)
		scene.EntityGroupList = append(scene.EntityGroupList, entityGroupLists)
	}
	return scene
}

func (g *GamePlayer) GetMissionStatusBySceneInfo(foorMap map[uint32]*gdconf.LevelGroup) *proto.MissionStatusBySceneInfo {
	info := &proto.MissionStatusBySceneInfo{
		AcceptMainMissionIdList: make([]uint32, 0),
		MainMissionIdList:       make([]uint32, 0),
		SceneSubMissionList:     make([]*proto.Mission, 0),
	}
	if foorMap == nil {
		return info
	}
	mainMissionList := g.GetMainMissionList()
	finishMainMissionList := g.GetFinishMainMissionList()
	subMissionList := g.GetSubMainMissionList()
	finishSubMissionList := g.GetFinishSubMainMissionList()
	for _, groupInfo := range foorMap {
		if groupInfo.OwnerMainMissionID != 0 {
			var isAdd = true
			for _, id := range info.AcceptMainMissionIdList {
				if id == groupInfo.OwnerMainMissionID {
					isAdd = false
					break
				}
			}
			for _, id := range info.MainMissionIdList {
				if id == groupInfo.OwnerMainMissionID {
					isAdd = false
					break
				}
			}
			if isAdd {
				if mainMissionList[groupInfo.OwnerMainMissionID] != nil {
					info.AcceptMainMissionIdList = append(info.AcceptMainMissionIdList, groupInfo.OwnerMainMissionID)
				}
				if finishMainMissionList[groupInfo.OwnerMainMissionID] != nil {
					info.MainMissionIdList = append(info.MainMissionIdList, groupInfo.OwnerMainMissionID)
				}
			}
		}
		if groupInfo.AtmosphereCondition != nil {
			for _, conditions := range groupInfo.AtmosphereCondition.Conditions {
				var isAdd = true
				for _, v := range info.SceneSubMissionList {
					if v.Id == conditions.SubMissionID {
						isAdd = false
						break
					}
				}
				if isAdd {
					db := finishSubMissionList[conditions.SubMissionID]
					if db == nil {
						db = subMissionList[conditions.SubMissionID]
					}
					if db != nil {
						info.SceneSubMissionList = append(info.SceneSubMissionList, &proto.Mission{
							Status:   proto.MissionStatus(db.Status),
							Progress: db.Progress,
							Id:       conditions.SubMissionID,
						})
					}
					// info.SceneSubMissionList = append(info.SceneSubMissionList, &proto.Mission{
					// 	Id: conditions.SubMissionID,
					// })
				}
			}
		}
		if groupInfo.LoadCondition != nil {
			for _, conditions := range groupInfo.LoadCondition.Conditions {
				var isAdd = true
				if conditions.Type != "SubMission" {
					continue
				}
				for _, v := range info.SceneSubMissionList {
					if v.Id == conditions.ID {
						isAdd = false
						break
					}
				}
				if isAdd {
					db := finishSubMissionList[conditions.ID]
					if db == nil {
						db = subMissionList[conditions.ID]
					}
					if db != nil {
						info.SceneSubMissionList = append(info.SceneSubMissionList, &proto.Mission{
							Status:   proto.MissionStatus(db.Status),
							Progress: db.Progress,
							Id:       conditions.ID,
						})
					}
					// info.SceneSubMissionList = append(info.SceneSubMissionList, &proto.Mission{
					// 	Id: conditions.ID,
					// })
				}
			}
		}
		if groupInfo.UnloadCondition != nil {
			for _, conditions := range groupInfo.UnloadCondition.Conditions {
				var isAdd = true
				if conditions.Type != "SubMission" {
					continue
				}
				for _, v := range info.SceneSubMissionList {
					if v.Id == conditions.ID {
						isAdd = false
						break
					}
				}
				if isAdd {
					db := finishSubMissionList[conditions.ID]
					if db == nil {
						db = subMissionList[conditions.ID]
					}
					if db != nil {
						info.SceneSubMissionList = append(info.SceneSubMissionList, &proto.Mission{
							Status:   proto.MissionStatus(db.Status),
							Progress: db.Progress,
							Id:       conditions.ID,
						})
					}
					// info.SceneSubMissionList = append(info.SceneSubMissionList, &proto.Mission{
					// 	Id: conditions.ID,
					// })
				}
			}
		}
	}
	return info
}

// 删除怪物
func (g *GamePlayer) GetDelSceneGroupRefreshInfo(mem []uint32) []*proto.GroupRefreshInfo {
	sceneGroupRefreshInfo := make([]*proto.GroupRefreshInfo, 0)
	for _, id := range mem {
		entity := g.GetMonsterEntityById(id)
		if entity == nil {
			continue
		}
		sgri := &proto.GroupRefreshInfo{
			State:   0,
			GroupId: entity.GroupId,
			RefreshEntity: []*proto.SceneEntityRefreshInfo{
				{
					DeleteEntity: entity.EntityId,
				},
			},
			RefreshType: proto.SceneGroupRefreshType_SCENE_GROUP_REFRESH_TYPE_LOADED,
		}
		sceneGroupRefreshInfo = append(sceneGroupRefreshInfo, sgri)
	}
	return sceneGroupRefreshInfo
}

// 添加怪物
func (g *GamePlayer) AddMonsterSceneEntityRefreshInfo(mazeGroupID uint32, monsterList map[uint32]*gdconf.MonsterList) []*proto.SceneEntityRefreshInfo {
	sceneEntityRefreshInfo := make([]*proto.SceneEntityRefreshInfo, 0)
	for _, monster := range monsterList {
		entityId := g.GetNextGameObjectGuid()
		monsterPos := &proto.Vector{
			X: int32(monster.PosX * 1000),
			Y: int32(monster.PosY * 1000),
			Z: int32(monster.PosZ * 1000),
		}
		monsterRot := &proto.Vector{
			X: int32(monster.RotX * 1000),
			Y: int32(monster.RotY * 1000),
			Z: int32(monster.RotZ * 1000),
		}
		seri := &proto.SceneEntityRefreshInfo{
			AddEntity: &proto.SceneEntityInfo{
				GroupId:  mazeGroupID,
				InstId:   monster.ID,
				EntityId: entityId,
				Motion: &proto.MotionInfo{
					Pos: monsterPos,
					Rot: monsterRot,
				},
				NpcMonster: &proto.SceneNpcMonsterInfo{
					WorldLevel: g.GetWorldLevel(),
					MonsterId:  monster.NPCMonsterID,
					EventId:    monster.EventID,
				},
			},
		}
		// 添加怪物实体
		g.AddEntity(mazeGroupID, &MonsterEntity{
			Entity: Entity{
				EntityId: entityId,
				GroupId:  mazeGroupID,
				Pos:      monsterPos,
				Rot:      monsterRot,
				InstId:   monster.ID,
			},
			EventID: monster.EventID,
		})
		sceneEntityRefreshInfo = append(sceneEntityRefreshInfo, seri)
	}
	return sceneEntityRefreshInfo
}

// 添加Npc
func (g *GamePlayer) AddNpcSceneEntityRefreshInfo(mazeGroupID uint32, npcList map[uint32]*gdconf.NPCList) []*proto.SceneEntityRefreshInfo {
	sceneEntityRefreshInfo := make([]*proto.SceneEntityRefreshInfo, 0)
	for _, npc := range npcList {
		entityId := g.GetNextGameObjectGuid()
		pos := &proto.Vector{
			X: int32(npc.PosX * 1000),
			Y: int32(npc.PosY * 1000),
			Z: int32(npc.PosZ * 1000),
		}
		rot := &proto.Vector{
			X: int32(npc.RotX * 1000),
			Y: int32(npc.RotY * 1000),
			Z: int32(npc.RotZ * 1000),
		}
		seri := &proto.SceneEntityRefreshInfo{
			AddEntity: &proto.SceneEntityInfo{
				GroupId:  mazeGroupID,
				InstId:   npc.ID,
				EntityId: entityId,
				Motion: &proto.MotionInfo{
					Pos: pos,
					Rot: rot,
				},
				Npc: &proto.SceneNpcInfo{
					ExtraInfo: nil,
					NpcId:     npc.NPCID,
				},
			},
		}
		// 添加Npc实体
		g.AddEntity(mazeGroupID, &NpcEntity{
			Entity: Entity{
				EntityId: entityId,
				GroupId:  mazeGroupID,
				Pos:      pos,
				Rot:      rot,
				InstId:   npc.ID,
			},
			NpcId: npc.NPCID,
		})
		sceneEntityRefreshInfo = append(sceneEntityRefreshInfo, seri)
	}
	return sceneEntityRefreshInfo
}

// 添加物品实体
func (g *GamePlayer) AddPropSceneEntityRefreshInfo(mazeGroupID uint32, propList map[uint32]*gdconf.PropList, db *spb.BlockBin) []*proto.SceneEntityRefreshInfo {
	sceneEntityRefreshInfo := make([]*proto.SceneEntityRefreshInfo, 0)
	for _, prop := range propList {
		g.StageObjectCapture(prop, mazeGroupID, db)
		entityId := g.GetNextGameObjectGuid()
		pos := &proto.Vector{
			X: int32(prop.PosX * 1000),
			Y: int32(prop.PosY * 1000),
			Z: int32(prop.PosZ * 1000),
		}
		rot := &proto.Vector{
			X: int32(prop.RotX * 1000),
			Y: int32(prop.RotY * 1000),
			Z: int32(prop.RotZ * 1000),
		}
		seri := &proto.SceneEntityRefreshInfo{
			AddEntity: &proto.SceneEntityInfo{
				GroupId:  mazeGroupID,
				InstId:   prop.ID,
				EntityId: entityId,
				Motion: &proto.MotionInfo{
					Pos: pos,
					Rot: rot,
				},
				Prop: &proto.ScenePropInfo{
					PropId:    prop.PropID, // PropID
					PropState: g.GetPropState(db, mazeGroupID, prop.ID, prop.State),
				},
			},
		}
		// 添加物品实体
		g.AddEntity(mazeGroupID, &PropEntity{
			Entity: Entity{
				EntityId: entityId,
				InstId:   prop.ID,
				EntryId:  db.EntryId,
				GroupId:  mazeGroupID,
				Pos:      pos,
				Rot:      rot,
			},
			PropId: prop.PropID,
		})
		sceneEntityRefreshInfo = append(sceneEntityRefreshInfo, seri)
	}
	return sceneEntityRefreshInfo
}

// 添加角色
func (g *GamePlayer) GetAddAvatarSceneEntityRefreshInfo(lineUp *spb.Line, pos, rot *proto.Vector) []*proto.SceneEntityRefreshInfo {
	sceneEntityRefreshInfo := make([]*proto.SceneEntityRefreshInfo, 0)
	for _, lineAvatar := range lineUp.AvatarIdList {
		if lineAvatar.AvatarId == 0 {
			continue
		}
		actor := &proto.SceneActorInfo{
			AvatarType:   proto.AvatarType_AVATAR_FORMAL_TYPE,
			BaseAvatarId: lineAvatar.AvatarId,
		}
		if lineAvatar.LineAvatarType == spb.LineAvatarType_LineAvatarType_TRIAL {
			actor.AvatarType = proto.AvatarType_AVATAR_TRIAL_TYPE
		}
		entityList := &proto.SceneEntityRefreshInfo{
			AddEntity: &proto.SceneEntityInfo{
				Actor: actor,
				Motion: &proto.MotionInfo{
					Pos: pos,
					Rot: rot,
				},
				EntityId: g.GetNextGameObjectGuid(),
			},
		}
		g.AddEntity(0, &AvatarEntity{
			Entity: Entity{
				EntityId: entityList.AddEntity.EntityId,
				GroupId:  0,
				Pos:      pos,
				Rot:      rot,
			},
			AvatarId: lineAvatar.AvatarId,
		})
		sceneEntityRefreshInfo = append(sceneEntityRefreshInfo, entityList)
	}
	return sceneEntityRefreshInfo
}

// 添加Buff
func (g *GamePlayer) GetAddBuffSceneEntityRefreshInfo(casterEntityId, summonId uint32, pos, rot *proto.Vector) []*proto.GroupRefreshInfo {
	groupRefreshInfo := make([]*proto.GroupRefreshInfo, 0)
	sceneGroupRefreshInfo := &proto.GroupRefreshInfo{
		RefreshEntity: make([]*proto.SceneEntityRefreshInfo, 0),
	}
	// for _, lineAvatar := range lineUp.AvatarIdList {
	sceneEntityRefreshInfo := &proto.SceneEntityRefreshInfo{
		AddEntity: &proto.SceneEntityInfo{
			Motion: &proto.MotionInfo{
				Pos: pos,
				Rot: &proto.Vector{Y: 139439},
			},
			EntityId: g.GetNextGameObjectGuid(),
			SummonUnit: &proto.SceneSummonUnitInfo{
				CasterEntityId:  casterEntityId,
				AttachEntityId:  casterEntityId,
				SummonUnitId:    summonId,
				CreateTimeMs:    uint64(time.Now().UnixMilli()),
				TriggerNameList: make([]string, 0),
				LifeTimeMs:      -1,
			},
		},
	}
	// g.AddEntity(0, &AvatarEntity{
	// 	Entity: Entity{
	// 		EntityId: entityList.AddEntity.EntityId,
	// 		GroupId:  0,
	// 		Pos:      pos,
	// 		Rot:      rot,
	// 	},
	// 	AvatarId: lineAvatar.AvatarId,
	// })
	sceneGroupRefreshInfo.RefreshEntity = append(sceneGroupRefreshInfo.RefreshEntity, sceneEntityRefreshInfo)
	// }
	groupRefreshInfo = append(groupRefreshInfo, sceneGroupRefreshInfo)
	return groupRefreshInfo
}

func (g *GamePlayer) GetSceneGroupRefreshInfoByLineUP(lineUp *spb.Line, pos, rot *proto.Vector) []*proto.GroupRefreshInfo {
	groupRefreshInfo := make([]*proto.GroupRefreshInfo, 0)
	sceneGroupRefreshInfo := &proto.GroupRefreshInfo{
		RefreshEntity: make([]*proto.SceneEntityRefreshInfo, 0),
	}
	for _, lineAvatar := range lineUp.AvatarIdList {
		actor := &proto.SceneActorInfo{
			AvatarType:   proto.AvatarType_AVATAR_FORMAL_TYPE,
			BaseAvatarId: lineAvatar.AvatarId,
		}
		if lineAvatar.LineAvatarType == spb.LineAvatarType_LineAvatarType_TRIAL {
			actor.AvatarType = proto.AvatarType_AVATAR_TRIAL_TYPE
		}
		entityId := g.GetNextGameObjectGuid()
		sceneEntityRefreshInfo := &proto.SceneEntityRefreshInfo{
			AddEntity: &proto.SceneEntityInfo{
				Actor: actor,
				Motion: &proto.MotionInfo{
					Pos: pos,
					Rot: rot,
				},
				EntityId: entityId,
			},
		}
		g.AddEntity(0, &AvatarEntity{
			Entity: Entity{
				EntityId: entityId,
				GroupId:  0,
				Pos:      pos,
				Rot:      rot,
			},
			AvatarId: lineAvatar.AvatarId,
		})
		sceneGroupRefreshInfo.RefreshEntity = append(sceneGroupRefreshInfo.RefreshEntity, sceneEntityRefreshInfo)
	}
	groupRefreshInfo = append(groupRefreshInfo, sceneGroupRefreshInfo)
	return groupRefreshInfo
}

// 获取忘却之庭世界
func (g *GamePlayer) GetChallengeScene() *proto.SceneInfo {
	curChallenge := g.GetCurChallenge()
	leaderEntityId := g.GetNextGameObjectGuid()
	lineUp := g.GetChallengesLineUp()
	mazeGroupID := g.GetChallengesMazeGroupID()
	configList := g.GetChallengesConfigList()
	npcMonsterIDList := g.GetChallengesNpcMonsterIDList()
	eventIDList := g.GetChallengesEventIDList()
	challengeMazeConfig := gdconf.GetChallengeMazeConfigById(curChallenge.ChallengeId)
	if challengeMazeConfig == nil {
		return nil
	}
	mapEntrance := gdconf.GetMapEntranceById(challengeMazeConfig.MapEntranceID)
	if mapEntrance == nil {
		return nil
	}
	foorMap := gdconf.GetServerGroupById(mapEntrance.PlaneID, mapEntrance.FloorID, mazeGroupID)
	if foorMap == nil || lineUp == nil || len(npcMonsterIDList) != len(eventIDList) || len(eventIDList) != len(configList) {
		return nil
	}
	pos, rot := g.GetChallengesAnchor(foorMap.AnchorList)
	if pos == nil || rot == nil {
		return nil
	}
	// 获取映射信息
	scene := &proto.SceneInfo{
		ClientPosVersion:   0,
		PlaneId:            mapEntrance.PlaneID,
		FloorId:            mapEntrance.FloorID,
		LeaderEntityId:     leaderEntityId,
		WorldId:            gdconf.GetMazePlaneById(mapEntrance.PlaneID).WorldID,
		EntryId:            challengeMazeConfig.MapEntranceID,
		GameModeType:       gdconf.GetPlaneType(gdconf.GetMazePlaneById(mapEntrance.PlaneID).PlaneType),
		EntityGroupList:    make([]*proto.SceneEntityGroupInfo, 0),
		SyncBuffInfo:       make([]*proto.BuffInfo, 0),
		GroupIdList:        make([]uint32, 0),
		GroupStateList:     make([]*proto.SceneGroupState, 0),
		LightenSectionList: []uint32{0},
		EntityList:         make([]*proto.SceneEntityInfo, 0),
	}

	// 添加场景buff
	if curChallenge.MazeBuffId != 0 {
		scene.SyncBuffInfo = append(scene.SyncBuffInfo, &proto.BuffInfo{
			Count:     4294967295,
			LifeTime:  -1,
			BuffId:    curChallenge.MazeBuffId,
			AddTimeMs: uint64(time.Now().UnixMilli()),
			Level:     1,
		})
	}
	// 添加自选buff
	if g.GetCurChallengeBuffId() != 0 {
		scene.SyncBuffInfo = append(scene.SyncBuffInfo, &proto.BuffInfo{
			Count:     4294967295,
			LifeTime:  -1,
			BuffId:    g.GetCurChallengeBuffId(),
			AddTimeMs: uint64(time.Now().UnixMilli()),
			Level:     1,
		})
	}
	// 将进入场景的角色添加到实体列表里
	entityGroup := &proto.SceneEntityGroupInfo{
		GroupId:    0,
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	g.GetSceneAvatarByLineUP(entityGroup, lineUp, leaderEntityId, pos, rot)
	scene.EntityGroupList = append(scene.EntityGroupList, entityGroup)
	// 添加怪物实体
	monsterEntityGroup := &proto.SceneEntityGroupInfo{
		GroupId:    mazeGroupID,
		EntityList: make([]*proto.SceneEntityInfo, 0),
	}
	for id, config := range configList {
		for _, monsterList := range foorMap.MonsterList {
			if monsterList.ID != config {
				continue
			}
			entityId := g.GetNextGameObjectGuid()
			monsterPos := &proto.Vector{
				X: int32(monsterList.PosX * 1000),
				Y: int32(monsterList.PosY * 1000),
				Z: int32(monsterList.PosZ * 1000),
			}
			monsterRot := &proto.Vector{
				X: int32(monsterList.RotX * 1000),
				Y: int32(monsterList.RotY * 1000),
				Z: int32(monsterList.RotZ * 1000),
			}
			entityList := &proto.SceneEntityInfo{
				GroupId:  mazeGroupID,
				InstId:   monsterList.ID,
				EntityId: entityId,
				Motion: &proto.MotionInfo{
					Pos: monsterPos,
					Rot: monsterRot,
				},
				NpcMonster: &proto.SceneNpcMonsterInfo{
					MonsterId: npcMonsterIDList[id],
					EventId:   eventIDList[id],
				},
			}
			// 添加怪物实体
			g.AddEntity(mazeGroupID, &MonsterEntity{
				Entity: Entity{
					EntityId: entityId,
					GroupId:  mazeGroupID,
					Pos:      monsterPos,
					Rot:      monsterRot,
					InstId:   monsterList.ID,
				},
				EventID: eventIDList[id],
			})
			monsterEntityGroup.EntityList = append(monsterEntityGroup.EntityList, entityList)
		}
	}
	scene.EntityGroupList = append(scene.EntityGroupList, monsterEntityGroup)
	return scene
}

func (g *GamePlayer) GetSpringRecoverConfig() *proto.SpringRecoverConfig {
	info := &proto.SpringRecoverConfig{
		RecoverPct:  10000,
		AutoRecover: true,
		// MANPHKHEFPC: nil,
	}
	return info
}

func (g *GamePlayer) GetHealPoolInfo() *proto.HealPoolInfo {
	info := &proto.HealPoolInfo{
		RefreshTime: time.Now().Unix(),
		HealPool:    23500,
	}

	return info
}
