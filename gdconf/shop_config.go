package gdconf

import (
	"encoding/base64"
	"fmt"
	"hkrpg/gameData"
	"os"

	"github.com/hjson/hjson-go/v4"
	"hkrpg/pkg/logger"
)

type ShopConfig struct {
	ShopID          uint32   `json:"ShopID"`
	ShopMainType    string   `json:"ShopMainType"`
	ShopType        uint32   `json:"ShopType"`
	ShopBar         string   `json:"ShopBar"`
	ShopSortID      uint32   `json:"ShopSortID"`
	LimitType1      string   `json:"LimitType1"`
	LimitValue1List []uint32 `json:"LimitValue1List"`
	LimitValue2List []uint32 `json:"LimitValue2List"`
	IsOpen          bool     `json:"IsOpen"`
	ScheduleDataID  uint32   `json:"ScheduleDataID"`
	HideRemainTime  bool     `json:"HideRemainTime"`
}

func (g *GameDataConfig) loadShopConfig() {
	g.ShopConfigMap = make(map[uint32][]uint32)
	shopConfigMap := make(map[string]*ShopConfig)
	//playerElementsFilePath := g.excelPrefix + "ShopConfig.json"
	//playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFile, err := base64.StdEncoding.DecodeString(gameData.ShopConfig)
	if err != nil {
		logger.Error("get ShopConfig error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFile, &shopConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	for _, shopGoodsConfig := range shopConfigMap {
		if g.ShopConfigMap[shopGoodsConfig.ShopType] == nil {
			g.ShopConfigMap[shopGoodsConfig.ShopType] = make([]uint32, 0)
		}
		g.ShopConfigMap[shopGoodsConfig.ShopType] = append(g.ShopConfigMap[shopGoodsConfig.ShopType], shopGoodsConfig.ShopID)
	}

	logger.Info("load %v ShopConfig", len(g.ShopConfigMap))
}

func GetShopConfigByTypeId(typeId uint32) []uint32 {
	return CONF.ShopConfigMap[typeId]
}
