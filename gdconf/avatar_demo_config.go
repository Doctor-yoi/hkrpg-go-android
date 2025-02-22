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

type AvatarDemoConfig struct {
	StageID           uint32   `json:"StageID"`
	AvatarID          uint32   `json:"AvatarID"`
	TrialAvatarList   []uint32 `json:"TrialAvatarList"`
	RewardID          uint32   `json:"RewardID"`
	RaidID            uint32   `json:"RaidID"`
	ScoringGroupID    uint32   `json:"ScoringGroupID"`
	GuideGroupID      uint32   `json:"GuideGroupID"`
	PlaneID           uint32   `json:"PlaneID"`
	FloorID           uint32   `json:"FloorID"`
	BattleAreaGroupID uint32   `json:"BattleAreaGroupID"`
	BattleAreaID      uint32   `json:"BattleAreaID"`
	MapEntranceID     uint32   `json:"MapEntranceID"`
	MazeGroupID1      uint32   `json:"MazeGroupID1"`
	ConfigList1       []uint32 `json:"ConfigList1"`
	NpcMonsterIDList1 []uint32 `json:"NpcMonsterIDList1"`
	EventIDList1      []uint32 `json:"EventIDList1"`
}

func (g *GameDataConfig) loadAvatarDemoConfig() {
	g.AvatarDemoConfigMap = make(map[string]*AvatarDemoConfig)
	//playerElementsFilePath := g.excelPrefix + "AvatarDemoConfig.json"
	//playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFile, err := base64.StdEncoding.DecodeString(gameData.AvatarDemoConfig)
	if err != nil {
		logger.Error("get AvatarDemoConfig error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.AvatarDemoConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v AvatarDemoConfig", len(g.AvatarDemoConfigMap))
}

func GetAvatarDemoConfigById(stageID uint32) *AvatarDemoConfig {
	return CONF.AvatarDemoConfigMap[strconv.Itoa(int(stageID))]
}
