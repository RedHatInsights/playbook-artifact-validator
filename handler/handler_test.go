package handler

import (
	"playbook-artifact-validator/ingress"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handler", func() {
	Describe("File size", func() {
		It("Rejects an upload over max size", func() {
			req := &ingress.Request{
				Size: 128 * 1024,
			}

			res := onRequest(req)
			Expect(res.Validation).To(Equal("failure"))
		})
	})
})
