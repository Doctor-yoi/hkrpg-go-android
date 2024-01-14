package gdconf

import (
	"encoding/base64"
	"fmt"
	"hkrpg/gameData"
	"os"

	"github.com/hjson/hjson-go/v4"
	"hkrpg/pkg/logger"
)

type ItemConfigRelic struct {
	ID                  uint32              `json:"ID"`
	ItemMainType        string              `json:"ItemMainType"`
	ItemSubType         string              `json:"ItemSubType"`
	InventoryDisplayTag uint32              `json:"InventoryDisplayTag"`
	Rarity              string              `json:"Rarity"`
	IsVisible           bool                `json:"isVisible"`
	PileLimit           uint32              `json:"PileLimit"`
	IsSellable          bool                `json:"IsSellable"`
	ReturnItemIDList    []*ReturnItemIDList `json:"ReturnItemIDList"`
	SellType            string              `json:"SellType"`
}

func (g *GameDataConfig) loadItemConfigRelic() {
	g.ItemConfigRelicMap = make(map[string]*ItemConfigRelic)
	//playerElementsFilePath := g.excelPrefix + "ItemConfigRelic.json"
	//playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFile, err := base64.StdEncoding.DecodeString(gameData.ItemConfigRelic)
	if err != nil {
		logger.Error("get ItemConfigRelic error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.ItemConfigRelicMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v ItemConfigRelic", len(g.ItemConfigRelicMap))
}

func GetItemConfigRelicById(ID string) *ItemConfigRelic {
	return CONF.ItemConfigRelicMap[ID]
}

func GetItemConfigRelicMap() map[string]*ItemConfigRelic {
	return CONF.ItemConfigRelicMap
}
