package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ungerik/go-rss"
)

// Config is top level object that the config gets
// loaded into
type Config struct {
	Filename    string                 `json:"-"`
	RssFeeds    []string               `json:"rssFeeds"`
	Checkpoints map[string]*Checkpoint `json:"checkpoints"`
}

// Checkpoint holds the information for each
// rss feed.
type Checkpoint struct {
	// Most recent time the content of the channel was modified
	LastBuildDate rss.Date `json:"lastBuildDate"`
	// The date of the last processed article
	LatestArticleDate rss.Date `json:"latestArticleDate"`
}

// GetCheckpoint takes a checkpoint key and returns the associated
// Checkpoint. It will create a new Checkpoint if it doesn't exist.
// Returns a pointer to Checkpoint so it can be amended by the caller.
func (c *Config) GetCheckpoint(key string) *Checkpoint {
	if _, ok := c.Checkpoints[key]; !ok {
		c.Checkpoints[key] = &Checkpoint{}
	}

	return c.Checkpoints[key]
}

// New creates a pointer to a new Config object
func New(filename string) (c *Config) {
	return &Config{Filename: filename}
}

// Load loads the config file and returns a pointer to a
// Config
func (c *Config) Load() error {
	jsonFile, err := os.Open(c.Filename)
	if err != nil {
		return fmt.Errorf("Unable to open config file: %+v", err)
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("Unable to read config file: %+v", err)
	}

	return json.Unmarshal(byteValue, &c)
}

// Save saves the current Config object to
// the original filename
func (c *Config) Save() error {
	j, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return fmt.Errorf("Unable to marshal config: %+v", err)
	}

	return ioutil.WriteFile(c.Filename, j, 0644)
}
