package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type NPCData struct {
	ID               uint32 `json:"ID"`
	ConfigEntityPath string `json:"ConfigEntityPath"`
	JsonPath         string `json:"JsonPath"`
	SubType          string `json:"SubType"`
	AnimGroupID      uint32 `json:"AnimGroupID"`
	SeriesID         uint32 `json:"SeriesID"`
}

func (g *GameDataConfig) loadNPCData() {
	g.NPCDataMap = make(map[uint32]*NPCData)
	nPCDataMap := make([]*NPCData, 0)
	playerElementsFilePath := g.excelPrefix + "NPCData.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &nPCDataMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range nPCDataMap {
		g.NPCDataMap[v.ID] = v
	}
	logger.Info("load %v NPCData", len(g.NPCDataMap))

}

func GetNPCDataId(id uint32) *NPCData {
	return CONF.NPCDataMap[id]
}
