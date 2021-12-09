package main

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	assertions "github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

func TestHonyakutterStack_TweetLambdaFunction(t *testing.T) {
	// Environment value for test
	t.Setenv("TWITTER_API_KEY", "twitter_api_key_for_test")
	t.Setenv("TWITTER_API_KEY_SECRET", "twitter_api_key_secret_for_test")
	t.Setenv("TWITTER_ACCESS_TOKEN", "twitter_access_token_for_test")
	t.Setenv("TWITTER_ACCESS_TOKEN_SECRET", "twitter_access_token_secret_for_test")

	app := awscdk.NewApp(nil)
	stack := NewHonyakutterStack(app, "TestStack", nil)
	template := assertions.Template_FromStack(stack)

	// Lambda Function
	template.HasResourceProperties(jsii.String("AWS::Lambda::Function"), map[string]interface{}{
		"FunctionName": "honyakutter-go-tweet-function",
		"Description":  "Tweet text with current time.",
		"Runtime":      "go1.x",
		"Handler":      "main",
		"MemorySize":   128,
		"Timeout":      60,
		"Environment": map[string]interface{}{
			"Variables": map[string]string{
				"GOTWI_API_KEY":             "twitter_api_key_for_test",
				"GOTWI_API_KEY_SECRET":      "twitter_api_key_secret_for_test",
				"GOTWI_ACCESS_TOKEN":        "twitter_access_token_for_test",
				"GOTWI_ACCESS_TOKEN_SECRET": "twitter_access_token_secret_for_test",
			},
		},
	})
}

func TestHonyakutterStack_TranslateTweetStateMachine(t *testing.T) {
	app := awscdk.NewApp(nil)
	stack := NewHonyakutterStack(app, "TestStack", nil)
	template := assertions.Template_FromStack(stack)

	// Lambda Function
	template.HasResourceProperties(jsii.String("AWS::StepFunctions::StateMachine"), map[string]interface{}{
		"StateMachineName": "honyakutter-go-translate-tweet-state-machine",
		"StateMachineType": "EXPRESS",
	})
}
