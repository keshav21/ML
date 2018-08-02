package tree

import "code.cloudfoundry.org/lager"

//DecisionTreeClassifier is interface for DecisionTreeClassifier
type DecisionTreeClassifier interface {
}

//NewDecisionTreeClassifier is a function to initialize a Classifier
func NewDecisionTreeClassifier(logger lager.Logger) DecisionTreeClassifier {
	logger = logger.Session("new-cf-client")
	logger.Info("started")
	defer logger.Info("completed")

	return &decisionTreeClassifierImpl{
		logger.Session("cf-client"),
	}
}

type decisionTreeClassifierImpl struct {
	logger lager.Logger
}
