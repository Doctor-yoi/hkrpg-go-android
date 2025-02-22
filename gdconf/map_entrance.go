package gdconf

import (
	"encoding/base64"
	"fmt"
	"hkrpg/gameData"
	"os"

	"github.com/hjson/hjson-go/v4"
	"hkrpg/pkg/logger"
)

type MapEntrance struct {
	ID                    uint32   `json:"ID"`
	EntranceType          string   `json:"EntranceType"`
	PlaneID               uint32   `json:"PlaneID"`
	FloorID               uint32   `json:"FloorID"`
	StartGroupID          uint32   `json:"StartGroupID"`
	StartAnchorID         uint32   `json:"StartAnchorID"`
	BeginMainMissionList  []uint32 `json:"BeginMainMissionList"`
	FinishMainMissionList []uint32 `json:"FinishMainMissionList"`
	FinishSubMissionList  []uint32 `json:"FinishSubMissionList"`
}

func (g *GameDataConfig) loadMapEntrance() {
	g.MapEntranceMap = make(map[string]*MapEntrance)
	//playerElementsFilePath := g.excelPrefix + "MapEntrance.json"
	//playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFile, err := base64.StdEncoding.DecodeString(gameData.MapEntrance)
	if err != nil {
		logger.Error("get MapEntrance error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.MapEntranceMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v MapEntrance", len(g.MapEntranceMap))
}

func GetMapEntranceById(entryId string) *MapEntrance {
	return CONF.MapEntranceMap[entryId]
}

func GetMapEntranceMap() map[string]*MapEntrance {
	return CONF.MapEntranceMap
}

func GetEntryIdList() []uint32 {
	var entryIdList []uint32
	for _, id := range CONF.MapEntranceMap {
		entryIdList = append(entryIdList, id.ID)
	}
	return entryIdList
}
