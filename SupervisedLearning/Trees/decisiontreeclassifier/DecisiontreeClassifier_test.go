package tree_test

import (
	"code.cloudfoundry.org/lager/lagertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	decisiontreeclassifier "github.ibm.com/Keshav21/ML/SupervisedLearning/Trees/decisiontreeclassifier"
)

var _ = Describe("DecisiontreeClassifier", func() {

	FDescribe("DecisionTreeClassifier", func() {
		var (
			logger                   *lagertest.TestLogger
			decisionclassifierclient decisiontreeclassifier.DecisionTreeClassifier
		)
		BeforeEach(func() {
			logger = lagertest.NewTestLogger("test")
		})

		JustBeforeEach(func() {
			decisionclassifierclient = decisiontreeclassifier.NewDecisionTreeClassifier(logger)
		})

		It("returns a new decisionclassifier client", func() {
			Expect(decisionclassifierclient).ToNot(BeNil())
		})

	})
})
