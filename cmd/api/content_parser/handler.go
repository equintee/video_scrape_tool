package content_parser

import (
	"github.com/labstack/echo/v4"
)

type ContentHandler interface {
	Scrape(c echo.Context) error
	GetContent(c echo.Context) error
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
	t.service.Scrape(c)
	return nil
}

func (t Handler) GetContent(c echo.Context) error {
	t.service.GetContent(c)
	return nil
}
