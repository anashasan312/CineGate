package server

import (
	"fmt"

	log "github.com/anashasan312/movie-reservation-system/logger"
	config "github.com/anashasan312/movie-reservation-system/pkg/infrastructure/config"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	PublicEngine  *gin.Engine
	PrivateEngine *gin.Engine
}

func NewServer(cfg *config.Appconfig) (*HttpServer, error) {
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

	return &HttpServer{}, nil
}
