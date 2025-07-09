package content_parser

import (
	"fmt"
	"os"
	"video_scrape_tool/cmd/parsers/twitter"

	"github.com/labstack/echo/v4"
)

type ContentDto struct {
	File    *os.File
	Request ScrapeRequest
}
type ScrapeRequest struct {
	Url  string `json:"url"`
	Type string `json:"type"`
}

type ContentService interface {
	Scrape(c echo.Context) error
}

type Service struct {
	twitterService twitter.Service
}

func NewService() *Service {
	return &Service{
		twitterService: twitter.NewService(),
	}
}

func (s *Service) Scrape(c echo.Context) error {
	var request ScrapeRequest
	if err := c.Bind(&request); err != nil {
		return nil
	}
	var content ContentDto

	switch request.Type {
	case "twitter":
		content.File, _ = s.twitterService.Scrape(request.Url)
	default:
		return nil
	}

	fmt.Println(content)
	return nil
}
