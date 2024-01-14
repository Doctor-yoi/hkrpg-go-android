package gdconf

import (
	"encoding/base64"
	"fmt"
	"hkrpg/gameData"
	"os"

	"github.com/hjson/hjson-go/v4"
	"hkrpg/pkg/logger"
)

type ActivityScheduling struct {
	ActivityId uint32 `json:"activityId"`
	BeginTime  int64  `json:"beginTime"`
	EndTime    int64  `json:"endTime"`
	ModuleId   uint32 `json:"moduleId"`
}

func (g *GameDataConfig) loadActivityScheduling() {
	g.ActivitySchedulingMap = make([]*ActivityScheduling, 0)
	//playerElementsFilePath := g.dataPrefix + "ActivityScheduling.json"
	//playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFile, err := base64.StdEncoding.DecodeString(gameData.ActivityScheduling)
	if err != nil {
		logger.Error("get ActivityScheduling error")
		os.Exit(-1)
	}
	err = hjson.Unmarshal(playerElementsFile, &g.ActivitySchedulingMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	logger.Info("load %v ActivityScheduling", len(g.ActivitySchedulingMap))
}

func GetActivitySchedulingMap() []*ActivityScheduling {
	return CONF.ActivitySchedulingMap
}
