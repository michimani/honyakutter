package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweets"
	"github.com/michimani/gotwi/tweets/types"
)

const (
	OAuthTokenEnvKeyName       = "GOTWI_ACCESS_TOKEN"
	OAuthTokenSecretEnvKeyName = "GOTWI_ACCESS_TOKEN_SECRET"
)

type TweetEvent struct {
	Text TweetText `json:"tweetText"`
}

type TweetText string

func (t TweetText) withCurrentTime() string {
	now := time.Now()
	return fmt.Sprintf("[%s] %s", now, t)
}

func handleRequest(ctx context.Context, event TweetEvent) (string, error) {
	fmt.Printf("%#+v\n", event)

	c, err := newTiwtterClient()
	if err != nil {
		return "", err
	}

	tweetText := event.Text.withCurrentTime()

	tweetID, err := tweet(c, tweetText)
	if err != nil {
		return "", err
	}

	return tweetID, nil
}

func newTiwtterClient() (*gotwi.GotwiClient, error) {
	in := &gotwi.NewGotwiClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           os.Getenv(OAuthTokenEnvKeyName),
		OAuthTokenSecret:     os.Getenv(OAuthTokenSecretEnvKeyName),
	}

	return gotwi.NewGotwiClient(in)
}

func tweet(c *gotwi.GotwiClient, text string) (string, error) {
	p := &types.ManageTweetsPostParams{
		Text: gotwi.String(text),
	}

	res, err := tweets.ManageTweetsPost(context.Background(), c, p)
	if err != nil {
		return "", err
	}

	return gotwi.StringValue(res.Data.ID), nil
}

func main() {
	lambda.Start(handleRequest)
}
