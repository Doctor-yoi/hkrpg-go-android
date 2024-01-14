package gdconf

import (
	"encoding/base64"
	"fmt"
	"hkrpg/gameData"
	"os"
	"strconv"

	"github.com/hjson/hjson-go/v4"
	"hkrpg/pkg/logger"
)

type ActivityLoginConfig struct {
	ID               uint32   `json:"ID"`
	RewardList       []uint32 `json:"RewardList"`
	ActivityModuleID uint32   `json:"ActivityModuleID"`
}

func (g *GameDataConfig) loadActivityLoginConfig() {
	g.ActivityLoginConfigMap = make(map[string]*ActivityLoginConfig)
	//playerElementsFilePath := g.excelPrefix + "ActivityLoginConfig.json"
	//playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFile, err := base64.StdEncoding.DecodeString(gameData.ActivityLoginConfig)
	if err != nil {
		logger.Error("get ActivityLoginConfig error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.ActivityLoginConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v ActivityLoginConfig", len(g.ActivityLoginConfigMap))
}

func GetActivityLoginConfigById(id uint32) *ActivityLoginConfig {
	return CONF.ActivityLoginConfigMap[strconv.Itoa(int(id))]
}

func GetActivityLoginListById() []uint32 {
	var activityLoginList []uint32
	for _, conf := range CONF.ActivityLoginConfigMap {
		activityLoginList = append(activityLoginList, conf.ID)
	}
	return activityLoginList
}
