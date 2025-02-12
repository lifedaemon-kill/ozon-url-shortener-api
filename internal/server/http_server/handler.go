package http_server

import (
	"github.com/gin-gonic/gin"
	"github.com/lifedaemon-kill/ozon-url-shortener-api/internal/service"
	"log/slog"
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

type RequestData struct {
	URL string `json:"url"`
}

func (h *Handler) CreateAliasURL(c *gin.Context) {
	var req RequestData

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(400, gin.H{"bad request": "you should send xml or json 'url' filed with data"})
		return
	}

	alias, err := h.urlService.CreateAlias(req.URL)
	if err == nil {
		c.JSON(200, gin.H{"url": alias})
		return
	}
	if err.Error() == "bad request" {
		c.JSON(400, gin.H{"bad request": "url should be 10 symbols, either UPPER- or lower-case eng symbols, digits or _ (lower underlining)"})
		return
	} else {
		c.JSON(500, gin.H{"error": "internal error"})
		return
	}
}

func (h *Handler) FetchSourceURL(c *gin.Context) {
	var req RequestData

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(400, gin.H{"bad request": "you should send xml or json 'url' filed with data"})
		return
	}

	source, err := h.urlService.FetchSource(req.URL)
	if err == nil {
		c.JSON(200, gin.H{"url": source})
		return
	} else {
		c.JSON(500, gin.H{"error": "internal error"})
		return
	}
}
