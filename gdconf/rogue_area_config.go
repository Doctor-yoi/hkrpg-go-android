package gdconf

import (
	"encoding/base64"
	"fmt"
	"hkrpg/gameData"
	"os"

	"github.com/hjson/hjson-go/v4"
	"hkrpg/pkg/logger"
)

type RogueAreaConfig struct {
	RogueAreaID              uint32            `json:"RogueAreaID"`
	AreaProgress             uint32            `json:"AreaProgress"`
	Difficulty               uint32            `json:"Difficulty"`
	RecommendLevel           uint32            `json:"RecommendLevel"`
	RecommendNature          []string          `json:"RecommendNature"`
	AreaIcon                 string            `json:"AreaIcon"`
	AreaFigure               string            `json:"AreaFigure"`
	DisplayMonsterMap        map[string]uint32 `json:"DisplayMonsterMap"`
	DisplayMonsterMap2       map[string]uint32 `json:"DisplayMonsterMap2"`
	FirstReward              uint32            `json:"FirstReward"`
	UnlockID                 uint32            `json:"UnlockID"`
	MapDisplayItemList       []*RewardList     `json:"MapDisplayItemList"`
	ChestDisplayItemList     []*RewardList     `json:"ChestDisplayItemList"`
	MonsterDisplayItemList   []*RewardList     `json:"MonsterDisplayItemList"`
	ScoreMap                 map[string]uint32 `json:"ScoreMap"`
	RecommendSkillTreePoints uint32            `json:"RecommendSkillTreePoints"`
	AreaTipsIcon             string            `json:"AreaTipsIcon"`
}

func (g *GameDataConfig) loadRogueAreaConfig() {
	g.RogueAreaConfigMap = make(map[string]*RogueAreaConfig)
	//playerElementsFilePath := g.excelPrefix + "RogueAreaConfig.json"
	//playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFile, err := base64.StdEncoding.DecodeString(gameData.RogueAreaConfig)
	if err != nil {
		logger.Error("get RogueAreaConfig error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.RogueAreaConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v RogueAreaConfig", len(g.RogueAreaConfigMap))
}

func GetRogueAreaMap() map[string]*RogueAreaConfig {
	return CONF.RogueAreaConfigMap
}
func GetRogueAreaConfigById(AreaID string) *RogueAreaConfig {
	return CONF.RogueAreaConfigMap[AreaID]
}
