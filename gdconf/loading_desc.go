package gdconf

import (
	"encoding/base64"
	"fmt"
	"hkrpg/gameData"
	"math/rand"
	"os"

	"github.com/hjson/hjson-go/v4"
	"hkrpg/pkg/logger"
)

type LoadingDesc struct {
	ID       uint32 `json:"ID"`       // 随机种子
	MinLevel uint32 `json:"MinLevel"` // 最小等级
	MaxLevel uint32 `json:"MaxLevel"` // 最大等级
	Weight   uint32 `json:"Weight"`   // 比重(貌似全部都是20)
	// TODO 还有部分字段没读，需要再读
}

func (g *GameDataConfig) loadLoadingDesc() {
	g.LoadingDescMap = make(map[string]*LoadingDesc)
	//playerElementsFilePath := g.excelPrefix + "LoadingDesc.json"
	//playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	//if err != nil {
	//	info := fmt.Sprintf("open file error: %v", err)
	//	panic(info)
	//}
	playerElementsFile, err := base64.StdEncoding.DecodeString(gameData.LoadingDesc)
	if err != nil {
		logger.Error("get LoadingDesc error")
		os.Exit(-1)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.LoadingDescMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}

	logger.Info("load %v LoadingDesc", len(g.LoadingDescMap))
}

func GetLoadingDesc() uint32 {
	var list []uint32
	idIndex := rand.Intn(len(CONF.LoadingDescMap))
	for _, id := range CONF.LoadingDescMap {
		list = append(list, id.ID)
	}
	return list[idIndex]
}
