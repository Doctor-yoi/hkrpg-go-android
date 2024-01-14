package gdconf

import (
	"encoding/base64"
	"fmt"
	"hkrpg/gameData"
	"os"

	"github.com/hjson/hjson-go/v4"
	"hkrpg/pkg/logger"
)

type TextJoinConfig struct {
	TextJoinID       uint32   `json:"TextJoinID"`
	DefaultItem      uint32   `json:"DefaultItem"`
	TextJoinItemList []uint32 `json:"TextJoinItemList"`
}

func (g *GameDataConfig) loadTextJoinConfig() {
	g.TextJoinConfigMap = make(map[string]*TextJoinConfig)
	//playerElementsFilePath := g.excelPrefix + "TextJoinConfig.json"
	//playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFile, err := base64.StdEncoding.DecodeString(gameData.TextJoinConfigs)
	if err != nil {
		logger.Error("get TextJoinConfig error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.TextJoinConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v TextJoinConfig", len(g.TextJoinConfigMap))
}

func GetTextJoinConfigById(ID string) *TextJoinConfig {
	return CONF.TextJoinConfigMap[ID]
}

func GetTextJoinConfigMap() map[string]*TextJoinConfig {
	return CONF.TextJoinConfigMap
}
