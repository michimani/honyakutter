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

type resultSelectorItem struct {
	Key   string
	Value interface{}
}

func TranslateTweetStateMaschine(stack constructs.Construct, translateFunc, tweetFunc awslambda.Function) {
	initSt := awsstepfunctions.NewPass(stack, jsii.String("init"), &awsstepfunctions.PassProps{
		Comment: jsii.String("init state"),
	})

	// Translate state
	translateResultSelector := []resultSelectorItem{
		{Key: "inputText.$", Value: "$.Payload"},
	}
	translatSt := lambdaFunctionToTask(stack, translateFunc, translateResultSelector...)

	// Tweet state
	tweetSt := lambdaFunctionToTask(stack, tweetFunc)

	definition := initSt.Next(translatSt).Next(tweetSt)

	awsstepfunctions.NewStateMachine(stack, jsii.String("TranslateTweetStateMaschine"), &awsstepfunctions.StateMachineProps{
		StateMachineName: jsii.String("translate-tweet-state-maschine"),
		StateMachineType: awsstepfunctions.StateMachineType_EXPRESS,
		Timeout:          awscdk.Duration_Millis(&stateMachineTimeout),
		Definition:       definition,
	})
}

func lambdaFunctionToTask(stack constructs.Construct, lfn awslambda.Function, resultSelectors ...resultSelectorItem) awsstepfunctionstasks.LambdaInvoke {
	props := awsstepfunctionstasks.LambdaInvokeProps{
		LambdaFunction: lfn,
		InvocationType: awsstepfunctionstasks.LambdaInvocationType_REQUEST_RESPONSE,
		Timeout:        awscdk.Duration_Millis(&taskTimeout),
	}

	rs := map[string]interface{}{}
	for _, rsi := range resultSelectors {
		rs[rsi.Key] = rsi.Value
	}

	if len(rs) > 0 {
		props.ResultSelector = &rs
	}

	return awsstepfunctionstasks.NewLambdaInvoke(stack, lfn.FunctionName(), &props)
}
