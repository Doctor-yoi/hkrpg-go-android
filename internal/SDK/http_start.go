package SDK

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"hkrpg/internal/DataBase"
	"hkrpg/pkg/config"
	"hkrpg/pkg/logger"
)

type Server struct {
	Config     *config.Config
	Store      *DataBase.Store
	Router     *gin.Engine
	server     *http.Server
	AutoCreate sync.Mutex
}

func (s *Server) Start() error {
	// 初始化路由
	s.InitRouter()
	httpsAddr := s.Config.Http.Addr + ":" + strconv.FormatInt(s.Config.Http.Port, 10)
	err := s.startServer(httpsAddr)
	return err
}

func (s *Server) startServer(addr string) error {
	var err error
	server := &http.Server{Addr: addr, Handler: s.Router}

  err = server.ListenAndServe()
		
	 if err != nil && err != http.ErrServerClosed {
		logger.Error("hkrpg-go SDK 服务器启动失败, 原因: %s", err)
		return err
	}
	logger.Info("hkrpg-go SDK Http 在 %s 启动", addr)
	return nil
}

func (s *Server) Shutdown(context.Context) error {
	if s.server == nil {
		return nil
	}
	return s.server.Close()
}

func clientIPMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip, _, err := net.SplitHostPort(c.Request.RemoteAddr)
		if err != nil {
			c.Next()
			return
		}

		// 将 IP 信息存储在 gin.Context 中
		c.Set("IP", ip)
		c.Next()
	}
}
