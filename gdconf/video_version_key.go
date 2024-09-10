package gdconf

import (
	"os"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type VideoVersionKey struct {
	ActivityVideoKeyInfoList []*VideoKeyInfo `json:"activityVideoKeyInfoList"`
	VideoKeyInfoList         []*VideoKeyInfo `json:"videoKeyInfoList"`
}

type VideoKeyInfo struct {
	Id       uint32 `json:"id"`
	VideoKey uint64 `json:"videoKey"`
}

func (g *GameDataConfig) loadVideoVersionKey() {
	g.VideoVersionKey = new(VideoVersionKey)
	playerElementsFilePath := g.dataPrefix + "VideoVersionKey.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		logger.Error("open file error: %v", err)
		return
	}

	err = hjson.Unmarshal(playerElementsFile, &g.VideoVersionKey)
	if err != nil {
		logger.Error("parse file error: %v", err)
		return
	}
	logger.Info("load %v VideoVersionKey", len(g.VideoVersionKey.VideoKeyInfoList)+len(g.VideoVersionKey.ActivityVideoKeyInfoList))
}

func GetVideoVersionKey() *VideoVersionKey {
	return CONF.VideoVersionKey
}
