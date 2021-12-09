package resources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func StateMachineLogGroup(stack constructs.Construct, groupName string) awslogs.LogGroup {
	return awslogs.NewLogGroup(stack, jsii.String("StateMachineLogGroup"), &awslogs.LogGroupProps{
		LogGroupName: jsii.String(groupName),
	})
}
