// Code generated from Feature.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Feature

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by FeatureParser.
type FeatureVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FeatureParser#feature.
	VisitFeature(ctx *FeatureContext) interface{}

	// Visit a parse tree produced by FeatureParser#featureHeader.
	VisitFeatureHeader(ctx *FeatureHeaderContext) interface{}

	// Visit a parse tree produced by FeatureParser#featureBody.
	VisitFeatureBody(ctx *FeatureBodyContext) interface{}

	// Visit a parse tree produced by FeatureParser#background.
	VisitBackground(ctx *BackgroundContext) interface{}

	// Visit a parse tree produced by FeatureParser#blockBody.
	VisitBlockBody(ctx *BlockBodyContext) interface{}

	// Visit a parse tree produced by FeatureParser#scenario.
	VisitScenario(ctx *ScenarioContext) interface{}

	// Visit a parse tree produced by FeatureParser#tags.
	VisitTags(ctx *TagsContext) interface{}

	// Visit a parse tree produced by FeatureParser#anyText.
	VisitAnyText(ctx *AnyTextContext) interface{}

	// Visit a parse tree produced by FeatureParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by FeatureParser#given.
	VisitGiven(ctx *GivenContext) interface{}

	// Visit a parse tree produced by FeatureParser#when.
	VisitWhen(ctx *WhenContext) interface{}

	// Visit a parse tree produced by FeatureParser#or.
	VisitOr(ctx *OrContext) interface{}

	// Visit a parse tree produced by FeatureParser#and.
	VisitAnd(ctx *AndContext) interface{}

	// Visit a parse tree produced by FeatureParser#then.
	VisitThen(ctx *ThenContext) interface{}

	// Visit a parse tree produced by FeatureParser#example.
	VisitExample(ctx *ExampleContext) interface{}

	// Visit a parse tree produced by FeatureParser#step.
	VisitStep(ctx *StepContext) interface{}

	// Visit a parse tree produced by FeatureParser#stepContent.
	VisitStepContent(ctx *StepContentContext) interface{}

	// Visit a parse tree produced by FeatureParser#stepText.
	VisitStepText(ctx *StepTextContext) interface{}

	// Visit a parse tree produced by FeatureParser#table.
	VisitTable(ctx *TableContext) interface{}

	// Visit a parse tree produced by FeatureParser#tableHeader.
	VisitTableHeader(ctx *TableHeaderContext) interface{}

	// Visit a parse tree produced by FeatureParser#row.
	VisitRow(ctx *RowContext) interface{}

	// Visit a parse tree produced by FeatureParser#cell.
	VisitCell(ctx *CellContext) interface{}

	// Visit a parse tree produced by FeatureParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by FeatureParser#contentNoQuotes.
	VisitContentNoQuotes(ctx *ContentNoQuotesContext) interface{}

	// Visit a parse tree produced by FeatureParser#contentNoPipes.
	VisitContentNoPipes(ctx *ContentNoPipesContext) interface{}

	// Visit a parse tree produced by FeatureParser#content.
	VisitContent(ctx *ContentContext) interface{}

	// Visit a parse tree produced by FeatureParser#comment.
	VisitComment(ctx *CommentContext) interface{}

	// Visit a parse tree produced by FeatureParser#commentText.
	VisitCommentText(ctx *CommentTextContext) interface{}

	// Visit a parse tree produced by FeatureParser#commentValue.
	VisitCommentValue(ctx *CommentValueContext) interface{}
}
