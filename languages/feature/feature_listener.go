// Code generated from Feature.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Feature

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FeatureListener is a complete listener for a parse tree produced by FeatureParser.
type FeatureListener interface {
	antlr.ParseTreeListener

	// EnterFeature is called when entering the feature production.
	EnterFeature(c *FeatureContext)

	// EnterFeatureHeader is called when entering the featureHeader production.
	EnterFeatureHeader(c *FeatureHeaderContext)

	// EnterFeatureBody is called when entering the featureBody production.
	EnterFeatureBody(c *FeatureBodyContext)

	// EnterBackground is called when entering the background production.
	EnterBackground(c *BackgroundContext)

	// EnterBlockBody is called when entering the blockBody production.
	EnterBlockBody(c *BlockBodyContext)

	// EnterScenario is called when entering the scenario production.
	EnterScenario(c *ScenarioContext)

	// EnterTags is called when entering the tags production.
	EnterTags(c *TagsContext)

	// EnterAnyText is called when entering the anyText production.
	EnterAnyText(c *AnyTextContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterGiven is called when entering the given production.
	EnterGiven(c *GivenContext)

	// EnterWhen is called when entering the when production.
	EnterWhen(c *WhenContext)

	// EnterOr is called when entering the or production.
	EnterOr(c *OrContext)

	// EnterAnd is called when entering the and production.
	EnterAnd(c *AndContext)

	// EnterThen is called when entering the then production.
	EnterThen(c *ThenContext)

	// EnterExample is called when entering the example production.
	EnterExample(c *ExampleContext)

	// EnterStep is called when entering the step production.
	EnterStep(c *StepContext)

	// EnterStepContent is called when entering the stepContent production.
	EnterStepContent(c *StepContentContext)

	// EnterStepText is called when entering the stepText production.
	EnterStepText(c *StepTextContext)

	// EnterTable is called when entering the table production.
	EnterTable(c *TableContext)

	// EnterTableHeader is called when entering the tableHeader production.
	EnterTableHeader(c *TableHeaderContext)

	// EnterRow is called when entering the row production.
	EnterRow(c *RowContext)

	// EnterCell is called when entering the cell production.
	EnterCell(c *CellContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterContentNoQuotes is called when entering the contentNoQuotes production.
	EnterContentNoQuotes(c *ContentNoQuotesContext)

	// EnterContentNoPipes is called when entering the contentNoPipes production.
	EnterContentNoPipes(c *ContentNoPipesContext)

	// EnterContent is called when entering the content production.
	EnterContent(c *ContentContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterCommentText is called when entering the commentText production.
	EnterCommentText(c *CommentTextContext)

	// EnterCommentValue is called when entering the commentValue production.
	EnterCommentValue(c *CommentValueContext)

	// ExitFeature is called when exiting the feature production.
	ExitFeature(c *FeatureContext)

	// ExitFeatureHeader is called when exiting the featureHeader production.
	ExitFeatureHeader(c *FeatureHeaderContext)

	// ExitFeatureBody is called when exiting the featureBody production.
	ExitFeatureBody(c *FeatureBodyContext)

	// ExitBackground is called when exiting the background production.
	ExitBackground(c *BackgroundContext)

	// ExitBlockBody is called when exiting the blockBody production.
	ExitBlockBody(c *BlockBodyContext)

	// ExitScenario is called when exiting the scenario production.
	ExitScenario(c *ScenarioContext)

	// ExitTags is called when exiting the tags production.
	ExitTags(c *TagsContext)

	// ExitAnyText is called when exiting the anyText production.
	ExitAnyText(c *AnyTextContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitGiven is called when exiting the given production.
	ExitGiven(c *GivenContext)

	// ExitWhen is called when exiting the when production.
	ExitWhen(c *WhenContext)

	// ExitOr is called when exiting the or production.
	ExitOr(c *OrContext)

	// ExitAnd is called when exiting the and production.
	ExitAnd(c *AndContext)

	// ExitThen is called when exiting the then production.
	ExitThen(c *ThenContext)

	// ExitExample is called when exiting the example production.
	ExitExample(c *ExampleContext)

	// ExitStep is called when exiting the step production.
	ExitStep(c *StepContext)

	// ExitStepContent is called when exiting the stepContent production.
	ExitStepContent(c *StepContentContext)

	// ExitStepText is called when exiting the stepText production.
	ExitStepText(c *StepTextContext)

	// ExitTable is called when exiting the table production.
	ExitTable(c *TableContext)

	// ExitTableHeader is called when exiting the tableHeader production.
	ExitTableHeader(c *TableHeaderContext)

	// ExitRow is called when exiting the row production.
	ExitRow(c *RowContext)

	// ExitCell is called when exiting the cell production.
	ExitCell(c *CellContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitContentNoQuotes is called when exiting the contentNoQuotes production.
	ExitContentNoQuotes(c *ContentNoQuotesContext)

	// ExitContentNoPipes is called when exiting the contentNoPipes production.
	ExitContentNoPipes(c *ContentNoPipesContext)

	// ExitContent is called when exiting the content production.
	ExitContent(c *ContentContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitCommentText is called when exiting the commentText production.
	ExitCommentText(c *CommentTextContext)

	// ExitCommentValue is called when exiting the commentValue production.
	ExitCommentValue(c *CommentValueContext)
}
