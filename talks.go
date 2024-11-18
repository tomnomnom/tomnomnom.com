package main

import (
	"cmp"
	_ "embed"
	"encoding/json"
	"slices"
)

//go:embed static/talks.json
var talksJSON []byte

type talk struct {
	ID        string `json:"id"`
	Published string `json:"published"`
	Title     string `json:"title"`
	Channel   string `json:"channel"`
	ChannelID string `json:"channelId"`
}

func getTalks() ([]talk, error) {

	var talks []talk
	err := json.Unmarshal(talksJSON, &talks)

	if err != nil {
		return talks, err
	}

	slices.SortFunc(talks, func(a, b talk) int {
		return cmp.Compare(b.Published, a.Published)
	})

	return talks, err
}
