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

Follow the steps [to create a new bucket and set up access policy](https://github.com/RedHatInsights/insights-ingress-go/tree/master/development#running)

### Running the service

Run `docker-compose up --build` to start the service and its dependencies

To test the service manually run `make sample_upload`.
This uploads the `upload.txt` file via the ingress service.
Afterwards, look for `Payload valid` message in validator logs.

### Running tests

`make test`
