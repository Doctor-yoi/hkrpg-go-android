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

type RelicExpType struct {
	ExpType uint32 `json:"ExpType"`
	Level   uint32 `json:"Level"`
	Exp     uint32 `json:"Exp"`
}

func (g *GameDataConfig) loadRelicExpType() {
	g.RelicExpTypeMap = make(map[string]map[string]*RelicExpType)
	//playerElementsFilePath := g.excelPrefix + "RelicExpType.json"
	//playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFile, err := base64.StdEncoding.DecodeString(gameData.RelicExpType)
	if err != nil {
		logger.Error("get RelicExpType error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.RelicExpTypeMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v RelicExpType", len(g.RelicExpTypeMap))
}

func GetRelicExpByLevel(relicType, exp, level, relicId uint32) (uint32, uint32) {
	maxLevel := GetRelicMaxLevel(relicId)
	for ; level <= maxLevel; level++ {
		if exp < CONF.RelicExpTypeMap[strconv.Itoa(int(relicType))][strconv.Itoa(int(level))].Exp {
			return level, exp
		} else {
			exp -= CONF.RelicExpTypeMap[strconv.Itoa(int(relicType))][strconv.Itoa(int(level))].Exp
		}
	}
	newExp := CONF.RelicExpTypeMap[strconv.Itoa(int(relicType))][strconv.Itoa(int(maxLevel))].Exp
	return maxLevel, newExp
}
