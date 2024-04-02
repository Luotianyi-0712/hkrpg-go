package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type RogueRoom struct {
	RogueRoomID       uint32            `json:"RogueRoomID"`
	RogueRoomType     uint32            `json:"RogueRoomType"`
	MapEntrance       uint32            `json:"MapEntrance"`
	GroupID           uint32            `json:"GroupID"`
	GroupWithContent  map[string]uint32 `json:"GroupWithContent"`
	RogueRoomSections []uint32          `json:"RogueRoomSections"`
}

func (g *GameDataConfig) loadRogueRoom() {
	g.RogueRoomMap = make(map[uint32]*RogueRoom)
	rogueRoomMap := make(map[string]*RogueRoom)
	playerElementsFilePath := g.excelPrefix + "RogueRoom.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueRoomMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	for _, rogueRoom := range rogueRoomMap {
		g.RogueRoomMap[rogueRoom.RogueRoomID] = rogueRoom
	}

	logger.Info("load %v RogueRoom", len(g.RogueRoomMap))
}

func GetRogueRoomById(roomId uint32) *RogueRoom {
	return CONF.RogueRoomMap[roomId]
}
