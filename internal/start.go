package internal

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"hkrpg/gameData"
	"hkrpg/internal/DataBase"
	"hkrpg/internal/SDK"
	"hkrpg/pkg/config"
	"hkrpg/pkg/logger"
	"hkrpg/pkg/random"
)

// 初始化所有服务
func NewServer(cfg *config.Config) *SDK.Server {
	s := &SDK.Server{}
	s.Config = cfg
	if cfg.UseDatabase {
		s.Store = DataBase.NewStore(s.Config) // 初始化数据库连接
	} else {
		s.Store = nil
	}
	gin.SetMode(gin.ReleaseMode) // 初始化gin
	s.Router = gin.New()         // gin.Default()
	s.Router.Use(gin.Recovery())
	cfg.Ec2b = getEc2b() // 读取ec2b密钥

	return s
}

func getEc2b() *random.Ec2b {
	ec2b_b64 := gameData.Ec2b
	if ec2b_b64 == "" {
		ec2p := random.NewEc2b().Bytes()
		gameData.Ec2b = base64.StdEncoding.EncodeToString(ec2p)
	}
	ec2p, err := base64.StdEncoding.DecodeString(gameData.Ec2b)
	if err != nil {
		logger.Error("read Ec2b error")
		return nil
	}
	ec2b, err := random.LoadEc2bKey(ec2p)
	if err != nil {
		logger.Error("parse region ec2b error: %v", err)
		return nil
	}
	return ec2b
	//open, err := os.Open("/storage/emulated/0/Documents/hkrpg-go/GameData/Ec2b.bin")
	//defer open.Close()
	//if err != nil {
	//	ec2p := random.NewEc2b().Bytes()
	//	ioutil.WriteFile("/storage/emulated/0/Documents/hkrpg-go/GameData/Ec2b.bin", ec2p, 0644)
	//	logger.Info("ec2b不存在,生成ec2b文件中")
	//	ec2b, err := random.LoadEc2bKey(ec2p)
	//	if err != nil {
	//		logger.Error("parse region ec2b error: %v", err)
	//		return nil
	//	}
	//	return ec2b
	//} else {
	//	ec2p, err := io.ReadAll(open)
	//	if err != nil {
	//		logger.Error("read Ec2b error")
	//		return nil
	//	}
	//	defer open.Close()
	//	ec2b, err := random.LoadEc2bKey(ec2p)
	//	if err != nil {
	//		logger.Error("parse region ec2b error: %v", err)
	//		return nil
	//	}
	//	return ec2b
	//}
}
