// Code generated from Feature.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Feature

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFeatureListener is a complete listener for a parse tree produced by FeatureParser.
type BaseFeatureListener struct{}

var _ FeatureListener = &BaseFeatureListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFeatureListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFeatureListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFeatureListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFeatureListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFeature is called when production feature is entered.
func (s *BaseFeatureListener) EnterFeature(ctx *FeatureContext) {}

// ExitFeature is called when production feature is exited.
func (s *BaseFeatureListener) ExitFeature(ctx *FeatureContext) {}

// EnterFeatureHeader is called when production featureHeader is entered.
func (s *BaseFeatureListener) EnterFeatureHeader(ctx *FeatureHeaderContext) {}

// ExitFeatureHeader is called when production featureHeader is exited.
func (s *BaseFeatureListener) ExitFeatureHeader(ctx *FeatureHeaderContext) {}

// EnterFeatureBody is called when production featureBody is entered.
func (s *BaseFeatureListener) EnterFeatureBody(ctx *FeatureBodyContext) {}

// ExitFeatureBody is called when production featureBody is exited.
func (s *BaseFeatureListener) ExitFeatureBody(ctx *FeatureBodyContext) {}

// EnterBackground is called when production background is entered.
func (s *BaseFeatureListener) EnterBackground(ctx *BackgroundContext) {}

// ExitBackground is called when production background is exited.
func (s *BaseFeatureListener) ExitBackground(ctx *BackgroundContext) {}

// EnterBlockBody is called when production blockBody is entered.
func (s *BaseFeatureListener) EnterBlockBody(ctx *BlockBodyContext) {}

// ExitBlockBody is called when production blockBody is exited.
func (s *BaseFeatureListener) ExitBlockBody(ctx *BlockBodyContext) {}

// EnterScenario is called when production scenario is entered.
func (s *BaseFeatureListener) EnterScenario(ctx *ScenarioContext) {}

// ExitScenario is called when production scenario is exited.
func (s *BaseFeatureListener) ExitScenario(ctx *ScenarioContext) {}

// EnterTags is called when production tags is entered.
func (s *BaseFeatureListener) EnterTags(ctx *TagsContext) {}

// ExitTags is called when production tags is exited.
func (s *BaseFeatureListener) ExitTags(ctx *TagsContext) {}

// EnterAnyText is called when production anyText is entered.
func (s *BaseFeatureListener) EnterAnyText(ctx *AnyTextContext) {}

// ExitAnyText is called when production anyText is exited.
func (s *BaseFeatureListener) ExitAnyText(ctx *AnyTextContext) {}

// EnterValue is called when production value is entered.
func (s *BaseFeatureListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseFeatureListener) ExitValue(ctx *ValueContext) {}

// EnterGiven is called when production given is entered.
func (s *BaseFeatureListener) EnterGiven(ctx *GivenContext) {}

// ExitGiven is called when production given is exited.
func (s *BaseFeatureListener) ExitGiven(ctx *GivenContext) {}

// EnterWhen is called when production when is entered.
func (s *BaseFeatureListener) EnterWhen(ctx *WhenContext) {}

// ExitWhen is called when production when is exited.
func (s *BaseFeatureListener) ExitWhen(ctx *WhenContext) {}

// EnterOr is called when production or is entered.
func (s *BaseFeatureListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production or is exited.
func (s *BaseFeatureListener) ExitOr(ctx *OrContext) {}

// EnterAnd is called when production and is entered.
func (s *BaseFeatureListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production and is exited.
func (s *BaseFeatureListener) ExitAnd(ctx *AndContext) {}

// EnterThen is called when production then is entered.
func (s *BaseFeatureListener) EnterThen(ctx *ThenContext) {}

// ExitThen is called when production then is exited.
func (s *BaseFeatureListener) ExitThen(ctx *ThenContext) {}

// EnterExample is called when production example is entered.
func (s *BaseFeatureListener) EnterExample(ctx *ExampleContext) {}

// ExitExample is called when production example is exited.
func (s *BaseFeatureListener) ExitExample(ctx *ExampleContext) {}

// EnterStep is called when production step is entered.
func (s *BaseFeatureListener) EnterStep(ctx *StepContext) {}

// ExitStep is called when production step is exited.
func (s *BaseFeatureListener) ExitStep(ctx *StepContext) {}

// EnterStepContent is called when production stepContent is entered.
func (s *BaseFeatureListener) EnterStepContent(ctx *StepContentContext) {}

// ExitStepContent is called when production stepContent is exited.
func (s *BaseFeatureListener) ExitStepContent(ctx *StepContentContext) {}

// EnterStepText is called when production stepText is entered.
func (s *BaseFeatureListener) EnterStepText(ctx *StepTextContext) {}

// ExitStepText is called when production stepText is exited.
func (s *BaseFeatureListener) ExitStepText(ctx *StepTextContext) {}

// EnterTable is called when production table is entered.
func (s *BaseFeatureListener) EnterTable(ctx *TableContext) {}

// ExitTable is called when production table is exited.
func (s *BaseFeatureListener) ExitTable(ctx *TableContext) {}

// EnterTableHeader is called when production tableHeader is entered.
func (s *BaseFeatureListener) EnterTableHeader(ctx *TableHeaderContext) {}

// ExitTableHeader is called when production tableHeader is exited.
func (s *BaseFeatureListener) ExitTableHeader(ctx *TableHeaderContext) {}

// EnterRow is called when production row is entered.
func (s *BaseFeatureListener) EnterRow(ctx *RowContext) {}

// ExitRow is called when production row is exited.
func (s *BaseFeatureListener) ExitRow(ctx *RowContext) {}

// EnterCell is called when production cell is entered.
func (s *BaseFeatureListener) EnterCell(ctx *CellContext) {}

// ExitCell is called when production cell is exited.
func (s *BaseFeatureListener) ExitCell(ctx *CellContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseFeatureListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseFeatureListener) ExitParameter(ctx *ParameterContext) {}

// EnterContentNoQuotes is called when production contentNoQuotes is entered.
func (s *BaseFeatureListener) EnterContentNoQuotes(ctx *ContentNoQuotesContext) {}

// ExitContentNoQuotes is called when production contentNoQuotes is exited.
func (s *BaseFeatureListener) ExitContentNoQuotes(ctx *ContentNoQuotesContext) {}

// EnterContentNoPipes is called when production contentNoPipes is entered.
func (s *BaseFeatureListener) EnterContentNoPipes(ctx *ContentNoPipesContext) {}

// ExitContentNoPipes is called when production contentNoPipes is exited.
func (s *BaseFeatureListener) ExitContentNoPipes(ctx *ContentNoPipesContext) {}

// EnterContent is called when production content is entered.
func (s *BaseFeatureListener) EnterContent(ctx *ContentContext) {}

// ExitContent is called when production content is exited.
func (s *BaseFeatureListener) ExitContent(ctx *ContentContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseFeatureListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseFeatureListener) ExitComment(ctx *CommentContext) {}

// EnterCommentText is called when production commentText is entered.
func (s *BaseFeatureListener) EnterCommentText(ctx *CommentTextContext) {}

// ExitCommentText is called when production commentText is exited.
func (s *BaseFeatureListener) ExitCommentText(ctx *CommentTextContext) {}

// EnterCommentValue is called when production commentValue is entered.
func (s *BaseFeatureListener) EnterCommentValue(ctx *CommentValueContext) {}

// ExitCommentValue is called when production commentValue is exited.
func (s *BaseFeatureListener) ExitCommentValue(ctx *CommentValueContext) {}
