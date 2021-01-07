run:
	ACG_CONFIG=$(shell pwd)/cdappconfig.json go run main.go

test:
	RUNNER_SCHEMA=$(shell pwd)/schemas/runner.yaml ACG_CONFIG=$(shell pwd)/cdappconfig.json go test -coverprofile cover.out ./...

sample_upload:
	curl -v -F "file=@upload.txt;type=application/vnd.redhat.playbook.upload+tgz" -H "x-rh-identity: eyJpZGVudGl0eSI6IHsiYWNjb3VudF9udW1iZXIiOiAiMDAwMDAwMSIsICJpbnRlcm5hbCI6IHsib3JnX2lkIjogIjAwMDAwMSJ9fX0=" -H "x-rh-request_id: testtesttest" http://localhost:8080/api/ingress/v1/upload
