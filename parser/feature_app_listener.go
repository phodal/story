package parser

import (
	. "../languages/feature"
)

var comments = make(map[string]string)

func NewFeatureAppListener() *FeatureAppListener {
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
