package rss

import (
	"fmt"
	"sort"
	"strings"

	"github.com/ungerik/go-rss"

	"github.com/realcodetips/rsstwitter/internal/config"
	"github.com/realcodetips/rsstwitter/internal/twitter"
)

func removeWhitespace(str string) string {
	return strings.NewReplacer(" ", "", "\n", "").Replace(str)
}

func processArticles(articles []rss.Item, cp *config.Checkpoint) error {
	// Sort posts from oldest -> newest
	sort.Slice(articles, func(i, j int) bool {
		id, _ := articles[i].PubDate.Parse()
		jd, _ := articles[j].PubDate.Parse()
		return id.Before(jd)
	})

	for _, article := range articles {
		if cp.LatestArticleDate == "" {
			err := twitter.SendTweet(article)
			if err != nil {
				return fmt.Errorf("Error tweeting article - %s - %+v", article.Title, err)
			}

			cp.LatestArticleDate = article.PubDate
		} else {
			lp, err := cp.LatestArticleDate.Parse()
			if err != nil {
				return err
			}

			pd, err := article.PubDate.Parse()
			if err != nil {
				return err
			}

			if pd.After(lp) {
				err := twitter.SendTweet(article)
				if err != nil {
					return fmt.Errorf("Error tweeting article - %s - %+v", article.Title, err)
				}
				cp.LatestArticleDate = article.PubDate
			}
		}
	}

	return nil
}

// ProcessRSS comment
func ProcessRSS(url string, conf *config.Config) error {
	channel, err := rss.Read(url)
	if err != nil {
		return err
	}

	name := removeWhitespace(channel.Title)
	checkpoint := conf.GetCheckpoint(name)

	// If the RSS hasn't been used before, go ahead
	// and process the articles.
	if checkpoint.LastBuildDate == "" {
		err = processArticles(channel.Item, checkpoint)
		if err != nil {
			return err
		}

		checkpoint.LastBuildDate = channel.LastBuildDate
	} else {
		// Otherwise we need to do some extra checks
		// so we don't reprocess an RSS feed that hasn't
		// changed
		rssBuild, err := channel.LastBuildDate.Parse()
		if err != nil {
			return fmt.Errorf("Unable to parse rss build date: %+v", err)
		}

		checkpointBuild, err := checkpoint.LastBuildDate.Parse()
		if err != nil {
			return fmt.Errorf("Unable to parse checkpoint build date: %+v", err)
		}

		if rssBuild.After(checkpointBuild) {
			err = processArticles(channel.Item, checkpoint)
			if err != nil {
				return err
			}

			checkpoint.LastBuildDate = channel.LastBuildDate
		}
	}

	return nil
}
