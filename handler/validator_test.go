package handler

import (
	"playbook-artifact-validator/ingress"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validator", func() {
	It("Rejects every payload", func() {
		request := &ingress.Request{}
		response := validateArtifacts(request)
		Expect(response.Validation).To(Equal("failure"))
	})
})
