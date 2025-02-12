package http_server

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	ierrors "github.com/lifedaemon-kill/ozon-url-shortener-api/internal/pkg/internal_errors"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/service"
	"log/slog"
	"time"
)

type Handler struct {
	log        *slog.Logger
	urlService service.UrlShortener
}

func NewHandler(log *slog.Logger, urlService service.UrlShortener) *Handler {
	return &Handler{
		log:        log,
		urlService: urlService,
	}
}

func (h *Handler) CreateAliasURL(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c, 15*time.Second)
	defer cancel()

	req := c.GetHeader("X-URL")

	if req == "" {
		h.log.Error("http_server.CreateAliasURL", "err", "no X-URL header")
		c.JSON(400, gin.H{"bad request": "you should send xml or json 'url' filed with data"})
		return
	}
	h.log.Debug("start http_server.CreateAliasURL", "url", req)

	alias, err := h.urlService.CreateAlias(ctx, req)
	if err == nil {
		h.log.Debug("http_server.CreateAliasURL success", "url", req)
		c.JSON(200, gin.H{"url": alias})
		return
	}
	if errors.Is(err, ierrors.InvalidURL) {
		h.log.Error("http_server.CreateAliasURL failed Invalid URL", "url", req)
		c.JSON(400, gin.H{"bad request": "url not match pattern "})
		return
	}
	if errors.Is(err, ierrors.SourceAlreadyExist) {
		h.log.Warn("http_server.CreateAliasURL failed SourceAlreadyExist", "url", req)
		c.JSON(400, gin.H{"bad request": "source already exists"})
	} else {
		h.log.Error("http_server.CreateAliasURL failed", "err", err)
		c.JSON(500, gin.H{"error": "internal error"})
		return
	}
}

func (h *Handler) FetchSourceURL(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, 15*time.Second)
	defer cancel()
	req := c.GetHeader("X-URL")

	if req == "" {
		h.log.Error("http_server.FetchSourceURL", "err", "no X-URL header")
		c.JSON(400, gin.H{"bad request": "you should send xml or json 'url' filed with data"})
		return
	}
	h.log.Debug("http_server.FetchSourceURL after binding", "url", req)

	source, err := h.urlService.FetchSource(ctx, req)
	if err == nil {
		h.log.Debug("http_server.FetchSourceURL success", "url", req)
		c.JSON(200, gin.H{"url": source})
		return
	} else {
		h.log.Error("http_server.FetchSourceURL failed", "err", err)
		c.JSON(500, gin.H{"error": "internal error"})
		return
	}
}
