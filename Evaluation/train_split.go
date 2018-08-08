package Evaluation

import (
	"strings"

	"code.cloudfoundry.org/lager"
)

// Xtrain ..
type Xtrain map[int][]string

// Ytrain ..
type Ytrain map[int]string

// Xtest ..
type Xtest map[int][]string

// Ytest ..
type Ytest map[int]string

// Evaluation ...
type Evaluation interface {
	TrainTestSplit(logger lager.Logger, attributeFile []byte) (Xtrain, Ytrain, Xtest, Ytest)
}

type evaluationImpl struct {
	logger lager.Logger
}

// NewEvaluation implements the evaluation interface
func NewEvaluation(logger lager.Logger) Evaluation {
	logger = logger.Session("new-evaluation-client")
	logger.Info("started")
	defer logger.Info("completed")

	return &evaluationImpl{
		logger.Session("evaluation-client"),
	}
}

func (evaluation *evaluationImpl) TrainTestSplit(logger lager.Logger, attributeFile []byte) (Xtrain, Ytrain, Xtest, Ytest) {
	var (
		X map[int][]string
		Y map[int]string
	)
	X = make(map[int][]string)
	Y = make(map[int]string)
	fileData := string(attributeFile)
	filelines := strings.Split(fileData, "\n")
	for index, line := range filelines {
		data := strings.Split(line, ",")
		len := len(data) - 1
		Y[index] = data[len]
		X[index] = data[:len]
	}
	return Xtrain{}, Ytrain{}, Xtest{}, Ytest{}
}
