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

type ExpType struct {
	ExpType uint32 `json:"ExpType"`
	Level   uint32 `json:"Level"`
	Exp     uint32 `json:"Exp"`
}

func (g *GameDataConfig) loadExpType() {
	g.ExpTypeMap = make(map[string]map[string]*ExpType)
	//playerElementsFilePath := g.excelPrefix + "ExpType.json"
	//playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFile, err := base64.StdEncoding.DecodeString(gameData.ExpType)
	if err != nil {
		logger.Error("get ExpType error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.ExpTypeMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v ExpType", len(g.ExpTypeMap))
}

func GetExpTypeByLevel(expType, exp, level, promotion, avatarId uint32) (uint32, uint32, uint32) {
	maxLevel := GetAvatarMaxLevel(avatarId, promotion)
	for ; level <= maxLevel; level++ {
		if exp < CONF.ExpTypeMap[strconv.Itoa(int(expType))][strconv.Itoa(int(level))].Exp {
			return level, exp, 0
		} else {
			exp -= CONF.ExpTypeMap[strconv.Itoa(int(expType))][strconv.Itoa(int(level))].Exp
		}
	}
	newExp := CONF.ExpTypeMap[strconv.Itoa(int(expType))][strconv.Itoa(int(maxLevel))].Exp
	return maxLevel, newExp, exp - newExp
}
