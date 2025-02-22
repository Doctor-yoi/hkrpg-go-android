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

type StageConfig struct {
	StageID          uint32              `json:"StageID"`          // 具体怪物id群
	StageType        string              `json:"StageType"`        // 怪物类型
	HardLevelGroup   uint32              `json:"HardLevelGroup"`   // 强度等级
	MonsterList      []map[string]uint32 `json:"MonsterList"`      // 怪物id
	ForbidExitBattle bool                `json:"ForbidExitBattle"` // 禁止退出
}

func (g *GameDataConfig) loadStageConfig() {
	g.StageConfigMap = make(map[string]*StageConfig)
	//playerElementsFilePath := g.excelPrefix + "StageConfig.json"
	//playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFile, err := base64.StdEncoding.DecodeString(gameData.StageConfig)
	if err != nil {
		logger.Error("get StageConfig error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.StageConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	logger.Info("load %v StageConfig", len(g.StageConfigMap))
}

func GetStageConfigById(stageID uint32) *StageConfig {
	return CONF.StageConfigMap[strconv.Itoa(int(stageID))]
}
