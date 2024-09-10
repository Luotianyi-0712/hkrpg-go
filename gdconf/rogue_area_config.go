package gdconf

import (
	"fmt"
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type RogueAreaConfig struct {
	RogueAreaID              uint32            `json:"RogueAreaID"`
	AreaProgress             uint32            `json:"AreaProgress"`
	Difficulty               uint32            `json:"Difficulty"`
	RecommendLevel           uint32            `json:"RecommendLevel"`
	RecommendNature          []string          `json:"RecommendNature"`
	AreaIcon                 string            `json:"AreaIcon"`
	AreaFigure               string            `json:"AreaFigure"`
	DisplayMonsterMap        map[string]uint32 `json:"DisplayMonsterMap"`
	DisplayMonsterMap2       map[string]uint32 `json:"DisplayMonsterMap2"`
	FirstReward              uint32            `json:"FirstReward"`
	UnlockID                 uint32            `json:"UnlockID"`
	MapDisplayItemList       []*RewardList     `json:"MapDisplayItemList"`
	ChestDisplayItemList     []*RewardList     `json:"ChestDisplayItemList"`
	MonsterDisplayItemList   []*RewardList     `json:"MonsterDisplayItemList"`
	ScoreMap                 map[string]uint32 `json:"ScoreMap"` // 得分
	RecommendSkillTreePoints uint32            `json:"RecommendSkillTreePoints"`
	AreaTipsIcon             string            `json:"AreaTipsIcon"`
}

func (g *GameDataConfig) loadRogueAreaConfig() {
	g.RogueAreaConfigMap = make(map[uint32]*RogueAreaConfig)
	rogueAreaConfigMap := make([]*RogueAreaConfig, 0)
	playerElementsFilePath := g.excelPrefix + "RogueAreaConfig.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &rogueAreaConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, v := range rogueAreaConfigMap {
		g.RogueAreaConfigMap[v.RogueAreaID] = v
	}
	logger.Info("load %v RogueAreaConfig", len(g.RogueAreaConfigMap))
}

func GetRogueAreaMap() map[uint32]*RogueAreaConfig {
	return CONF.RogueAreaConfigMap
}
func GetRogueAreaConfigById(AreaID uint32) *RogueAreaConfig {
	return CONF.RogueAreaConfigMap[AreaID]
}
