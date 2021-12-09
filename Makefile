.PHONY: build-tweet build load-env prepare

build-tweet:
	cd resources/lambda_functions/tweet && GOARCH=amd64 GOOS=linux go build -o bin/main

build: build-tweet

load-env:
	source ".env"

prepare: load-env build