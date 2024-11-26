package main

import (
	"bufio"
	"cmp"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"slices"
	"strings"
	"time"
)

const videosFile = "static/videos.json"

//go:embed static/videos.json
var videosJSON []byte

type video struct {
	ID            string   `json:"id"`
	Published     string   `json:"published"`
	Title         string   `json:"title"`
	OriginalTitle string   `json:"originalTitle"`
	Channel       string   `json:"channel"`
	ChannelID     string   `json:"channelId"`
	Description   string   `json:"description"`
	Tags          []string `json:"tags"`

	Date string `json:"-"`
}

func (t video) getInfoURL(youtubeKey string) string {
	return fmt.Sprintf(
		"https://www.googleapis.com/youtube/v3/videos?part=snippet&id=%s&key=%s",
		t.ID,
		youtubeKey,
	)
}

func (t *video) update(youtubeKey string) error {
	if t.Published != "" {
		//return nil
	}

	resp, err := http.Get(t.getInfoURL(youtubeKey))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	var wrapper struct {
		Items []struct {
			Snippet struct {
				PublishedAt  string `json:"publishedAt"`
				ChannelID    string `json:"channelID"`
				ChannelTitle string `json:"channelTitle"`
				Title        string `json:"title"`
				Thumnails    struct {
					Standard struct {
						URL string `json:"url"`
					} `json:"standard"`
				} `json:"thumbnails"`
			} `json:"snippet"`
		} `json:"items"`
	}

	err = dec.Decode(&wrapper)
	if err != nil {
		return err
	}

	if len(wrapper.Items) == 0 {
		return errors.New("no items returned from youtube")
	}

	sn := wrapper.Items[0].Snippet

	t.Published = sn.PublishedAt
	t.ChannelID = sn.ChannelID
	t.Channel = sn.ChannelTitle
	t.OriginalTitle = sn.Title

	if t.Title == "" {
		t.Title = sn.Title
	}

	if t.Tags == nil {
		t.Tags = make([]string, 0)
	}

	return nil
}

func getVideos() ([]*video, error) {

	var videos []*video
	err := json.Unmarshal(videosJSON, &videos)

	if err != nil {
		return videos, err
	}

	slices.SortFunc(videos, func(a, b *video) int {
		return cmp.Compare(b.Published, a.Published)
	})

	for _, video := range videos {
		t, err := time.Parse(time.RFC3339, video.Published)
		if err != nil {
			return videos, err
		}

		video.Date = t.Format("January 2, 2006")
	}

	return videos, err
}

func updateVideos() error {

	youtubeKey := os.Getenv("YOUTUBE_KEY")
	if youtubeKey == "" {
		return errors.New("YOUTUBE_KEY environment variable must be set")
	}

	f, err := os.OpenFile(videosFile, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	dec := json.NewDecoder(f)

	var videos []*video
	err = dec.Decode(&videos)
	if err != nil {
		return err
	}

	for _, video := range videos {
		err = video.update(youtubeKey)
		if err != nil {
			return err
		}
	}

	j, err := json.MarshalIndent(videos, "", "  ")
	if err != nil {
		return err
	}

	fmt.Printf("%s\n\n", j)

	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("Write this to the videos file? y/n")
	sc.Scan()
	answer := strings.TrimSpace(sc.Text())

	if answer != "y" {
		fmt.Println("aborting")
		return nil
	}

	err = f.Truncate(0)
	if err != nil {
		return err
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = f.Write(j)
	if err != nil {
		return err
	}

	return nil
}
