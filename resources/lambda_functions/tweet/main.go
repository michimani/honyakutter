package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/michimani/gotwi"
	mt "github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

const (
	OAuthTokenEnvKeyName       = "GOTWI_ACCESS_TOKEN"
	OAuthTokenSecretEnvKeyName = "GOTWI_ACCESS_TOKEN_SECRET"
)

type TweetEvent struct {
	Text TweetText `json:"inputText"`
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

func newTiwtterClient() (*gotwi.Client, error) {
	in := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           os.Getenv(OAuthTokenEnvKeyName),
		OAuthTokenSecret:     os.Getenv(OAuthTokenSecretEnvKeyName),
	}

	return gotwi.NewClient(in)
}

func tweet(c *gotwi.Client, text string) (string, error) {
	p := &types.CreateInput{
		Text: gotwi.String(text),
	}

	res, err := mt.Create(context.Background(), c, p)
	if err != nil {
		return "", err
	}

	return gotwi.StringValue(res.Data.ID), nil
}

func main() {
	lambda.Start(handleRequest)
}
