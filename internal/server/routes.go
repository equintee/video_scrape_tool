package server

import (
	"net/http"
	"video_scrape_tool/cmd/api/content_parser"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var handlers handlerStruct

type handlerStruct struct {
	contentHandler content_parser.ContentHandler
}

func init() {
	handlers = handlerStruct{
		contentHandler: content_parser.NewContentHandler(),
	}
}
func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	e.GET("/", handlers.contentHandler.Scrape)

	return e
}
