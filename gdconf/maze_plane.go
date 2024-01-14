package gdconf

import (
	"encoding/base64"
	"fmt"
	"hkrpg/gameData"
	"os"

	"github.com/hjson/hjson-go/v4"
	"hkrpg/pkg/logger"
)

type MazePlane struct {
	PlaneID      uint32   `json:"PlaneID"`
	PlaneType    string   `json:"PlaneType"`
	MazePoolType uint32   `json:"MazePoolType"`
	WorldID      uint32   `json:"WorldID"`
	StartFloorID uint32   `json:"StartFloorID"`
	FloorIDList  []uint32 `json:"FloorIDList"`
}

func (g *GameDataConfig) loadMazePlane() {
	g.MazePlaneMap = make(map[string]*MazePlane)
	//playerElementsFilePath := g.excelPrefix + "MazePlane.json"
	//playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFile, err := base64.StdEncoding.DecodeString(gameData.MazePlane)
	if err != nil {
		logger.Error("get MazePlane error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.MazePlaneMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v MazePlane", len(g.MazePlaneMap))
}

func GetMazePlaneById(planeID string) *MazePlane {
	return CONF.MazePlaneMap[planeID]
}

func GetMazePlaneMap() map[string]*MazePlane {
	return CONF.MazePlaneMap
}

func GetPlaneType(state string) uint32 {
	stateMap := map[string]uint32{
		"Unknown":       0,
		"Town":          1,
		"Maze":          2,
		"Train":         3,
		"Challenge":     4,
		"Rogue":         5,
		"Raid":          6,
		"AetherDivide":  7,
		"TrialActivity": 8,
	}

	value, ok := stateMap[state]
	if !ok {
		return 0
	}

	return value
}
