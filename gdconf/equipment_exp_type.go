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

type EquipmentExp struct {
	ExpType uint32 `json:"ExpType"`
	Level   uint32 `json:"Level"`
	Exp     uint32 `json:"Exp"`
}

func (g *GameDataConfig) loadEquipmentExpType() {
	g.EquipmentExpTypeMap = make(map[string]map[string]*EquipmentExp)
	//playerElementsFilePath := g.excelPrefix + "EquipmentExpType.json"
	//playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFile, err := base64.StdEncoding.DecodeString(gameData.EquipmentExpType)
	if err != nil {
		logger.Error("get EquipmentExpType error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.EquipmentExpTypeMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v EquipmentExpType", len(g.EquipmentExpTypeMap))
}

func GetEquipmentExpByLevel(equipmentType, exp, level, promotion, equipmentId uint32) (uint32, uint32) {
	maxLevel := GetEquipmentMaxLevel(equipmentId, promotion)
	for ; level <= maxLevel; level++ {
		if exp < CONF.EquipmentExpTypeMap[strconv.Itoa(int(equipmentType))][strconv.Itoa(int(level))].Exp {
			return level, exp
		} else {
			exp -= CONF.EquipmentExpTypeMap[strconv.Itoa(int(equipmentType))][strconv.Itoa(int(level))].Exp
		}
	}
	newExp := CONF.EquipmentExpTypeMap[strconv.Itoa(int(equipmentType))][strconv.Itoa(int(maxLevel))].Exp
	return maxLevel, newExp
}
