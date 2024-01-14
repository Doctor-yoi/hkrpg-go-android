package gdconf

import (
	"encoding/base64"
	"fmt"
	"hkrpg/gameData"
	"os"

	"github.com/hjson/hjson-go/v4"
	"hkrpg/pkg/logger"
)

type QuestData struct {
	QuestID     uint32 `json:"QuestID"`
	QuestType   uint32 `json:"QuestType"`
	UnlockType  string `json:"UnlockType"`
	RewardID    uint32 `json:"RewardID"`
	FinishWayID uint32 `json:"FinishWayID"`
	GotoID      uint32 `json:"GotoID"`
}

func (g *GameDataConfig) loadQuestData() {
	g.QuestDataMap = make(map[string]*QuestData)
	//playerElementsFilePath := g.excelPrefix + "QuestData.json"
	//playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFile, err := base64.StdEncoding.DecodeString(gameData.QuestData)
	if err != nil {
		logger.Error("get QuestData error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.QuestDataMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v QuestData", len(g.QuestDataMap))
}

func GetQuestDataById(questID string) *QuestData {
	return CONF.QuestDataMap[questID]
}

func GetQuestDataMap() map[string]*QuestData {
	return CONF.QuestDataMap
}
