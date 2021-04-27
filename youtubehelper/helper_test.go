package youtubehelper

import (
	"fmt"
	"time"
	"testing"
	"io/ioutil"
	"google.golang.org/api/youtube/v3"
)

func readKey(filePath string) (string) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func TestSearchChannels(t *testing.T) {
	apikey := readKey("./secret")
	helper := NewYoutubeHelper(apikey, YoutubeHelperVerbose(true), YoutubeHelperTimeout(30 * time.Second))
	err := helper.SearchChannels(
		func(result *youtube.SearchResult) {
			fmt.Printf("%v, %v\n", result.Id.Kind, result.Id.ChannelId)
		},
		YoutubeHelperSearchLimit(50),
	)
	if err != nil {
		t.Error(err)
	}
}

func TestSearchPlayLists(t *testing.T) {
	apikey := readKey("./secret")
	helper := NewYoutubeHelper(apikey, YoutubeHelperVerbose(true), YoutubeHelperTimeout(30 * time.Second))
	err := helper.SearchPlayLists(
		func(result *youtube.SearchResult) {
			fmt.Printf("%v, %v\n", result.Id.Kind, result.Id.PlaylistId)
		},
		YoutubeHelperSearchLimit(50),
	)
	if err != nil {
		t.Error(err)
	}
}

func TestSearchVideos(t *testing.T) {
	apikey := readKey("./secret")
	helper := NewYoutubeHelper(apikey, YoutubeHelperVerbose(true), YoutubeHelperTimeout(30 * time.Second))
	err := helper.SearchVideos(
		func(result *youtube.SearchResult) {
			fmt.Printf("%v, %v\n", result.Id.Kind, result.Id.VideoId)
		},
		YoutubeHelperSearchRelatedToVideoId("duhLuCOepso"),
	)
	if err != nil {
		t.Error(err)
	}
}

func TestSearchVideosByChannelId(t *testing.T) {
	apikey := readKey("./secret")
	helper := NewYoutubeHelper(apikey, YoutubeHelperVerbose(true), YoutubeHelperTimeout(30 * time.Second))
	err := helper.SearchVideosByChannelId(
		"UCv1fFr156jc65EMiLbaLImw",
		func(result *youtube.SearchResult) {
			fmt.Printf("%v, %v\n", result.Id.Kind, result.Id.VideoId)
		},
	)
	if err != nil {
		t.Error(err)
	}
}

func TestChannels(t *testing.T) {
	apikey := readKey("./secret")
	helper := NewYoutubeHelper(apikey, YoutubeHelperVerbose(true), YoutubeHelperTimeout(30 * time.Second))
	err := helper.Channels(
		func(result *youtube.Channel) {
			fmt.Printf("%v\n", result.Id)
		},
		YoutubeHelperChannelsForUsername("test"),
	)
	if err != nil {
		t.Error(err)
	}
}

func TestChannelsByChannelIds(t *testing.T) {
	count := 0
	apikey := readKey("./secret")
	helper := NewYoutubeHelper(apikey, YoutubeHelperVerbose(true), YoutubeHelperTimeout(30 * time.Second))
	err := helper.ChannelsByChannelIds(
		[]string{"UCv1fFr156jc65EMiLbaLImw", "UCD-miitqNY3nyukJ4Fnf4_A"},
		func(result *youtube.Channel) {
			fmt.Printf("%v\n", result.Id)
			count += 1
		},
	)
	if err != nil {
		t.Error(err)
	}
	if count != 2 {
		t.Error(err)
	}
}

func TestVideos(t *testing.T) {
	apikey := readKey("./secret")
	helper := NewYoutubeHelper(apikey, YoutubeHelperVerbose(true), YoutubeHelperTimeout(30 * time.Second))
	err := helper.Videos(
		func(result *youtube.Video) {
			fmt.Printf("%v\n", result.Id)
		},
		YoutubeHelperVideosChart("mostPopular"),
	)
	if err != nil {
		t.Error(err)
	}
}

func TestVideosByVideoIds(t *testing.T) {
	count := 0
	apikey := readKey("./secret")
	helper := NewYoutubeHelper(apikey, YoutubeHelperVerbose(true), YoutubeHelperTimeout(30 * time.Second))
	err := helper.VideosByVideoIds(
		[]string{ "duhLuCOepso", "qHnRdR3CKyI" },
		func(result *youtube.Video) {
			fmt.Printf("%v\n", result.Id)
			count += 1
		},
	)
	if err != nil {
		t.Error(err)
	}
	if count != 2 {
		t.Error(err)
	}
}


