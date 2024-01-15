package gdconf

import (
	"encoding/base64"
	"fmt"
	"hkrpg/gameData"
	"os"

	"github.com/hjson/hjson-go/v4"
	"hkrpg/pkg/logger"
)

type MonsterConfig struct {
	MonsterID         uint32 `json:"MonsterID"`
	MonsterTemplateID uint32 `json:"MonsterTemplateID"`
	HardLevelGroup    uint32 `json:"HardLevelGroup"`
	EliteGroup        uint32 `json:"EliteGroup"`
	// TODO 需要再加
}

func (g *GameDataConfig) loadMonsterConfig() {
	g.MonsterConfigMap = make(map[string]*MonsterConfig)
	//playerElementsFilePath := g.excelPrefix + "MonsterConfig.json"
	//playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFile, err := base64.StdEncoding.DecodeString(gameData.MonsterConfig)
	if err != nil {
		logger.Error("get MonsterConfig error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.MonsterConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v MonsterConfig", len(g.MonsterConfigMap))
}

func GetMonsterConfigById(monsterID string) *MonsterConfig {
	return CONF.MonsterConfigMap[monsterID]
}

func GetMonsterConfigMap() map[string]*MonsterConfig {
	return CONF.MonsterConfigMap
}
