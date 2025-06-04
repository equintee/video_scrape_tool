package twitter

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type TwitterHandler interface {
	Scrape(c echo.Context) error
}
type TwitterHandlerImpl struct {
	TwitterService TwitterService
}

var twitterHandler TwitterHandler

func NewTwitterHandler() TwitterHandler {
	if twitterHandler == nil {
		twitterHandler = TwitterHandlerImpl{
			TwitterService: NewService(),
		}
	}
	return twitterHandler
}

func (t TwitterHandlerImpl) Scrape(c echo.Context) error {
	tweetUrl := c.QueryParams().Get("url")
	if tweetUrl == "" {
		return c.JSON(http.StatusBadRequest, "url query parameter is required")
	}
	filePath, err := t.TwitterService.Scrape(tweetUrl)

	if err != nil {
		c.Error(err)
	} else {
		c.JSON(http.StatusOK, filePath)
	}

	return nil
}
