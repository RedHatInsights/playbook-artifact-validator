# Playbook Artifact Validator

Playbook Artifact Validator is a service that validates playbook run artifacts uploaded to cloud.redhat.com platform.

The service ensures that archives uploaded using [Ingress service](https://github.com/RedHatInsights/insights-ingress-go) contain valid [Ansible Runner](https://ansible-runner.readthedocs.io/en/stable/) output before other services do further processing of the archive.

## Details

The validator service consumes the `platform.upload.playbook` topic, where new uploads of the _playbook_ type are announced by the Ingress service.

Each upload is validated against a schema and the result (success/failure) is written to the `platform.upload.validation` topic.
If validation is successful, the uploaded archive is then made available to platform services.

## Expected input format

The service expects the uploaded file to contain Ansible Runner [job events](https://ansible-runner.readthedocs.io/en/stable/intro.html#runner-artifact-job-events-host-and-playbook-events).
Job events should be stored in the [newline-delimited JSON](https://jsonlines.org/) format.
Each line in the file matches one job event.

```jsonl
{"event": "playbook_on_start", "uuid": "cb93301e-5ff8-4f75-ade6-57d0ec2fc662", "counter": 0, "stdout": "", "start_line": 0, "end_line": 0}
{"event": "playbook_on_stats", "uuid": "998a4bd2-2d6b-4c31-905c-2d5ad7a7f8ab", "counter": 1, "stdout": "", "start_line": 0, "end_line": 0}
```

The structure of each event is validated using a JSON Schema defined in [runner.yaml](./schemas/runner.yaml).
Note that additional attributes (not defined by the schema) are allowed.

The expected content type of the uploaded file is `application/vnd.redhat.playbook.events+jsonl` for a plain file or `application/vnd.redhat.playbook.events+tgz` if the content is compressed (not implemented yet).

## Development

### Prerequisities

* Golang >= 1.12
* docker-compose

### Setup

Follow the steps [to create a new bucket and set up access policy](https://github.com/RedHatInsights/insights-ingress-go/tree/master/development#running)

### Running the service

Run `docker-compose up --build` to start the service and its dependencies

To test the service manually run `make sample_upload`.
This uploads the `upload.jsonl` file via the ingress service.
Afterwards, look for `Payload valid` message in validator logs.

### Running tests

`make test`
