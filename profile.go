package main

import (
	"regexp"
	"strconv"
	"strings"
	"testing"
)

type FlatChannel struct {
	Id string `json:"id"`
}

type FlatPost struct {
	FlatChannelId string `json:"channel_id"`
	Message       string `json:"message"`
	Hashtags      string `json:"hashtags"`
}

func (p FlatPost) FlatParseTags() (*FlatPost, error) {
	var x, err = regexp.Compile(`#[^\s]{,999}`)
	if err != nil {
		return nil, err
	}
	for _, hashtag := range x.FindAll([]byte(p.Message), -1) {
		if !strings.Contains(p.Hashtags, string(hashtag)) {
			p.Hashtags = p.Hashtags + string(hashtag)
		}
	}
	return &p, nil
}

func NewId(prev string) string {
	i, _ := strconv.Atoi(prev)
	return strconv.Itoa(i + 1)
}

func TestGetFlatProfile(t *testing.T) {
	var flat_channel = FlatChannel{"000aaa"}
	var flat_post = &FlatPost{flat_channel.Id, "Hello #world, I am #here to learn about #apis", ""}
	flat_post, _ = flat_post.FlatParseTags()
	_ = flat_post
}
