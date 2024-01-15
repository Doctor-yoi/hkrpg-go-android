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

type ChallengeTargetConfig struct {
	ID                    uint32 `json:"ID"`
	ChallengeTargetType   string `json:"ChallengeTargetType"`
	ChallengeTargetParam1 uint32 `json:"ChallengeTargetParam1"`
	RewardID              uint32 `json:"RewardID"`
}

func (g *GameDataConfig) loadChallengeTargetConfig() {
	g.ChallengeTargetConfigMap = make(map[string]*ChallengeTargetConfig)
	//playerElementsFilePath := g.excelPrefix + "ChallengeTargetConfig.json"
	//playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFile, err := base64.StdEncoding.DecodeString(gameData.ChallengeTargetConfig)
	if err != nil {
		logger.Error("get ChallengeTargetConfig error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.ChallengeTargetConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	playerElementsFileStory, err := base64.StdEncoding.DecodeString(gameData.ChallengeStoryTargetConfig)
	if err != nil {
		logger.Error("get ChallengeStoryTargetConfig error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFileStory, &g.ChallengeTargetConfigMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	logger.Info("load %v ChallengeTargetConfig", len(g.ChallengeTargetConfigMap))
}

func GetChallengeTargetConfigById(id uint32) *ChallengeTargetConfig {
	return CONF.ChallengeTargetConfigMap[strconv.Itoa(int(id))]
}
