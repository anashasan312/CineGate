package server

import (
	"fmt"
	"net"
	"net/http"

	log "github.com/anashasan312/movie-reservation-system/logger"
	config "github.com/anashasan312/movie-reservation-system/pkg/infrastructure/config"
	"github.com/gin-gonic/gin"
)

var (
	BasePath = "/ayman"
)

type HttpServer struct {
	PublicEngine  *gin.Engine
	PrivateEngine *gin.Engine
	Config        *config.AppConfig
}

func (h *HttpServer) SetupRoutes() {
	SetupPublicRoutes(h)
	// SetupPrivateRoutes(h)
}

type engine struct {
	serverType string
	server     *http.Server
}

func NewServer(cfg *config.AppConfig) (*HttpServer, error) {
	log.Boot(fmt.Sprintf("%v: Creating new server", ""))
	gin.SetMode(gin.DebugMode)

	publicEngine := gin.New()
	privateEngine := gin.New()

	publicEngine.Use(
		gin.Recovery(),
	)

	privateEngine.Use(
		gin.Recovery(),
	)

	return &HttpServer{
		PublicEngine:  publicEngine,
		PrivateEngine: privateEngine,
		Config:        cfg,
	}, nil
}

func (h *HttpServer) Run() {
	engines := []engine{
		{
			server: &http.Server{
				Addr:    net.JoinHostPort(h.Config.Server.Host, fmt.Sprint(h.Config.Server.PublicPort)),
				Handler: h.PublicEngine,
			},
		},
	}

	for _, e := range engines {
		if e.server == nil {
			continue
		}
	}
}
