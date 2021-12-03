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

func TranslateTweetStateMaschine(stack constructs.Construct, translateFunc, tweetFunc awslambda.Function) {
	chain := awsstepfunctions.NewPass(stack, jsii.String("init"), &awsstepfunctions.PassProps{
		Comment: jsii.String("init state"),
	})

	// Translate state
	chain.Next(lambdaFunctionToTask(stack, translateFunc))

	awsstepfunctions.NewStateMachine(stack, jsii.String("TranslateTweetStateMaschine"), &awsstepfunctions.StateMachineProps{
		StateMachineName: jsii.String("translate-tweet-state-maschine"),
		StateMachineType: awsstepfunctions.StateMachineType_EXPRESS,
		Timeout:          awscdk.Duration_Millis(&stateMachineTimeout),
		Definition:       chain,
	})
}

func lambdaFunctionToTask(stack constructs.Construct, lfn awslambda.Function) awsstepfunctionstasks.LambdaInvoke {
	return awsstepfunctionstasks.NewLambdaInvoke(stack, lfn.FunctionName(), &awsstepfunctionstasks.LambdaInvokeProps{
		LambdaFunction: lfn,
		InvocationType: awsstepfunctionstasks.LambdaInvocationType_REQUEST_RESPONSE,
		Timeout:        awscdk.Duration_Millis(&taskTimeout),
		ResultSelector: &map[string]interface{}{
			"inputType.$": "$.Payload",
		},
	})
}
