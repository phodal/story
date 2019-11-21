package parser

import (
	. "../languages/feature"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func NewFeatureApp() *FeatureApp {
	return &FeatureApp{}
}

type FeatureApp struct {

}

func (j *FeatureApp) Start(path string) {
	context := (*FeatureApp)(nil).ProcessFile(path).Feature()
	listener := NewFeatureAppListener()

	antlr.NewParseTreeWalker().Walk(listener, context)

	//getComments := listener.getComments()
}

func (j *FeatureApp) ProcessFile(path string) *FeatureParser {
	is, _ := antlr.NewFileStream(path)
	lexer := NewFeatureLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := NewFeatureParser(stream)
	return parser
}
