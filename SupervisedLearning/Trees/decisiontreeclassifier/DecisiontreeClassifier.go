package decisiontreeclassifier

import (
	"strings"

	"code.cloudfoundry.org/lager"
)

//DecisionTreeClassifier is interface for DecisionTreeClassifier
type DecisionTreeClassifier interface {
	CreateClassAttribute(logger lager.Logger, attributeFile []byte) DecisionTreeAttributes
}

//NewDecisionTreeClassifier is a function to initialize a Classifier
func NewDecisionTreeClassifier(logger lager.Logger) DecisionTreeClassifier {
	logger = logger.Session("new-decison-tree-client")
	logger.Info("started")
	defer logger.Info("completed")

	return &decisionTreeClassifierImpl{
		logger.Session("decison-tree-client"),
	}
}

// AttributeNode is the node in the Attribute
type AttributeNode struct {
	nodeName  string
	nodeCount int
}

// Attribute is the ClassAttribute
type Attribute struct {
	attributeName     string
	attributeNodeList []AttributeNode
}

//DecisionTreeAttributes store the attributes for the decision tree
type DecisionTreeAttributes struct {
	attributeList []Attribute
}
type decisionTreeClassifierImpl struct {
	logger lager.Logger
}

func (decisiontreeclassifier *decisionTreeClassifierImpl) CreateClassAttribute(logger lager.Logger, attributeFile []byte) DecisionTreeAttributes {
	logger = logger.Session("create-class-attribute")
	logger.Info("started")
	defer logger.Info("completed")
	var (
		decisiontreeAttribute DecisionTreeAttributes
		attributeList         []Attribute
	)
	attributeList = []Attribute{}
	fileData := string(attributeFile)
	filelines := strings.Split(fileData, "\n")
	for _, line := range filelines {
		data := strings.Split(line, ",")
		attribute := decisiontreeclassifier.createAttributeNodes(logger, data)
		attributeList = append(attributeList, attribute)
	}
	decisiontreeAttribute.attributeList = attributeList
	return decisiontreeAttribute
}

func (decisiontreeclassifier *decisionTreeClassifierImpl) createAttributeNodes(logger lager.Logger, attributeData []string) Attribute {

	logger = logger.Session("create-attribute-node")
	logger.Info("started")
	defer logger.Info("completed")
	var (
		attribute     Attribute
		attributeNode AttributeNode
		attributeList []AttributeNode
	)
	attribute = Attribute{}
	attributeNode = AttributeNode{}
	attribute.attributeName = attributeData[0]
	attributeList = []AttributeNode{}
	for _, node := range attributeData[1:] {
		attributeNode.nodeName = node
		attributeList = append(attributeList, attributeNode)
		attributeNode = AttributeNode{}
	}
	attribute.attributeNodeList = attributeList
	return attribute
}

//https://machinelearningmastery.com/implement-decision-tree-algorithm-scratch-python/
// https://github.com/sjwhitworth/golearn/tree/master/trees
// func (decisiontreeclassifier *decisionTreeClassifierImpl) CrossValidationSplit(logger lager.Logger, attributeData []byte) *list.List {
// 	return list{}
// }
