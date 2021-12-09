honyakutter
===

This is an application that translates text entered in Japanese into English and tweets it, defined using AWS CDK v2 (golang) and using AWS Step Functions to connect the translation and tweeting processes.

# Preparation

1. Check your AWS CDK version.

    ```bash
    cdk --version
    ```
    
    If you have not installed it, please use the following command to install it.
    
    ```bash
    npm install -g aws-cdk
    ```

    This application is intended for use with AWS CDK `v2.0.0` or higher.

2. Getting Twitter API's some tokens.

    ⚠️ You will need the Twitter API key and secret, as well as the access token and access token secret. Please create an app on the Twitter Developer page and obtain each token.

1. Create `.env` file.

    ```bash
    cp .env.sample .env
    ```

    And, replace values of each environment variables.

2. Load environment variables.

    ```bash
    source .env
    ```

# Build

1. Build Tweet Lambda Function

    ```bash
    cd resources/lambda_functions/tweet \
    && GOARCH=amd64 GOOS=linux go build -o bin/main
    ```

# Deploying

1. bootstrap 

    ```bash
    cdk bootstrap
    ```

1. Build Lambda functions

    ```bash
    make build-tweet
    ```

1. Load environment values 

    ```bash
    source .env
    ```

1. Generate CFn template

    ```bash
    cdk synth
    ```

1. deploy

    ```bash
    cdk deploy
    ```


# Testing

```bash
go test .
```

# Manual execution

Execute Lambda functions and Step Functions state machine manually by using AWS CLI.. (The latest versions at the time of this writing are `v2.4.4` and `1.22.18`.)

## Tweet Lambda Function

```bash
aws lambda invoke \
--function-name honyakutter-go-tweet-function \
--invocation-type Event \
--region ap-northeast-1 \
--payload fileb://testdata/tweet_lambda_payload.json \
out
```

## Start state machine

1. Get state machine ARN

    ```bash
    STATEMACHINE_ARN=$(
        aws stepfunctions list-state-machines \
        --query "stateMachines[?name=='honyakutter-go-translate-tweet-state-machine'].stateMachineArn" \
        --output text
    ) && echo "${STATEMACHINE_ARN}"
    ```

2. execute

    ```bash
    aws stepfunctions start-execution \
    --state-machine-arn "${STATEMACHINE_ARN}" \
    --input file://testdata/statemachine_input.json
    ```

# Licence

[MIT](https://github.com/michimani/honyakutter/blob/main/LICENCE)

# Author

[michimani210](https://twitter.com/michimani210)

