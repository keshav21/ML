package Evaluation_test

import (
	"io/ioutil"

	"code.cloudfoundry.org/lager/lagertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.ibm.com/Keshav21/ML/Evaluation"
)

var _ = Describe("TrainSplit", func() {
	var (
		logger           *lagertest.TestLogger
		evaluationClient Evaluation
	)

	BeforeEach(func() {
		logger = lagertest.NewTestLogger("test")
	})

	Describe("NewEvaluationClassifier", func() {

		JustBeforeEach(func() {
			evaluationClient = NewEvaluation(logger)
		})

		It("returns a new evaluation client", func() {
			Expect(evaluationClient).ToNot(BeNil())
		})

	})

	Describe("TrainTestSplit", func() {

		BeforeEach(func() {
			evaluationClient = NewEvaluation(logger)
		})
		JustBeforeEach(func() {
			attributeData, err := ioutil.ReadFile("./SampledataFile.csv")
			Expect(err).ToNot(HaveOccurred())
			_, _, _, _ = evaluationClient.TrainTestSplit(logger, attributeData)
		})
		It("returns a new evaluation client", func() {
			Expect(evaluationClient).ToNot(BeNil())
		})

	})
})
