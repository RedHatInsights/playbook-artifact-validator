run:
	ACG_CONFIG=$(shell pwd)/cdappconfig.json go run main.go

test:
	RUNNER_SCHEMA=$(shell pwd)/schemas/runner.yaml ACG_CONFIG=$(shell pwd)/cdappconfig.json go test -coverprofile cover.out ./...
