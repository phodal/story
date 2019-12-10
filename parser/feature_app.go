package parser

import (
	. "story/languages/feature"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type CommentStruct struct {
	CommentsMap map[string]string
	CommentPosition
}

var commentStruct CommentStruct

func NewFeatureApp() *FeatureApp {
	commentStruct = *&CommentStruct{}
	return &FeatureApp{}
}

type FeatureApp struct {
}

func (j *FeatureApp) Start(path string) CommentStruct {
	context := (*FeatureApp)(nil).ProcessFile(path).Feature()
	listener := NewFeatureAppListener()

	antlr.NewParseTreeWalker().Walk(listener, context)

	commentStruct.CommentsMap = listener.getComments()

	position := listener.getCommentPosition()
	commentStruct.Start = position.Start
	commentStruct.End = position.End

	return commentStruct
}

func (j *FeatureApp) ProcessFile(path string) *FeatureParser {
	is, _ := antlr.NewFileStream(path)
	lexer := NewFeatureLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := NewFeatureParser(stream)
	return parser
}
