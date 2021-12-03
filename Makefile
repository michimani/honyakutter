.PHONY: build-translate build-tweet build load-env prepare

build-translate:
	cd resources/lambda_functions/translate && GOARCH=amd64 GOOS=linux go build -o bin/main

build-tweet:
	cd resources/lambda_functions/tweet && GOARCH=amd64 GOOS=linux go build -o bin/main

build: build-translate build-tweet

load-env:
	source ".env"

prepare: load-env build