package http_server

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"os"
)

func NewGinRouter(env string, handler *Handler, log *slog.Logger) *gin.Engine {
	r := gin.New()
	err := r.SetTrustedProxies(nil)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	switch env {
	case "prod":
	case "local":
	default:
		log.Error("unknown env", "env", env)
		os.Exit(1)
	}

	//	gin.SetMode(gin.ReleaseMode)

	r.POST("/link", handler.CreateAliasURL)
	r.GET("/link", handler.FetchSourceURL)

	return r
}
