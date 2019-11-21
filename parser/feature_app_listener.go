package parser

import (
	. "../languages/feature"
	"fmt"
	"reflect"
)

var comments []string

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
	//text := ctx.CommentText().GetText()
	//comments = append(comments, text)
	//fmt.Println(ctx.CommentText().GetChild(0))
	//fmt.Println(ctx.CommentText().GetChild(2))
	fmt.Println(ctx.CommentText().GetChildCount())
	fmt.Println(reflect.TypeOf(ctx.CommentText().GetChild(0)).String())
}

func (s *FeatureAppListener) EnterStep(ctx *StepContext) {

}

func (s *FeatureAppListener) EnterRow(ctx *RowContext) {

}

func (s *FeatureAppListener) EnterCell(ctx *CellContext) {
	//fmt.Println(ctx.ContentNoPipes().GetText())
}

func (s *FeatureAppListener) EnterTableHeader(ctx *TableHeaderContext) {
	//cells := ctx.AllCell()
	//for _, cell := range cells {
	//	content := cell.(*CellContext).ContentNoPipes()
	//	fmt.Println(content.GetText())
	//}
}

func (s *FeatureAppListener) getComments() []string {
	return comments
}