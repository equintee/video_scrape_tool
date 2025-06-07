package content_parser

import (
	"net/http"
	"video_scrape_tool/cmd/api/parsers/twitter"

	"github.com/labstack/echo/v4"
)

type ContentHandler interface {
	Scrape(c echo.Context) error
}

type Instance struct {
	twitterService twitter.TwitterService
}

var instance *Instance

func NewContentHandler() *Instance {
	if instance == nil {
		instance = &Instance{
			twitterService: twitter.NewService(),
		}
	}
	return instance
}

func (t Instance) Scrape(c echo.Context) error {
	tweetUrl := c.QueryParams().Get("url")
	if tweetUrl == "" {
		return c.JSON(http.StatusBadRequest, "url query parameter is required")
	}
	filePath, err := t.twitterService.Scrape(tweetUrl)
	if err != nil {
		c.Error(err)
	} else {
		c.JSON(http.StatusOK, filePath)
	}

	return nil
}
