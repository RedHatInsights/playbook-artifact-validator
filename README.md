# Playbook Artifact Validator

Playbook Artifact Validator is a service that validates playbook run artifacts uploaded to cloud.redhat.com platform.

The service ensures that archives uploaded using [Ingress service](https://github.com/RedHatInsights/insights-ingress-go) contain valid [Ansible Runner](https://ansible-runner.readthedocs.io/en/stable/) output before other services do further processing of the archive.

## Details

The validator service consumes the `platform.upload.playbook` topic, where new uploads of the _playbook_ type are announced by the Ingress service.

Each upload is validated against a schema and the result (success/failure) is written to the `platform.upload.validation` topic.
If validation is successful, the uploaded archive is then made available to platform services.

## Development

### Prerequisities

* Golang >= 1.12
* docker-compose

### Setup

1. Run `sudo echo "127.0.0.1 kafka minio" >> /etc/hosts`
1. Run `docker-compose up` to start dependencies
1. Follow the steps [to create a new bucket and set up access policy](https://github.com/RedHatInsights/insights-ingress-go/tree/master/development#running)

### Running the service

Run `ACG_CONFIG=$(pwd)/cdappconfig.json go run main.go` to start the validator service

To test the service manually run `curl -v -F "file=@somefile.txt;type=application/vnd.redhat.playbook.somefile+tgz" -H "x-rh-identity: eyJpZGVudGl0eSI6IHsiYWNjb3VudF9udW1iZXIiOiAiMDAwMDAwMSIsICJpbnRlcm5hbCI6IHsib3JnX2lkIjogIjAwMDAwMSJ9fX0=" -H "x-rh-request_id: 012345" http://localhost:8080/api/ingress/v1/upload` where `somefile.txt` is the name of a file you wish to upload

### Running tests

`ACG_CONFIG=$(pwd)/cdappconfig.json go test -coverprofile cover.out ./...`
