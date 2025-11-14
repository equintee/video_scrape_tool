package twitter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"video_scrape_tool/cmd/util"
)

type Service interface {
	Scrape(tweetUrl string) (*os.File, error)
	getGuestToken(tweetUrl string) string
	scrapeM3u8Url(tweetId string, token string) VideoVariant
}

var instance Service

type twitterServiceImpl struct {
	BaseUrl string
}

func NewService() Service {
	if instance == nil {
		instance = &twitterServiceImpl{
			BaseUrl: generateBaseEndpoint(),
		}
	}
	return instance
}

func generateBaseEndpoint() string {
	documentationUrl, _ := url.Parse("https://raw.githubusercontent.com/fa0311/TwitterInternalAPIDocument/refs/heads/master/docs/json/GraphQL.json")
	request := http.Request{
		URL: documentationUrl,
	}
	response, err := http.DefaultClient.Do(&request)
	if err != nil {
		panic(err)
	}

	rawData, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var data []map[string]interface{}
	json.Unmarshal(rawData, &data)

	var endpointSuffix string
	for _, v := range data {
		exports := v["exports"].(map[string]interface{})
		if exports["operationName"] == "TweetResultByRestId" {
			fmt.Println(exports["queryId"].(string))
			endpointSuffix = exports["queryId"].(string)
		}
	}
	if endpointSuffix == "" {
		panic("Could not find EndpointSuffix")
	}

	baseUrl := fmt.Sprintf("https://api.x.com/graphql/%s/TweetResultByRestId?", endpointSuffix)
	return baseUrl
}

func (t *twitterServiceImpl) Scrape(tweetUrl string) (*os.File, error) {
	regex := regexp.MustCompile(`https://x.com/(?P<Username>\w*)/status/(?P<TweetId>\d*)`)

	matches := regex.FindStringSubmatch(tweetUrl)
	tweetId := matches[2]

	headers := t.getGuestToken(tweetUrl)
	videoUrl := t.scrapeM3u8Url(tweetId, headers).URL
	videoClip, err := util.SaveVideoFromUrl(videoUrl)

	if err != nil {
		panic(err)
	}

	return videoClip, nil
}

func (t twitterServiceImpl) getGuestToken(tweetUrl string) string {
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

func (t twitterServiceImpl) scrapeM3u8Url(tweetId string, token string) VideoVariant {
	params := url.Values{}
	params.Add("variables", fmt.Sprintf("{\"tweetId\":\"%s\",\"withCommunity\":false,\"includePromotedContent\":false,\"withVoice\":false}&features={\"creator_subscriptions_tweet_preview_api_enabled\":true,\"premium_content_api_read_enabled\":false,\"communities_web_enable_tweet_community_results_fetch\":true,\"c9s_tweet_anatomy_moderator_badge_enabled\":true,\"responsive_web_grok_analyze_button_fetch_trends_enabled\":false,\"responsive_web_grok_analyze_post_followups_enabled\":false,\"responsive_web_jetfuel_frame\":false,\"responsive_web_grok_share_attachment_enabled\":true,\"articles_preview_enabled\":true,\"responsive_web_edit_tweet_api_enabled\":true,\"graphql_is_translatable_rweb_tweet_is_translatable_enabled\":true,\"view_counts_everywhere_api_enabled\":true,\"longform_notetweets_consumption_enabled\":true,\"responsive_web_twitter_article_tweet_consumption_enabled\":true,\"tweet_awards_web_tipping_enabled\":false,\"responsive_web_grok_show_grok_translated_post\":false,\"responsive_web_grok_analysis_button_from_backend\":false,\"creator_subscriptions_quote_tweet_preview_enabled\":false,\"freedom_of_speech_not_reach_fetch_enabled\":true,\"standardized_nudges_misinfo\":true,\"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled\":true,\"longform_notetweets_rich_text_read_enabled\":true,\"longform_notetweets_inline_media_enabled\":true,\"profile_label_improvements_pcf_label_in_post_enabled\":true,\"rweb_tipjar_consumption_enabled\":true,\"verified_phone_label_enabled\":false,\"responsive_web_grok_image_annotation_enabled\":true,\"responsive_web_graphql_skip_user_profile_image_extensions_enabled\":false,\"responsive_web_graphql_timeline_navigation_enabled\":true,\"responsive_web_enhance_cards_enabled\":false}&fieldToggles={\"withArticleRichContentState\":true,\"withArticlePlainText\":false,\"withGrokAnalyze\":false,\"withDisallowedReplyControls\":false}", tweetId))
	params.Add("features", "{\"creator_subscriptions_tweet_preview_api_enabled\":true,\"premium_content_api_read_enabled\":false,\"communities_web_enable_tweet_community_results_fetch\":true,\"c9s_tweet_anatomy_moderator_badge_enabled\":true,\"responsive_web_grok_analyze_button_fetch_trends_enabled\":false,\"responsive_web_grok_analyze_post_followups_enabled\":true,\"responsive_web_jetfuel_frame\":true,\"responsive_web_grok_share_attachment_enabled\":true,\"articles_preview_enabled\":true,\"responsive_web_edit_tweet_api_enabled\":true,\"graphql_is_translatable_rweb_tweet_is_translatable_enabled\":true,\"view_counts_everywhere_api_enabled\":true,\"longform_notetweets_consumption_enabled\":true,\"responsive_web_twitter_article_tweet_consumption_enabled\":true,\"tweet_awards_web_tipping_enabled\":false,\"responsive_web_grok_show_grok_translated_post\":false,\"responsive_web_grok_analysis_button_from_backend\":true,\"creator_subscriptions_quote_tweet_preview_enabled\":false,\"freedom_of_speech_not_reach_fetch_enabled\":true,\"standardized_nudges_misinfo\":true,\"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled\":true,\"longform_notetweets_rich_text_read_enabled\":true,\"longform_notetweets_inline_media_enabled\":true,\"payments_enabled\":false,\"profile_label_improvements_pcf_label_in_post_enabled\":true,\"responsive_web_profile_redirect_enabled\":false,\"rweb_tipjar_consumption_enabled\":true,\"verified_phone_label_enabled\":false,\"responsive_web_grok_image_annotation_enabled\":true,\"responsive_web_grok_imagine_annotation_enabled\":true,\"responsive_web_grok_community_note_auto_translation_is_enabled\":false,\"responsive_web_graphql_skip_user_profile_image_extensions_enabled\":false,\"responsive_web_graphql_timeline_navigation_enabled\":true,\"responsive_web_enhance_cards_enabled\":false}")
	params.Add("fieldToggles", "{\"withArticleRichContentState\":true,\"withArticlePlainText\":false,\"withGrokAnalyze\":false,\"withDisallowedReplyControls\":false}")

	baseUrl := t.BaseUrl + params.Encode()
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
	parsedRespose, _ := io.ReadAll(response.Body)

	var jsonResponse WebApiResponse

	if response.StatusCode != 200 {
		pirt := make(map[string]interface{})
		json.Unmarshal(parsedRespose, &pirt)
		panic(pirt)
	}
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
