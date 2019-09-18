package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	c := New("./test.json")
	c.Load()

	if len(c.RssFeeds) != 2 {
		t.Errorf("Expected length of RSS Feeds to be 2, but got %d", len(c.RssFeeds))
	}
}

func TestGetCheckpoint(t *testing.T) {
	c := New("./test.json")
	c.Load()

	cp := c.GetCheckpoint("testCheckpoint")

	if cp.LastBuildDate != "Wed, 18 Sep 2019 07:02:58 +0000" {
		t.Error("Incorrect LastBuildDate returned")
	}

	if cp.LatestArticleDate != "Wed, 18 Sep 2019 07:02:58 +0000" {
		t.Error("Incorrect LatestArticleDate returned")
	}

}
