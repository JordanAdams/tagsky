package ingestor

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/bluesky-social/jetstream/pkg/models"
)

type PostRecord struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func handlePost(post models.Event) error {
	var record PostRecord
	json.Unmarshal(post.Commit.Record, &record)

	re := regexp.MustCompile(`#\w+`)
	tags := re.FindAllString(record.Text, -1)

	for _, t := range tags {
		fmt.Printf("%s - %s\n", t, post.Commit.CID)
	}

	return nil
}

func handleCommits(commit models.Event) error {
	if commit.Commit.Collection == "app.bsky.feed.post" {
		return handlePost(commit)
	}

	return nil
}
