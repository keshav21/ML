package tree_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDecisionTreeClassifier(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DecisionTreeClassifier Suite")
}
