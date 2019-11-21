package parser

import (
	. "github.com/phodal/story/languages/feature"
)

var hasFirstComment = false

type CommentPosition struct {
	Start int
	End   int
}
var commentPos CommentPosition

var comments = make(map[string]string)

func NewFeatureAppListener() *FeatureAppListener {
	commentPos = *&CommentPosition{}
	return &FeatureAppListener{}
}

type FeatureAppListener struct {
	BaseFeatureListener
}

func (s *FeatureAppListener) EnterScenario(ctx *ScenarioContext) {

}

func (s *FeatureAppListener) EnterAnd(ctx *AndContext) {

}

func (s *FeatureAppListener) EnterComment(ctx *CommentContext) {
	startLine := ctx.GetStart().GetLine()

	if hasFirstComment == false {
		hasFirstComment = true
		commentPos.Start = startLine
	}
	commentPos.End = startLine

	commentTextCtx := ctx.CommentText().(*CommentTextContext)
	identifier := commentTextCtx.IDENTIFIER()
	commentValue := commentTextCtx.CommentValue()
	if identifier != nil && commentValue != nil {
		comments[identifier.GetText()] = commentValue.GetText()
	}
}

func (s *FeatureAppListener) EnterStep(ctx *StepContext) {

}

func (s *FeatureAppListener) EnterRow(ctx *RowContext) {

}

func (s *FeatureAppListener) EnterCell(ctx *CellContext) {

}

func (s *FeatureAppListener) EnterTableHeader(ctx *TableHeaderContext) {

}

func (s *FeatureAppListener) getComments() map[string]string {
	return comments
}
func (s *FeatureAppListener) getCommentPosition() CommentPosition {
	return commentPos
}
