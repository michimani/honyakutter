honyakutter
===

This is an application that translates text entered in Japanese into English and tweets it, defined using AWS CDK v2 (golang) and using AWS Step Functions to connect the translation and tweeting processes.

# TODO

- [ ] Create a Lambda function to translate Japanese into English.
- [ ] Create a Step Functions state machine that connects the translation and tweeting Lambda functions.

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

1. Build Lambda Function (that tweet a text)

    ```bash
    cd resources/lambda_functions/tweet \
    && GOARCH=amd64 GOOS=linux go build -o bin/main
    ```

# Deploying

1. bootstrap 

    ```bash
    cdk bootstrap
    ```
    
2. Generate CFn template

    ```bash
    cdk synth
    ```

3. deploy

    ```bash
    cdk deploy
    ```


# Testing

```bash
go test .
```

# Manual execution of Lambda functions

1. Tweet Lambda Function

    Using AWS CLI. (The latest versions at the time of this writing are `v2.4.4` and `1.22.18`.)

    ```bash
    aws lambda invoke \
    --function-name tweet-function \
    --invocation-type Event \
    --region ap-northeast-1 \
    --payload fileb://testdata/lambda_payload.json \
    out
    ```

# Licence

[MIT](https://github.com/michimani/gotwi/blob/main/LICENCE)

# Author

[michimani210](https://twitter.com/michimani210)

