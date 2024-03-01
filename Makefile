compile = env GOOS=linux  GOARCH=arm64  go build -v -ldflags '-s -w -v' -tags lambda.norpc -o
zipper = zip -j -r
test_to_file = go test -coverprofile=coverage.out
percent = go tool cover -func=coverage.out | sed 's~\([^/]\{1,\}/\)\{3\}~~' | sed '$d' | sort -g -r -k 3
profile = pet

build: clean gomodgen import
	$(compile) bin/cmd/createProcessorHandler/bootstrap cmd/createProcessorHandler/create_processor_handler.go

zip:
	$(zipper) bin/cmd/createProcessorHandler/createProcessorHandler.zip bin/cmd/createProcessorHandler/bootstrap

clean:
	go clean
	rm -rf ./bin ./vendor go.sum

deploy: build zip
	sls deploy --aws-profile $(profile) --verbose

undeploy:
	sls remove --aws-profile $(profile) --verbose

import:
	go mod download github.com/aws/aws-lambda-go
	go mod download github.com/diegocabrera89/ms-payment-core

	go get github.com/diegocabrera89/ms-payment-core/dynamodbcore
	go get github.com/diegocabrera89/ms-payment-core/helpers
	go get github.com/diegocabrera89/ms-payment-processor/utils

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
