package handler

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func RunTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Upload Suite")
}
