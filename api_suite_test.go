package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestWeb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Suite")
}
