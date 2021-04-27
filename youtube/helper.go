package youtube

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"time"
)

type youtubeHelperOptions struct {
        verbose bool
	timeout time.Duration
}

func defaultYoutubeHelperOptions() *youtubeHelperOptions {
        return &youtubeHelperOptions{
                verbose: false,
		timeout: 20 * time.Second,
        }
}

type YoutubeHelperOption func(*youtubeHelperOptions)

func YoutubeHelperVerbose(verbose bool) YoutubeHelperOption {
        return func(opts *youtubeHelperOptions) {
                opts.verbose = verbose
        }
}

func YoutubeHelperTimeout(timeout time.Duration) YoutubeHelperOption {
        return func(opts *youtubeHelperOptions) {
                opts.timeout = timeout
        }
}

type youtubeHelperSearchOptions struct {
	limit            int
	parts            []string
	relatedToVideoId string
	channelId        string
	channelType      string
	eventType        string
	order            string
	publishedAfter   time.Time
	publishedBefore  time.Time
	q                string
	regionCode       string
	safeSearch       string
	topicId          string
	searchType       string
	videoCaption     string
	videoCategoryId  string
	videoDefinition  string
	videoDimension   string
	videoDuration    string
	videoEmbeddable  string
	videoLicense     string
	videoSyndicated  string
	videoType        string
}

func defaultYoutubeHelperSearchOptions() *youtubeHelperSearchOptions {
        return &youtubeHelperSearchOptions{
                limit:            0,
		parts:            []string{ "id" },
		relatedToVideoId: "",
		channelId:        "",
		channelType:      "any",
		eventType:        "",
		order:            "date",
		publishedAfter:   time.Time{},
		publishedBefore:  time.Time{},
		q:                "",
		regionCode:       "",
		safeSearch:       "none",
		topicId:          "",
		searchType:       "",
		videoCaption:     "any",
		videoCategoryId:  "",
		videoDefinition:  "any",
		videoDimension:   "any",
		videoDuration:    "any",
		videoEmbeddable:  "any",
		videoLicense:     "any",
		videoSyndicated:  "any",
		videoType:        "any",
        }
}

type YoutubeHelperSearchOption func(*youtubeHelperSearchOptions)

func YoutubeHelperSearchLimit(limit int) YoutubeHelperSearchOption {
        return func(opts *youtubeHelperSearchOptions) {
                opts.limit = limit
        }
}

func YoutubeHelperSearchParts(parts ...string) YoutubeHelperSearchOption {
        return func(opts *youtubeHelperSearchOptions) {
                opts.parts = parts
        }
}

func YoutubeHelperSearchRelatedToVideoId(relatedToVideoId string) YoutubeHelperSearchOption {
        return func(opts *youtubeHelperSearchOptions) {
                opts.relatedToVideoId = relatedToVideoId
        }
}

func YoutubeHelperSearchChannelId(channelId string) YoutubeHelperSearchOption {
        return func(opts *youtubeHelperSearchOptions) {
                opts.channelId = channelId
        }
}

func YoutubeHelperSearchChannelType(channelType string) YoutubeHelperSearchOption {
        return func(opts *youtubeHelperSearchOptions) {
                opts.channelType = channelType
        }
}

func YoutubeHelperSearchEventType(eventType string) YoutubeHelperSearchOption {
        return func(opts *youtubeHelperSearchOptions) {
                opts.eventType = eventType
        }
}

func YoutubeHelperSearchOrder(order string) YoutubeHelperSearchOption {
        return func(opts *youtubeHelperSearchOptions) {
                opts.order = order
        }
}

func YoutubeHelperSearchPublishAfter(publishedAfter time.Time) YoutubeHelperSearchOption {
        return func(opts *youtubeHelperSearchOptions) {
                opts.publishedAfter = publishedAfter
        }
}

func YoutubeHelperSearchPublishBefore(publishedBefore time.Time) YoutubeHelperSearchOption {
        return func(opts *youtubeHelperSearchOptions) {
                opts.publishedBefore = publishedBefore
        }
}

func YoutubeHelperSearchQ(q string) YoutubeHelperSearchOption {
        return func(opts *youtubeHelperSearchOptions) {
                opts.q = q
        }
}

func YoutubeHelperSearchRegionCode(regionCode string) YoutubeHelperSearchOption {
        return func(opts *youtubeHelperSearchOptions) {
                opts.regionCode = regionCode
        }
}

func YoutubeHelperSearchSafeSearch(safeSearch string) YoutubeHelperSearchOption {
        return func(opts *youtubeHelperSearchOptions) {
                opts.safeSearch = safeSearch
        }
}

func YoutubeHelperSearchTopicId(topicId string) YoutubeHelperSearchOption {
        return func(opts *youtubeHelperSearchOptions) {
                opts.topicId = topicId
        }
}

func YoutubeHelperSearchType(searchType string) YoutubeHelperSearchOption {
        return func(opts *youtubeHelperSearchOptions) {
                opts.searchType = searchType
        }
}

type SearchCallBack func(*youtube.SearchResult)

type youtubeHelperChannelsOptions struct {
	parts       []string
	categoryId  string
	forUsername string
	ids         []string
}

func defaultYoutubeHelperChannelsOptions() *youtubeHelperChannelsOptions {
        return &youtubeHelperChannelsOptions{
		parts:       []string{ "id" },
		categoryId:  "",
		forUsername: "",
		ids:         nil,
        }
}

type YoutubeHelperChannelsOption func(*youtubeHelperChannelsOptions)

func YoutubeHelperChannelsParts(parts ...string) YoutubeHelperChannelsOption {
        return func(opts *youtubeHelperChannelsOptions) {
                opts.parts = parts
        }
}

func YoutubeHelperChannelsCategoryId(categoryId string) YoutubeHelperChannelsOption {
        return func(opts *youtubeHelperChannelsOptions) {
                opts.categoryId = categoryId
        }
}

func YoutubeHelperChannelsForUsername(forUsername string) YoutubeHelperChannelsOption {
        return func(opts *youtubeHelperChannelsOptions) {
                opts.forUsername = forUsername
        }
}

func YoutubeHelperChannelsIds(ids []string) YoutubeHelperChannelsOption {
        return func(opts *youtubeHelperChannelsOptions) {
                opts.ids = ids
        }
}

type ChannelsCallBack func(*youtube.Channel)

type youtubeHelperVideosOptions struct {
	parts           []string
	chart           string
	ids             []string
	regionCode      string
	videoCategoryId string
}

func defaultYoutubeHelperVideosOptions() *youtubeHelperVideosOptions {
        return &youtubeHelperVideosOptions{
		parts:           []string{ "id" },
		chart:           "",
		ids:             nil,
		regionCode:      "",
		videoCategoryId: "",
        }
}

type YoutubeHelperVideosOption func(*youtubeHelperVideosOptions)

func YoutubeHelperVideosParts(parts ...string) YoutubeHelperVideosOption {
        return func(opts *youtubeHelperVideosOptions) {
                opts.parts = parts
        }
}

func YoutubeHelperVideosChart(chart string) YoutubeHelperVideosOption {
        return func(opts *youtubeHelperVideosOptions) {
                opts.chart = chart
        }
}

func YoutubeHelperVideosIds(ids []string) YoutubeHelperVideosOption {
        return func(opts *youtubeHelperVideosOptions) {
                opts.ids = ids
        }
}

func YoutubeHelperVideosRegionCode(regionCode string) YoutubeHelperVideosOption {
        return func(opts *youtubeHelperVideosOptions) {
                opts.regionCode = regionCode
        }
}

func YoutubeHelperVideosVideoCategoryId(videoCategoryId string) YoutubeHelperVideosOption {
        return func(opts *youtubeHelperVideosOptions) {
                opts.videoCategoryId = videoCategoryId
        }
}

type VideosCallBack func(*youtube.Video)

type YoutubeHelper struct {
	verbose bool
	apiKey  string
	timeout time.Duration
}

func (y *YoutubeHelper) SearchChannels(cb SearchCallBack, opts ...YoutubeHelperSearchOption) (error) {
	opts = append(opts, YoutubeHelperSearchType("channel"))
	return y.Search(cb, opts...)
}

func (y *YoutubeHelper) SearchPlayLists(cb SearchCallBack, opts ...YoutubeHelperSearchOption) (error) {
	opts = append(opts, YoutubeHelperSearchType("playlist"))
	return y.Search(cb, opts...)
}

func (y *YoutubeHelper) SearchVideosByChannelId(channelId string, cb SearchCallBack, opts ...YoutubeHelperSearchOption) (error) {
	opts = append(opts, YoutubeHelperSearchChannelId(channelId))
	if err := y.SearchVideos(cb, opts...); err != nil {
		return fmt.Errorf("can not search video (channelId = %v): %w", channelId, err)
	}
	return nil
}

func (y *YoutubeHelper) SearchVideos(cb SearchCallBack, opts ...YoutubeHelperSearchOption) (error) {
	opts = append(opts, YoutubeHelperSearchType("video"))
	return y.Search(cb, opts...)
}

func (y *YoutubeHelper) Search(cb SearchCallBack, opts ...YoutubeHelperSearchOption) (error) {
	baseOpts := defaultYoutubeHelperSearchOptions()
        for _, opt := range opts {
                opt(baseOpts)
        }
	pageToken := ""
	count := 0
	maxResults := 50
	for {
		ctx, cancel := context.WithTimeout(context.Background(), y.timeout)
		defer cancel()
		youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(y.apiKey))
		if err != nil {
			return fmt.Errorf("can not create youtube service: %w", err)
		}
		searchListCall := youtubeService.Search.List(baseOpts.parts)
		if baseOpts.relatedToVideoId != "" &&  baseOpts.searchType == "video" {
			searchListCall.RelatedToVideoId(baseOpts.relatedToVideoId)
		}
		if baseOpts.channelId != "" {
			searchListCall.ChannelId(baseOpts.channelId)
		}
		searchListCall.ChannelType(baseOpts.channelType)
		if baseOpts.eventType != "" {
			searchListCall.EventType(baseOpts.eventType)
		}
		if baseOpts.limit != 0 {
			maxResults = ((baseOpts.limit - count) % 50) + 1
		}
		searchListCall.MaxResults(int64(maxResults))
		searchListCall.Order(baseOpts.order)
		searchListCall.PageToken(pageToken)
		if !baseOpts.publishedAfter.IsZero() {
			publishedAfterStr := baseOpts.publishedAfter.UTC().Format(time.RFC3339)
			searchListCall.PublishedAfter(publishedAfterStr)
		}
		if !baseOpts.publishedBefore.IsZero() {
			publishedBeforeStr := baseOpts.publishedBefore.UTC().Format(time.RFC3339)
			searchListCall.PublishedBefore(publishedBeforeStr)
		}
		if baseOpts.q != "" {
			searchListCall.Q(baseOpts.q)
		}
		if baseOpts.regionCode != "" {
			searchListCall.VideoCaption(baseOpts.regionCode)
		}
		searchListCall.VideoCaption(baseOpts.safeSearch)
		if baseOpts.topicId != "" {
			searchListCall.VideoCaption(baseOpts.topicId)
		}
		if baseOpts.searchType != "" {
			searchListCall.Type(baseOpts.searchType)
		}
		searchListCall.VideoCaption(baseOpts.videoCaption)
		if baseOpts.videoCategoryId != "" {
			searchListCall.VideoCategoryId(baseOpts.videoCategoryId)
		}
                searchListCall.VideoDefinition(baseOpts.videoDefinition)
                searchListCall.VideoDimension(baseOpts.videoDimension)
                searchListCall.VideoDuration(baseOpts.videoDuration)
                searchListCall.VideoEmbeddable(baseOpts.videoEmbeddable)
                searchListCall.VideoLicense(baseOpts.videoLicense)
                searchListCall.VideoSyndicated(baseOpts.videoSyndicated)
                searchListCall.VideoType(baseOpts.videoType)
		searchListResponse, err := searchListCall.Do()
		if err != nil {
			return fmt.Errorf("can not search %v: %w", baseOpts.searchType, err)
		}
		if searchListResponse.Items == nil || len(searchListResponse.Items) == 0 {
			break
		}
		for _, item := range searchListResponse.Items {
			cb(item)
		}
		if searchListResponse.NextPageToken == "" {
			break
                }
		pageToken = searchListResponse.NextPageToken
		count += len(searchListResponse.Items)
	}
	return nil
}

func (y *YoutubeHelper) ChannelsByChannelIds(channelIds []string, cb ChannelsCallBack, opts ...YoutubeHelperChannelsOption) (error) {
	opts = append(opts, YoutubeHelperChannelsIds(channelIds))
	if err := y.Channels(cb, opts...); err != nil {
		return fmt.Errorf("can not get channels (channelIds = %+v): %w", opts, err)
	}
	return nil
}

func (y *YoutubeHelper) Channels(cb ChannelsCallBack, opts ...YoutubeHelperChannelsOption) (error) {
	baseOpts := defaultYoutubeHelperChannelsOptions()
        for _, opt := range opts {
                opt(baseOpts)
        }
	pageToken := ""
	for {
		ctx, cancel := context.WithTimeout(context.Background(), y.timeout)
		defer cancel()
		youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(y.apiKey))
		if err != nil {
			return fmt.Errorf("can not create youtube service: %w", err)
		}
		channelsListCall := youtubeService.Channels.List(baseOpts.parts)
		if baseOpts.categoryId != "" {
			channelsListCall.CategoryId(baseOpts.categoryId)
		}
		if baseOpts.forUsername != "" {
			channelsListCall.ForUsername(baseOpts.forUsername)
		}
		if baseOpts.ids != nil {
			channelsListCall.Id(baseOpts.ids...)
		}
		channelsListCall.MaxResults(50)
		channelsListCall.PageToken(pageToken)
		channelsListResponse, err := channelsListCall.Do()
		if err != nil {
			return fmt.Errorf("can not get channels: %w", err)
		}
		if channelsListResponse.Items == nil || len(channelsListResponse.Items) == 0 {
			break
		}
		for _, item := range channelsListResponse.Items {
			cb(item)
		}
		if channelsListResponse.NextPageToken == "" {
			break
                }
		pageToken = channelsListResponse.NextPageToken
	}
	return nil
}

func (y *YoutubeHelper) VideosByVideoIds(videoIds []string, cb VideosCallBack, opts ...YoutubeHelperVideosOption) (error) {
	opts = append(opts, YoutubeHelperVideosIds(videoIds))
	if err := y.Videos(cb, opts...); err != nil {
		return fmt.Errorf("can not get videos (videoIds = %+v): %w", opts, err)
	}
	return nil
}

func (y *YoutubeHelper) Videos(cb VideosCallBack, opts ...YoutubeHelperVideosOption) (error) {
	baseOpts := defaultYoutubeHelperVideosOptions()
        for _, opt := range opts {
                opt(baseOpts)
        }
	pageToken := ""
	for {
		ctx, cancel := context.WithTimeout(context.Background(), y.timeout)
		defer cancel()
		youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(y.apiKey))
		if err != nil {
			return fmt.Errorf("can not create youtube service: %w", err)
		}
		videosListCall := youtubeService.Videos.List(baseOpts.parts)
		if baseOpts.chart != "" {
			videosListCall.Chart(baseOpts.chart)
		}
		if baseOpts.ids != nil {
			videosListCall.Id(baseOpts.ids...)
		}
		if baseOpts.ids != nil {
			videosListCall.MaxResults(50)
			videosListCall.PageToken(pageToken)
		}
		if baseOpts.regionCode != "" {
			videosListCall.RegionCode(baseOpts.regionCode)
		}
		if baseOpts.videoCategoryId != "" {
			videosListCall.VideoCategoryId(baseOpts.videoCategoryId)
		}
		videosListResponse, err := videosListCall.Do()
		if err != nil {
			return fmt.Errorf("can not get videos: %w", err)
		}
		if videosListResponse.Items == nil || len(videosListResponse.Items) == 0 {
			break
		}
		for _, item := range videosListResponse.Items {
			cb(item)
		}
		if videosListResponse.NextPageToken == "" {
			break
                }
		pageToken = videosListResponse.NextPageToken
	}
	return nil
}

func NewYoutubeHelper(apiKey string, opts ...YoutubeHelperOption) *YoutubeHelper {
	baseOpts := defaultYoutubeHelperOptions()
	for _, opt := range opts {
		opt(baseOpts)
	}
	return &YoutubeHelper{
		verbose: baseOpts.verbose,
		apiKey:  apiKey,
		timeout: baseOpts.timeout,
	}
}
