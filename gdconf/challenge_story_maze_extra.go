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

type ChallengeStoryMazeExtra struct {
	ID             uint32   `json:"ID"`
	TurnLimit      uint32   `json:"TurnLimit"`
	BattleTargetID []uint32 `json:"BattleTargetID"`
	ClearScore     uint32   `json:"ClearScore"`
}

func (g *GameDataConfig) loadChallengeStoryMazeExtra() {
	g.ChallengeStoryMazeExtraMap = make(map[string]*ChallengeStoryMazeExtra, 0)
	//playerElementsFilePath := g.excelPrefix + "ChallengeStoryMazeExtra.json"
	//playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFile, err := base64.StdEncoding.DecodeString(gameData.ChallengeStoryMazeExtra)
	if err != nil {
		logger.Error("get ChallengeStoryMazeExtra error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.ChallengeStoryMazeExtraMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	logger.Info("load %v ChallengeStoryMazeExtra", len(g.ChallengeStoryMazeExtraMap))
}

func GetChallengeStoryMazeExtraById(id uint32) *ChallengeStoryMazeExtra {
	return CONF.ChallengeStoryMazeExtraMap[strconv.Itoa(int(id))]
}
