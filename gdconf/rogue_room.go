package gdconf

import (
	"encoding/base64"
	"fmt"
	"hkrpg/gameData"
	"os"

	"github.com/hjson/hjson-go/v4"
	"hkrpg/pkg/logger"
)

type RogueRoom struct {
	RogueRoomID       uint32            `json:"RogueRoomID"`
	RogueRoomType     uint32            `json:"RogueRoomType"`
	MapEntrance       uint32            `json:"MapEntrance"`
	GroupID           uint32            `json:"GroupID"`
	GroupWithContent  map[string]uint32 `json:"GroupWithContent"`
	RogueRoomSections []uint32          `json:"RogueRoomSections"`
}

func (g *GameDataConfig) loadRogueRoom() {
	g.RogueRoomMap = make(map[string]*RogueRoom)
	//playerElementsFilePath := g.excelPrefix + "RogueRoom.json"
	//playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFile, err := base64.StdEncoding.DecodeString(gameData.RogueRoom)
	if err != nil {
		logger.Error("get RogueRoom error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.RogueRoomMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v RogueRoom", len(g.RogueRoomMap))
}

func GetRogueRoomById(roomId string) *RogueRoom {
	return CONF.RogueRoomMap[roomId]
}
