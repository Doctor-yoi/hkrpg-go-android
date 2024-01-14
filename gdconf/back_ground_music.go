package gdconf

import (
	"encoding/base64"
	"fmt"
	"hkrpg/gameData"
	"os"

	"github.com/hjson/hjson-go/v4"
	"hkrpg/pkg/logger"
)

type BackGroundMusic struct {
	ID              uint32 `json:"ID"`
	GroupID         uint32 `json:"GroupID"`
	MusicSwitchName string `json:"MusicSwitchName"`
	BPM             uint32 `json:"BPM"`
}

func (g *GameDataConfig) loadBackGroundMusic() {
	g.BackGroundMusicMap = make(map[string]*BackGroundMusic)
	//playerElementsFilePath := g.excelPrefix + "BackGroundMusic.json"
	//playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFile, err := base64.StdEncoding.DecodeString(gameData.BackGroundMusic)
	if err != nil {
		logger.Error("get BackGroundMusic error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.BackGroundMusicMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v BackGroundMusic", len(g.BackGroundMusicMap))
}

func GetBackGroundMusicById(iD string) *BackGroundMusic {
	return CONF.BackGroundMusicMap[iD]
}

func GetBackGroundMusicMap() map[string]*BackGroundMusic {
	return CONF.BackGroundMusicMap
}
