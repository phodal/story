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

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseFeatureListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseFeatureListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterFeature_elements is called when production feature_elements is entered.
func (s *BaseFeatureListener) EnterFeature_elements(ctx *Feature_elementsContext) {}

// ExitFeature_elements is called when production feature_elements is exited.
func (s *BaseFeatureListener) ExitFeature_elements(ctx *Feature_elementsContext) {}

// EnterScenario is called when production scenario is entered.
func (s *BaseFeatureListener) EnterScenario(ctx *ScenarioContext) {}

// ExitScenario is called when production scenario is exited.
func (s *BaseFeatureListener) ExitScenario(ctx *ScenarioContext) {}

// EnterScenario_outline is called when production scenario_outline is entered.
func (s *BaseFeatureListener) EnterScenario_outline(ctx *Scenario_outlineContext) {}

// ExitScenario_outline is called when production scenario_outline is exited.
func (s *BaseFeatureListener) ExitScenario_outline(ctx *Scenario_outlineContext) {}

// EnterSteps is called when production steps is entered.
func (s *BaseFeatureListener) EnterSteps(ctx *StepsContext) {}

// ExitSteps is called when production steps is exited.
func (s *BaseFeatureListener) ExitSteps(ctx *StepsContext) {}

// EnterStep is called when production step is entered.
func (s *BaseFeatureListener) EnterStep(ctx *StepContext) {}

// ExitStep is called when production step is exited.
func (s *BaseFeatureListener) ExitStep(ctx *StepContext) {}

// EnterExamples_sections is called when production examples_sections is entered.
func (s *BaseFeatureListener) EnterExamples_sections(ctx *Examples_sectionsContext) {}

// ExitExamples_sections is called when production examples_sections is exited.
func (s *BaseFeatureListener) ExitExamples_sections(ctx *Examples_sectionsContext) {}

// EnterExamples is called when production examples is entered.
func (s *BaseFeatureListener) EnterExamples(ctx *ExamplesContext) {}

// ExitExamples is called when production examples is exited.
func (s *BaseFeatureListener) ExitExamples(ctx *ExamplesContext) {}

// EnterMultiline_arg is called when production multiline_arg is entered.
func (s *BaseFeatureListener) EnterMultiline_arg(ctx *Multiline_argContext) {}

// ExitMultiline_arg is called when production multiline_arg is exited.
func (s *BaseFeatureListener) ExitMultiline_arg(ctx *Multiline_argContext) {}

// EnterTable is called when production table is entered.
func (s *BaseFeatureListener) EnterTable(ctx *TableContext) {}

// ExitTable is called when production table is exited.
func (s *BaseFeatureListener) ExitTable(ctx *TableContext) {}

// EnterTable_row is called when production table_row is entered.
func (s *BaseFeatureListener) EnterTable_row(ctx *Table_rowContext) {}

// ExitTable_row is called when production table_row is exited.
func (s *BaseFeatureListener) ExitTable_row(ctx *Table_rowContext) {}

// EnterCell is called when production cell is entered.
func (s *BaseFeatureListener) EnterCell(ctx *CellContext) {}

// ExitCell is called when production cell is exited.
func (s *BaseFeatureListener) ExitCell(ctx *CellContext) {}

// EnterTags is called when production tags is entered.
func (s *BaseFeatureListener) EnterTags(ctx *TagsContext) {}

// ExitTags is called when production tags is exited.
func (s *BaseFeatureListener) ExitTags(ctx *TagsContext) {}

// EnterTag is called when production tag is entered.
func (s *BaseFeatureListener) EnterTag(ctx *TagContext) {}

// ExitTag is called when production tag is exited.
func (s *BaseFeatureListener) ExitTag(ctx *TagContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseFeatureListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseFeatureListener) ExitComment(ctx *CommentContext) {}

// EnterComment_line is called when production comment_line is entered.
func (s *BaseFeatureListener) EnterComment_line(ctx *Comment_lineContext) {}

// ExitComment_line is called when production comment_line is exited.
func (s *BaseFeatureListener) ExitComment_line(ctx *Comment_lineContext) {}

// EnterLine_to_eol is called when production line_to_eol is entered.
func (s *BaseFeatureListener) EnterLine_to_eol(ctx *Line_to_eolContext) {}

// ExitLine_to_eol is called when production line_to_eol is exited.
func (s *BaseFeatureListener) ExitLine_to_eol(ctx *Line_to_eolContext) {}

// EnterFeature_keyword is called when production feature_keyword is entered.
func (s *BaseFeatureListener) EnterFeature_keyword(ctx *Feature_keywordContext) {}

// ExitFeature_keyword is called when production feature_keyword is exited.
func (s *BaseFeatureListener) ExitFeature_keyword(ctx *Feature_keywordContext) {}

// EnterScenario_keyword is called when production scenario_keyword is entered.
func (s *BaseFeatureListener) EnterScenario_keyword(ctx *Scenario_keywordContext) {}

// ExitScenario_keyword is called when production scenario_keyword is exited.
func (s *BaseFeatureListener) ExitScenario_keyword(ctx *Scenario_keywordContext) {}

// EnterScenario_outline_keyword is called when production scenario_outline_keyword is entered.
func (s *BaseFeatureListener) EnterScenario_outline_keyword(ctx *Scenario_outline_keywordContext) {}

// ExitScenario_outline_keyword is called when production scenario_outline_keyword is exited.
func (s *BaseFeatureListener) ExitScenario_outline_keyword(ctx *Scenario_outline_keywordContext) {}

// EnterStep_keyword is called when production step_keyword is entered.
func (s *BaseFeatureListener) EnterStep_keyword(ctx *Step_keywordContext) {}

// ExitStep_keyword is called when production step_keyword is exited.
func (s *BaseFeatureListener) ExitStep_keyword(ctx *Step_keywordContext) {}

// EnterExamples_keyword is called when production examples_keyword is entered.
func (s *BaseFeatureListener) EnterExamples_keyword(ctx *Examples_keywordContext) {}

// ExitExamples_keyword is called when production examples_keyword is exited.
func (s *BaseFeatureListener) ExitExamples_keyword(ctx *Examples_keywordContext) {}
