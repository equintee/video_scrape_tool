package content_parser

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ContentHandler interface {
	Scrape(c echo.Context) error
}

type Handler struct {
	service *Service
}

var instance *Handler

func NewContentHandler() *Handler {
	if instance == nil {
		instance = &Handler{
			service: NewService(),
		}
	}
	return instance
}

func (t Handler) Scrape(c echo.Context) error {
	var request ScrapeRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	t.service.Scrape(c)
	return nil
}
