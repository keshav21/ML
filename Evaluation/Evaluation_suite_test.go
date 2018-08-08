package Evaluation_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEvaluation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Evaluation Suite")
}
