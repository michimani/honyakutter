package resources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

var (
	stateMachineTimeout float64 = 300 * 1000 // 300 sec
	taskTimeout         float64 = 60 * 1000  // 60 sec
)

func TranslateTweetStateMachine(stack constructs.Construct, tweetFunc awslambda.Function) {
	initSt := awsstepfunctions.NewPass(stack, jsii.String("init"), &awsstepfunctions.PassProps{
		Comment: jsii.String("init state"),
	})

	// Translate state
	translateResultSelector := map[string]interface{}{
		"inputText.$": "$.TranslatedText",
	}
	translatSt := awsstepfunctionstasks.NewCallAwsService(stack, jsii.String("TranslateSDKIntegration"), &awsstepfunctionstasks.CallAwsServiceProps{
		Service: jsii.String("Translate"),
		Action:  jsii.String("translateText"),
		IamResources: &[]*string{
			jsii.String("*"),
		},
		IamAction: jsii.String("translate:TranslateText"),
		Parameters: &map[string]interface{}{
			"SourceLanguageCode": awsstepfunctions.JsonPath_StringAt(jsii.String("$.sourceLang")),
			"TargetLanguageCode": awsstepfunctions.JsonPath_StringAt(jsii.String("$.targetLang")),
			"Text":               awsstepfunctions.JsonPath_StringAt(jsii.String("$.inputText")),
		},
		ResultSelector: &translateResultSelector,
	})

	// Tweet state
	tweetSt := lambdaFunctionToTask(stack, tweetFunc)

	definition := initSt.Next(translatSt).Next(tweetSt)

	logGroup := StateMachineLogGroup(stack, "HonyakutterGoLogGroup")

	awsstepfunctions.NewStateMachine(stack, jsii.String("TranslateTweetStateMachine"), &awsstepfunctions.StateMachineProps{
		StateMachineName: jsii.String("honyakutter-go-translate-tweet-state-machine"),
		StateMachineType: awsstepfunctions.StateMachineType_EXPRESS,
		Timeout:          awscdk.Duration_Millis(&stateMachineTimeout),
		Definition:       definition,
		Logs: &awsstepfunctions.LogOptions{
			Destination: logGroup,
			Level:       awsstepfunctions.LogLevel_ALL,
		},
	})
}

func lambdaFunctionToTask(stack constructs.Construct, lfn awslambda.Function, resultSelectors ...map[string]interface{}) awsstepfunctionstasks.LambdaInvoke {
	props := awsstepfunctionstasks.LambdaInvokeProps{
		LambdaFunction: lfn,
		InvocationType: awsstepfunctionstasks.LambdaInvocationType_REQUEST_RESPONSE,
		Timeout:        awscdk.Duration_Millis(&taskTimeout),
	}

	if len(resultSelectors) > 0 {
		props.ResultSelector = &resultSelectors[0]
	}

	return awsstepfunctionstasks.NewLambdaInvoke(stack, lfn.FunctionName(), &props)
}
