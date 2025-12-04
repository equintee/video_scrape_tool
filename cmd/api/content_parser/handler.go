package content_parser

import (
	"github.com/labstack/echo/v4"
)

type ContentHandler interface {
	Scrape(c echo.Context) error
	UpdateContent(c echo.Context) error
	GetContent(c echo.Context) error
	GetContentChunk(c echo.Context) error
	GetTags(c echo.Context) error
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

func (t Handler) UpdateContent(c echo.Context) error {
	t.service.UpdateContent(c)
	return nil
}

func (t Handler) GetContent(c echo.Context) error {
	t.service.GetContentMetaData(c)
	return nil
}

func (t Handler) GetContentChunk(c echo.Context) error {
	t.service.GetContentChunk(c)
	return nil
}

func (t Handler) GetTags(c echo.Context) error {
	t.service.GetTags(c)
	return nil
}
