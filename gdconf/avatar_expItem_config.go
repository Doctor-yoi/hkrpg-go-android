package gdconf

import (
	"encoding/base64"
	"fmt"
	"hkrpg/gameData"
	"os"

	"github.com/hjson/hjson-go/v4"
	"hkrpg/pkg/logger"
)

type AvatarExpItemConfig struct {
	ItemID uint32 `json:"ItemID"`
	Exp    uint32 `json:"Exp"`
}

func (g *GameDataConfig) loadAvatarExpItemConfig() {
	g.AvatarExpItemConfigMap = make(map[string]*AvatarExpItemConfig)
	//playerElementsFilePath := g.excelPrefix + "AvatarExpItemConfig.json"
	//playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFile, err := base64.StdEncoding.DecodeString(gameData.AvatarExpItemConfig)
	if err != nil {
		logger.Error("get AvatarExpItemConfig error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.AvatarExpItemConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v AvatarExpItemConfig", len(g.AvatarExpItemConfigMap))
}

func GetAvatarExpItemConfigById(itemID string) *AvatarExpItemConfig {
	return CONF.AvatarExpItemConfigMap[itemID]
}

func GetAvatarExpItemConfigMap() map[string]*AvatarExpItemConfig {
	return CONF.AvatarExpItemConfigMap
}
