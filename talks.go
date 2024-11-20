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

const talksFile = "static/talks.json"

//go:embed static/talks.json
var talksJSON []byte

type talk struct {
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

func (t talk) getInfoURL(youtubeKey string) string {
	return fmt.Sprintf(
		"https://www.googleapis.com/youtube/v3/videos?part=snippet&id=%s&key=%s",
		t.ID,
		youtubeKey,
	)
}

func (t *talk) update(youtubeKey string) error {
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

func getTalks() ([]*talk, error) {

	var talks []*talk
	err := json.Unmarshal(talksJSON, &talks)

	if err != nil {
		return talks, err
	}

	slices.SortFunc(talks, func(a, b *talk) int {
		return cmp.Compare(b.Published, a.Published)
	})

	for _, talk := range talks {
		t, err := time.Parse(time.RFC3339, talk.Published)
		if err != nil {
			return talks, err
		}

		talk.Date = t.Format("January 2, 2006")
	}

	return talks, err
}

func updateTalks() error {

	youtubeKey := os.Getenv("YOUTUBE_KEY")
	if youtubeKey == "" {
		return errors.New("YOUTUBE_KEY environment variable must be set")
	}

	f, err := os.OpenFile(talksFile, os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	dec := json.NewDecoder(f)

	var talks []*talk
	err = dec.Decode(&talks)
	if err != nil {
		return err
	}

	for _, talk := range talks {
		err = talk.update(youtubeKey)
		if err != nil {
			return err
		}
	}

	j, err := json.MarshalIndent(talks, "", "  ")
	if err != nil {
		return err
	}

	fmt.Printf("%s\n\n", j)

	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("Write this to the talks file? y/n")
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
