package twitter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"

	"github.com/labstack/echo/v4"
)

type TwitterService interface {
	Scrape(c echo.Context) error
}

type twitterServiceImpl struct{}

func NewService() TwitterService {
	return &twitterServiceImpl{}
}

func (t *twitterServiceImpl) Scrape(c echo.Context) error {
	tweetUrl := c.QueryParams().Get("url")
	if tweetUrl == "" {
		return c.JSON(http.StatusBadRequest, "url query parameter is required")
	}

	regex := regexp.MustCompile(`https://x.com/(?P<Username>\w*)/status/(?P<TweetId>\d*)`)

	matches := regex.FindStringSubmatch(tweetUrl)
	tweetId := matches[2]

	headers := getGuestToken(tweetUrl)
	scrapeM3u8Url(tweetId, headers)
	return nil
}

func getGuestToken(tweetUrl string) string {
	url, _ := url.Parse(tweetUrl)
	requestParams := http.Request{
		Method: "GET",
		URL:    url,
		Header: map[string][]string{
			"user-agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36"},
			"accept":     {"*/*"},
		},
	}

	response, err := http.DefaultClient.Do(&requestParams)
	if err != nil {
		panic(err)
	}

	body, _ := io.ReadAll(response.Body)
	regex, _ := regexp.Compile("document.cookie=\"gt=(?P<Token>\\d+);")

	return regex.FindAllStringSubmatch(string(body), -1)[0][1]
}

func scrapeM3u8Url(tweetId string, token string) VideoVariant {
	baseUrl := fmt.Sprintf("https://api.x.com/graphql/_y7SZqeOFfgEivILXIy3tQ/TweetResultByRestId?variables=%7B%22tweetId%22%3A%22%s%22%2C%22withCommunity%22%3Afalse%2C%22includePromotedContent%22%3Afalse%2C%22withVoice%22%3Afalse%7D&features=%7B%22creator_subscriptions_tweet_preview_api_enabled%22%3Atrue%2C%22premium_content_api_read_enabled%22%3Afalse%2C%22communities_web_enable_tweet_community_results_fetch%22%3Atrue%2C%22c9s_tweet_anatomy_moderator_badge_enabled%22%3Atrue%2C%22responsive_web_grok_analyze_button_fetch_trends_enabled%22%3Afalse%2C%22responsive_web_grok_analyze_post_followups_enabled%22%3Afalse%2C%22responsive_web_jetfuel_frame%22%3Afalse%2C%22responsive_web_grok_share_attachment_enabled%22%3Atrue%2C%22articles_preview_enabled%22%3Atrue%2C%22responsive_web_edit_tweet_api_enabled%22%3Atrue%2C%22graphql_is_translatable_rweb_tweet_is_translatable_enabled%22%3Atrue%2C%22view_counts_everywhere_api_enabled%22%3Atrue%2C%22longform_notetweets_consumption_enabled%22%3Atrue%2C%22responsive_web_twitter_article_tweet_consumption_enabled%22%3Atrue%2C%22tweet_awards_web_tipping_enabled%22%3Afalse%2C%22responsive_web_grok_analysis_button_from_backend%22%3Afalse%2C%22creator_subscriptions_quote_tweet_preview_enabled%22%3Afalse%2C%22freedom_of_speech_not_reach_fetch_enabled%22%3Atrue%2C%22standardized_nudges_misinfo%22%3Atrue%2C%22tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled%22%3Atrue%2C%22rweb_video_timestamps_enabled%22%3Atrue%2C%22longform_notetweets_rich_text_read_enabled%22%3Atrue%2C%22longform_notetweets_inline_media_enabled%22%3Atrue%2C%22profile_label_improvements_pcf_label_in_post_enabled%22%3Atrue%2C%22rweb_tipjar_consumption_enabled%22%3Atrue%2C%22responsive_web_graphql_exclude_directive_enabled%22%3Atrue%2C%22verified_phone_label_enabled%22%3Afalse%2C%22responsive_web_grok_image_annotation_enabled%22%3Afalse%2C%22responsive_web_graphql_skip_user_profile_image_extensions_enabled%22%3Afalse%2C%22responsive_web_graphql_timeline_navigation_enabled%22%3Atrue%2C%22responsive_web_enhance_cards_enabled%22%3Afalse%7D&fieldToggles=%7B%22withArticleRichContentState%22%3Atrue%2C%22withArticlePlainText%22%3Afalse%2C%22withGrokAnalyze%22%3Afalse%2C%22withDisallowedReplyControls%22%3Afalse%7D", tweetId)
	requestUrl, _ := url.Parse(baseUrl)
	requestParams := http.Request{
		Method: "GET",
		URL:    requestUrl,
		Header: map[string][]string{
			"user-agent":    {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36"},
			"Authorization": {"Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"},
			"x-guest-token": {token},
		},
	}

	response, _ := http.DefaultClient.Do(&requestParams)
	//Parsing fails
	parsedRespose, _ := io.ReadAll(response.Body)

	var jsonResponse WebApiResponse
	json.Unmarshal(parsedRespose, &jsonResponse)

	videoDetails := jsonResponse.Data.TweetResult.Result.Legacy.Entities.Media[0].VideoInfo.Variants

	var highestBitrateVideo VideoVariant

	highestBitrate := 0

	for _, videoDetail := range videoDetails {
		if videoDetail.Bitrate > highestBitrate {
			highestBitrate = videoDetail.Bitrate
			highestBitrateVideo = videoDetail
		}
	}

	return highestBitrateVideo
}
