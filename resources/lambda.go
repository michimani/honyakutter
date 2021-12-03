package resources

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

const (
	EnvKeyOfTwitterAPIKey            = "TWITTER_API_KEY"
	EnvKeyOfTwitterAPIKeySecret      = "TWITTER_API_KEY_SECRET"
	EnvKeyOfTwitterAccessToken       = "TWITTER_ACCESS_TOKEN"
	EnvKeyOfTwitterAccessTokenSecret = "TWITTER_ACCESS_TOKEN_SECRET"
)

var (
	memory  float64 = 128
	timeout float64 = 60 * 1000 // 30 sec
)

func TweetLambdaFunction(stack constructs.Construct) awslambda.Function {
	return awslambda.NewFunction(stack, jsii.String("TweetLambdaFunction"), &awslambda.FunctionProps{
		FunctionName: jsii.String("tweet-function"),
		Description:  jsii.String("Tweet text with current time."),
		Runtime:      awslambda.Runtime_GO_1_X(),
		Handler:      jsii.String("main"),
		Code:         awslambda.AssetCode_FromAsset(jsii.String("./resources/lambda_functions/tweet/bin"), nil),
		MemorySize:   &memory,
		Timeout:      awscdk.Duration_Millis(&timeout),
		Environment: &map[string]*string{
			"GOTWI_API_KEY":             jsii.String(os.Getenv(EnvKeyOfTwitterAPIKey)),
			"GOTWI_API_KEY_SECRET":      jsii.String(os.Getenv(EnvKeyOfTwitterAPIKeySecret)),
			"GOTWI_ACCESS_TOKEN":        jsii.String(os.Getenv(EnvKeyOfTwitterAccessToken)),
			"GOTWI_ACCESS_TOKEN_SECRET": jsii.String(os.Getenv(EnvKeyOfTwitterAccessTokenSecret)),
		},
	})
}

func TranslateLambdaFunction(stack constructs.Construct) awslambda.Function {
	lambdaFn := awslambda.NewFunction(stack, jsii.String("TranslateLambdaFunction"), &awslambda.FunctionProps{
		FunctionName: jsii.String("translate-function"),
		Description:  jsii.String("Translate text Japanese into English (default)."),
		Runtime:      awslambda.Runtime_GO_1_X(),
		Handler:      jsii.String("main"),
		Code:         awslambda.AssetCode_FromAsset(jsii.String("./resources/lambda_functions/translate/bin"), nil),
		MemorySize:   &memory,
		Timeout:      awscdk.Duration_Millis(&timeout),
	})

	lambdaFn.AddToRolePolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Effect: awsiam.Effect_ALLOW,
		Actions: &[]*string{
			jsii.String("translate:*"),
		},
		Resources: &[]*string{
			jsii.String("*"),
		},
	}))

	return lambdaFn
}
