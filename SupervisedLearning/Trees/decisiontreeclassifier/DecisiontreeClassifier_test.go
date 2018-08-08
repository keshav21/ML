package decisiontreeclassifier_test

import (
	"fmt"
	"io/ioutil"

	"code.cloudfoundry.org/lager/lagertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/Keshav21/ML/SupervisedLearning/Trees/decisiontreeclassifier"
)

var _ = Describe("DecisiontreeClassifier", func() {
	var (
		logger                   *lagertest.TestLogger
		decisionclassifierclient DecisionTreeClassifier
	)

	BeforeEach(func() {
		logger = lagertest.NewTestLogger("test")
	})

	Describe("DecisionTreeClassifier", func() {

		JustBeforeEach(func() {
			decisionclassifierclient = NewDecisionTreeClassifier(logger)
		})

		It("returns a new decisionclassifier client", func() {
			Expect(decisionclassifierclient).ToNot(BeNil())
		})

	})

	FDescribe("CreateClassAttribute", func() {
		var (
			descisionTreeAttribute DecisionTreeAttributes
		)

		BeforeEach(func() {
			decisionclassifierclient = NewDecisionTreeClassifier(logger)
		})

		JustBeforeEach(func() {
			attributeData, err := ioutil.ReadFile("./datasets/SampleattributeFile.txt")
			Expect(err).ToNot(HaveOccurred())
			descisionTreeAttribute = decisionclassifierclient.CreateClassAttribute(logger, attributeData)
		})

		It("creates the class Attribute", func() {
			fmt.Println(descisionTreeAttribute)
			Expect(descisionTreeAttribute).ToNot(BeNil())
		})
	})
})
