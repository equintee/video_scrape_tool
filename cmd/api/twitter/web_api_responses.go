package twitter

type WebApiResponse struct {
	Data struct {
		TweetResult struct {
			Result struct {
				Typename string `json:"__typename"`
				RestID   string `json:"rest_id"`
				Core     struct {
					UserResults struct {
						Result struct {
							Typename                   string `json:"__typename"`
							ID                         string `json:"id"`
							RestID                     string `json:"rest_id"`
							AffiliatesHighlightedLabel struct {
							} `json:"affiliates_highlighted_label"`
							ParodyCommentaryFanLabel string `json:"parody_commentary_fan_label"`
							IsBlueVerified           bool   `json:"is_blue_verified"`
							ProfileImageShape        string `json:"profile_image_shape"`
							Legacy                   struct {
								CreatedAt           string `json:"created_at"`
								DefaultProfile      bool   `json:"default_profile"`
								DefaultProfileImage bool   `json:"default_profile_image"`
								Description         string `json:"description"`
								Entities            struct {
									Description struct {
										Urls []any `json:"urls"`
									} `json:"description"`
								} `json:"entities"`
								FastFollowersCount      int    `json:"fast_followers_count"`
								FavouritesCount         int    `json:"favourites_count"`
								FollowersCount          int    `json:"followers_count"`
								FriendsCount            int    `json:"friends_count"`
								HasCustomTimelines      bool   `json:"has_custom_timelines"`
								IsTranslator            bool   `json:"is_translator"`
								ListedCount             int    `json:"listed_count"`
								Location                string `json:"location"`
								MediaCount              int    `json:"media_count"`
								Name                    string `json:"name"`
								NormalFollowersCount    int    `json:"normal_followers_count"`
								PinnedTweetIdsStr       []any  `json:"pinned_tweet_ids_str"`
								PossiblySensitive       bool   `json:"possibly_sensitive"`
								ProfileBannerURL        string `json:"profile_banner_url"`
								ProfileImageURLHTTPS    string `json:"profile_image_url_https"`
								ProfileInterstitialType string `json:"profile_interstitial_type"`
								ScreenName              string `json:"screen_name"`
								StatusesCount           int    `json:"statuses_count"`
								TranslatorType          string `json:"translator_type"`
								Verified                bool   `json:"verified"`
								WithheldInCountries     []any  `json:"withheld_in_countries"`
							} `json:"legacy"`
							TipjarSettings struct {
								IsEnabled      bool   `json:"is_enabled"`
								BandcampHandle string `json:"bandcamp_handle"`
								BitcoinHandle  string `json:"bitcoin_handle"`
								CashAppHandle  string `json:"cash_app_handle"`
								EthereumHandle string `json:"ethereum_handle"`
								GofundmeHandle string `json:"gofundme_handle"`
								PatreonHandle  string `json:"patreon_handle"`
								PayPalHandle   string `json:"pay_pal_handle"`
								VenmoHandle    string `json:"venmo_handle"`
							} `json:"tipjar_settings"`
						} `json:"result"`
					} `json:"user_results"`
				} `json:"core"`
				UnmentionData struct {
				} `json:"unmention_data"`
				EditControl struct {
					EditTweetIds       []string `json:"edit_tweet_ids"`
					EditableUntilMsecs string   `json:"editable_until_msecs"`
					IsEditEligible     bool     `json:"is_edit_eligible"`
					EditsRemaining     string   `json:"edits_remaining"`
				} `json:"edit_control"`
				IsTranslatable bool `json:"is_translatable"`
				Views          struct {
					Count string `json:"count"`
					State string `json:"state"`
				} `json:"views"`
				Source string `json:"source"`
				Legacy struct {
					BookmarkCount     int    `json:"bookmark_count"`
					Bookmarked        bool   `json:"bookmarked"`
					CreatedAt         string `json:"created_at"`
					ConversationIDStr string `json:"conversation_id_str"`
					DisplayTextRange  []int  `json:"display_text_range"`
					Entities          struct {
						Hashtags []any `json:"hashtags"`
						Media    []struct {
							DisplayURL          string `json:"display_url"`
							ExpandedURL         string `json:"expanded_url"`
							IDStr               string `json:"id_str"`
							Indices             []int  `json:"indices"`
							MediaKey            string `json:"media_key"`
							MediaURLHTTPS       string `json:"media_url_https"`
							SourceStatusIDStr   string `json:"source_status_id_str"`
							SourceUserIDStr     string `json:"source_user_id_str"`
							Type                string `json:"type"`
							URL                 string `json:"url"`
							AdditionalMediaInfo struct {
								Monetizable bool `json:"monetizable"`
								SourceUser  struct {
									UserResults struct {
										Result struct {
											Typename                   string `json:"__typename"`
											ID                         string `json:"id"`
											RestID                     string `json:"rest_id"`
											AffiliatesHighlightedLabel struct {
											} `json:"affiliates_highlighted_label"`
											ParodyCommentaryFanLabel string `json:"parody_commentary_fan_label"`
											IsBlueVerified           bool   `json:"is_blue_verified"`
											ProfileImageShape        string `json:"profile_image_shape"`
											Legacy                   struct {
												CreatedAt           string `json:"created_at"`
												DefaultProfile      bool   `json:"default_profile"`
												DefaultProfileImage bool   `json:"default_profile_image"`
												Description         string `json:"description"`
												Entities            struct {
													Description struct {
														Urls []any `json:"urls"`
													} `json:"description"`
												} `json:"entities"`
												FastFollowersCount      int    `json:"fast_followers_count"`
												FavouritesCount         int    `json:"favourites_count"`
												FollowersCount          int    `json:"followers_count"`
												FriendsCount            int    `json:"friends_count"`
												HasCustomTimelines      bool   `json:"has_custom_timelines"`
												IsTranslator            bool   `json:"is_translator"`
												ListedCount             int    `json:"listed_count"`
												Location                string `json:"location"`
												MediaCount              int    `json:"media_count"`
												Name                    string `json:"name"`
												NormalFollowersCount    int    `json:"normal_followers_count"`
												PinnedTweetIdsStr       []any  `json:"pinned_tweet_ids_str"`
												PossiblySensitive       bool   `json:"possibly_sensitive"`
												ProfileBannerURL        string `json:"profile_banner_url"`
												ProfileImageURLHTTPS    string `json:"profile_image_url_https"`
												ProfileInterstitialType string `json:"profile_interstitial_type"`
												ScreenName              string `json:"screen_name"`
												StatusesCount           int    `json:"statuses_count"`
												TranslatorType          string `json:"translator_type"`
												Verified                bool   `json:"verified"`
												WithheldInCountries     []any  `json:"withheld_in_countries"`
											} `json:"legacy"`
											TipjarSettings struct {
												IsEnabled      bool   `json:"is_enabled"`
												BandcampHandle string `json:"bandcamp_handle"`
												BitcoinHandle  string `json:"bitcoin_handle"`
												CashAppHandle  string `json:"cash_app_handle"`
												EthereumHandle string `json:"ethereum_handle"`
												GofundmeHandle string `json:"gofundme_handle"`
												PatreonHandle  string `json:"patreon_handle"`
												PayPalHandle   string `json:"pay_pal_handle"`
												VenmoHandle    string `json:"venmo_handle"`
											} `json:"tipjar_settings"`
										} `json:"result"`
									} `json:"user_results"`
								} `json:"source_user"`
							} `json:"additional_media_info"`
							ExtMediaAvailability struct {
								Status string `json:"status"`
							} `json:"ext_media_availability"`
							Sizes struct {
								Large struct {
									H      int    `json:"h"`
									W      int    `json:"w"`
									Resize string `json:"resize"`
								} `json:"large"`
								Medium struct {
									H      int    `json:"h"`
									W      int    `json:"w"`
									Resize string `json:"resize"`
								} `json:"medium"`
								Small struct {
									H      int    `json:"h"`
									W      int    `json:"w"`
									Resize string `json:"resize"`
								} `json:"small"`
								Thumb struct {
									H      int    `json:"h"`
									W      int    `json:"w"`
									Resize string `json:"resize"`
								} `json:"thumb"`
							} `json:"sizes"`
							OriginalInfo struct {
								Height     int   `json:"height"`
								Width      int   `json:"width"`
								FocusRects []any `json:"focus_rects"`
							} `json:"original_info"`
							VideoInfo struct {
								AspectRatio    []int          `json:"aspect_ratio"`
								DurationMillis int            `json:"duration_millis"`
								Variants       []VideoVariant `json:"variants"`
							} `json:"video_info"`
							MediaResults struct {
								Result struct {
									MediaKey string `json:"media_key"`
								} `json:"result"`
							} `json:"media_results"`
						} `json:"media"`
						Symbols      []any `json:"symbols"`
						Timestamps   []any `json:"timestamps"`
						Urls         []any `json:"urls"`
						UserMentions []any `json:"user_mentions"`
					} `json:"entities"`
					ExtendedEntities struct {
						Media []struct {
							DisplayURL          string `json:"display_url"`
							ExpandedURL         string `json:"expanded_url"`
							IDStr               string `json:"id_str"`
							Indices             []int  `json:"indices"`
							MediaKey            string `json:"media_key"`
							MediaURLHTTPS       string `json:"media_url_https"`
							SourceStatusIDStr   string `json:"source_status_id_str"`
							SourceUserIDStr     string `json:"source_user_id_str"`
							Type                string `json:"type"`
							URL                 string `json:"url"`
							AdditionalMediaInfo struct {
								Monetizable bool `json:"monetizable"`
								SourceUser  struct {
									UserResults struct {
										Result struct {
											Typename                   string `json:"__typename"`
											ID                         string `json:"id"`
											RestID                     string `json:"rest_id"`
											AffiliatesHighlightedLabel struct {
											} `json:"affiliates_highlighted_label"`
											ParodyCommentaryFanLabel string `json:"parody_commentary_fan_label"`
											IsBlueVerified           bool   `json:"is_blue_verified"`
											ProfileImageShape        string `json:"profile_image_shape"`
											Legacy                   struct {
												CreatedAt           string `json:"created_at"`
												DefaultProfile      bool   `json:"default_profile"`
												DefaultProfileImage bool   `json:"default_profile_image"`
												Description         string `json:"description"`
												Entities            struct {
													Description struct {
														Urls []any `json:"urls"`
													} `json:"description"`
												} `json:"entities"`
												FastFollowersCount      int    `json:"fast_followers_count"`
												FavouritesCount         int    `json:"favourites_count"`
												FollowersCount          int    `json:"followers_count"`
												FriendsCount            int    `json:"friends_count"`
												HasCustomTimelines      bool   `json:"has_custom_timelines"`
												IsTranslator            bool   `json:"is_translator"`
												ListedCount             int    `json:"listed_count"`
												Location                string `json:"location"`
												MediaCount              int    `json:"media_count"`
												Name                    string `json:"name"`
												NormalFollowersCount    int    `json:"normal_followers_count"`
												PinnedTweetIdsStr       []any  `json:"pinned_tweet_ids_str"`
												PossiblySensitive       bool   `json:"possibly_sensitive"`
												ProfileBannerURL        string `json:"profile_banner_url"`
												ProfileImageURLHTTPS    string `json:"profile_image_url_https"`
												ProfileInterstitialType string `json:"profile_interstitial_type"`
												ScreenName              string `json:"screen_name"`
												StatusesCount           int    `json:"statuses_count"`
												TranslatorType          string `json:"translator_type"`
												Verified                bool   `json:"verified"`
												WithheldInCountries     []any  `json:"withheld_in_countries"`
											} `json:"legacy"`
											TipjarSettings struct {
												IsEnabled      bool   `json:"is_enabled"`
												BandcampHandle string `json:"bandcamp_handle"`
												BitcoinHandle  string `json:"bitcoin_handle"`
												CashAppHandle  string `json:"cash_app_handle"`
												EthereumHandle string `json:"ethereum_handle"`
												GofundmeHandle string `json:"gofundme_handle"`
												PatreonHandle  string `json:"patreon_handle"`
												PayPalHandle   string `json:"pay_pal_handle"`
												VenmoHandle    string `json:"venmo_handle"`
											} `json:"tipjar_settings"`
										} `json:"result"`
									} `json:"user_results"`
								} `json:"source_user"`
							} `json:"additional_media_info"`
							ExtMediaAvailability struct {
								Status string `json:"status"`
							} `json:"ext_media_availability"`
							Sizes struct {
								Large struct {
									H      int    `json:"h"`
									W      int    `json:"w"`
									Resize string `json:"resize"`
								} `json:"large"`
								Medium struct {
									H      int    `json:"h"`
									W      int    `json:"w"`
									Resize string `json:"resize"`
								} `json:"medium"`
								Small struct {
									H      int    `json:"h"`
									W      int    `json:"w"`
									Resize string `json:"resize"`
								} `json:"small"`
								Thumb struct {
									H      int    `json:"h"`
									W      int    `json:"w"`
									Resize string `json:"resize"`
								} `json:"thumb"`
							} `json:"sizes"`
							OriginalInfo struct {
								Height     int   `json:"height"`
								Width      int   `json:"width"`
								FocusRects []any `json:"focus_rects"`
							} `json:"original_info"`
							VideoInfo struct {
								AspectRatio    []int          `json:"aspect_ratio"`
								DurationMillis int            `json:"duration_millis"`
								Variants       []VideoVariant `json:"variants"`
							} `json:"video_info"`
							MediaResults struct {
								Result struct {
									MediaKey string `json:"media_key"`
								} `json:"result"`
							} `json:"media_results"`
						} `json:"media"`
					} `json:"extended_entities"`
					FavoriteCount             int    `json:"favorite_count"`
					Favorited                 bool   `json:"favorited"`
					FullText                  string `json:"full_text"`
					IsQuoteStatus             bool   `json:"is_quote_status"`
					Lang                      string `json:"lang"`
					PossiblySensitive         bool   `json:"possibly_sensitive"`
					PossiblySensitiveEditable bool   `json:"possibly_sensitive_editable"`
					QuoteCount                int    `json:"quote_count"`
					ReplyCount                int    `json:"reply_count"`
					RetweetCount              int    `json:"retweet_count"`
					Retweeted                 bool   `json:"retweeted"`
					UserIDStr                 string `json:"user_id_str"`
					IDStr                     string `json:"id_str"`
				} `json:"legacy"`
			} `json:"result"`
		} `json:"tweetResult"`
	} `json:"data"`
}

type VideoVariant struct {
	Bitrate     int    `json:"bitrate"`
	ContentType string `json:"content_type"`
	URL         string `json:"url"`
}
