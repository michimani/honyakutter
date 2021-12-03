package main

import (
	"honyakutter/resources"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type HonyakutterStackProps struct {
	awscdk.StackProps
}

func NewHonyakutterStack(scope constructs.Construct, id string, props *HonyakutterStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Translate Lambda Function
	translateFn := resources.TranslateLambdaFunction(stack)

	// Tweet Lambda Function
	tweetFn := resources.TweetLambdaFunction(stack)

	// StateMaschine
	resources.TranslateTweetStateMaschine(stack, translateFn, tweetFn)

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewHonyakutterStack(app, "HonyakutterStack", &HonyakutterStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return nil
}
