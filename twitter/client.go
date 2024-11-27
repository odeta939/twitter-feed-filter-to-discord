package twitter

import (
	"context"
	"fmt"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/searchtweet"
	searchTypes "github.com/michimani/gotwi/tweet/searchtweet/types"
)

func GetClient(config *Config) (*gotwi.Client, error) {
	if config.BearerToken == "" {
		return nil, fmt.Errorf("no token found")
	}

	in := &gotwi.NewClientWithAccessTokenInput{
		AccessToken: config.BearerToken,
	}

	client, err := gotwi.NewClientWithAccessToken(in)
	if err != nil {
		return nil, fmt.Errorf("failed to create Twitter client: %w", err)
	}

	return client, nil
}

func GetRecentTweets(client *gotwi.Client, query string) (*searchTypes.ListRecentOutput, error) {
	p := &searchTypes.ListRecentInput{
			MaxResults:  10,
			Query:       query,
		}

	res, err := searchtweet.ListRecent(context.Background(), client, p)
	if err != nil {
		return nil, fmt.Errorf("failed to create Twitter client: %w", err)
	}

	return res, nil
}