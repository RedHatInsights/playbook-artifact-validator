package handler

import (
	"playbook-artifact-validator/ingress"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validator", func() {

	DescribeTable("Rejects invalid files",
		func(file string) {
			request := &ingress.Request{}
			response := validateArtifacts(request, []byte(file))
			Expect(response.Validation).To(Equal("failure"))
		},
		Entry("empty file", ""),
		Entry("whitespace-only file", "         "),
		Entry("newline-only file", "\n\n\n\n"),
		Entry("invalid JSON (trailing braces)", `{"event": "playbook_on_start", "uuid": "60049c81-4f4b-41a0-bcf4-84399bf1b693", "counter": 0, "stdout": "", "start_line": 0, "end_line": 0}}`),
		Entry("missing uuid", `{"event": "playbook_on_start", "counter": 0, "stdout": "", "start_line": 0, "end_line": 0}`),
		Entry("incorrect uuid format", `{"event": "playbook_on_start", "uuid": "abcd", "counter": 0, "stdout": "", "start_line": 0, "end_line": 0}`),
	)

	DescribeTable("Accepts valid files",
		func(file string) {
			request := &ingress.Request{}
			response := validateArtifacts(request, []byte(file))
			Expect(response.Validation).To(Equal("success"))
		},

		Entry("multiple events", `
		{"event": "playbook_on_start", "uuid": "cb93301e-5ff8-4f75-ade6-57d0ec2fc662", "counter": 0, "stdout": "", "start_line": 0, "end_line": 0}
		{"event": "playbook_on_stats", "uuid": "998a4bd2-2d6b-4c31-905c-2d5ad7a7f8ab", "counter": 1, "stdout": "", "start_line": 0, "end_line": 0}
		`),

		Entry("extra attributes", `{"event": "playbook_on_start", "uuid": "cb93301e-5ff8-4f75-ade6-57d0ec2fc662", "counter": 0, "stdout": "", "start_line": 0, "end_line": 0, "event_data": {"playbook": "ping.yml", "playbook_uuid": "db6da5c7-37a6-479f-b18a-1db5af7f0932", "uuid": "db6da5c7-37a6-479f-b18a-1db5af7f0932"}}`),
	)
})
