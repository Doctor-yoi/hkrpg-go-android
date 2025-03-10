package gdconf

import (
	"encoding/base64"
	"fmt"
	"hkrpg/gameData"
	"os"

	"github.com/hjson/hjson-go/v4"
	"hkrpg/pkg/logger"
)

type EquipmentConfig struct {
	EquipmentID          uint32    `json:"EquipmentID"`
	Release              bool      `json:"Release"`
	Rarity               string    `json:"Rarity"`
	AvatarBaseType       string    `json:"AvatarBaseType"`
	SkillID              uint32    `json:"SkillID"`
	ExpType              uint32    `json:"ExpType"`
	ExpProvide           uint32    `json:"ExpProvide"`
	CoinCost             uint32    `json:"CoinCost"`
	ItemRightPanelOffset []float64 `json:"ItemRightPanelOffset"`
	AvatarDetailOffset   []float64 `json:"AvatarDetailOffset"`
	BattleDialogOffset   []float64 `json:"BattleDialogOffset"`
	GachaResultOffset    []float64 `json:"GachaResultOffset"`
}

func (g *GameDataConfig) loadEquipmentConfig() {
	g.EquipmentConfigMap = make(map[string]*EquipmentConfig)
	//playerElementsFilePath := g.excelPrefix + "EquipmentConfig.json"
	//playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFile, err := base64.StdEncoding.DecodeString(gameData.EquipmentConfig)
	if err != nil {
		logger.Error("get EquipmentConfig error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.EquipmentConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	//playerElementsFilePaths := g.excelPrefix + "EquipmentExpItemConfig.json"
	//playerElementsFiles, err := os.ReadFile(playerElementsFilePaths)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFiles, err := base64.StdEncoding.DecodeString(gameData.EquipmentExpItemConfig)
	if err != nil {
		logger.Error("get EquipmentExpItemConfig error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFiles, &g.EquipmentConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	logger.Info("load %v EquipmentConfig", len(g.EquipmentConfigMap))
}

func GetEquipmentConfigById(ID string) *EquipmentConfig {
	return CONF.EquipmentConfigMap[ID]
}

func GetEquipmentConfigMap() map[string]*EquipmentConfig {
	return CONF.EquipmentConfigMap
}

func GetEquipmentList() []uint32 {
	var equipmentList []uint32
	for _, equipment := range CONF.EquipmentConfigMap {
		equipmentList = append(equipmentList, equipment.EquipmentID)
	}
	return equipmentList
}
