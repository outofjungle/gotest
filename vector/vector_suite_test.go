package vector_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestVector(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Vector Suite")
}
