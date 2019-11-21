// Code generated from Feature.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Feature

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by FeatureParser.
type FeatureVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FeatureParser#compilationUnit.
	VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

	// Visit a parse tree produced by FeatureParser#feature_elements.
	VisitFeature_elements(ctx *Feature_elementsContext) interface{}

	// Visit a parse tree produced by FeatureParser#scenario.
	VisitScenario(ctx *ScenarioContext) interface{}

	// Visit a parse tree produced by FeatureParser#scenario_outline.
	VisitScenario_outline(ctx *Scenario_outlineContext) interface{}

	// Visit a parse tree produced by FeatureParser#steps.
	VisitSteps(ctx *StepsContext) interface{}

	// Visit a parse tree produced by FeatureParser#step.
	VisitStep(ctx *StepContext) interface{}

	// Visit a parse tree produced by FeatureParser#examples_sections.
	VisitExamples_sections(ctx *Examples_sectionsContext) interface{}

	// Visit a parse tree produced by FeatureParser#examples.
	VisitExamples(ctx *ExamplesContext) interface{}

	// Visit a parse tree produced by FeatureParser#multiline_arg.
	VisitMultiline_arg(ctx *Multiline_argContext) interface{}

	// Visit a parse tree produced by FeatureParser#table.
	VisitTable(ctx *TableContext) interface{}

	// Visit a parse tree produced by FeatureParser#table_row.
	VisitTable_row(ctx *Table_rowContext) interface{}

	// Visit a parse tree produced by FeatureParser#cell.
	VisitCell(ctx *CellContext) interface{}

	// Visit a parse tree produced by FeatureParser#tags.
	VisitTags(ctx *TagsContext) interface{}

	// Visit a parse tree produced by FeatureParser#tag.
	VisitTag(ctx *TagContext) interface{}

	// Visit a parse tree produced by FeatureParser#comment.
	VisitComment(ctx *CommentContext) interface{}

	// Visit a parse tree produced by FeatureParser#comment_line.
	VisitComment_line(ctx *Comment_lineContext) interface{}

	// Visit a parse tree produced by FeatureParser#line_to_eol.
	VisitLine_to_eol(ctx *Line_to_eolContext) interface{}

	// Visit a parse tree produced by FeatureParser#feature_keyword.
	VisitFeature_keyword(ctx *Feature_keywordContext) interface{}

	// Visit a parse tree produced by FeatureParser#scenario_keyword.
	VisitScenario_keyword(ctx *Scenario_keywordContext) interface{}

	// Visit a parse tree produced by FeatureParser#scenario_outline_keyword.
	VisitScenario_outline_keyword(ctx *Scenario_outline_keywordContext) interface{}

	// Visit a parse tree produced by FeatureParser#step_keyword.
	VisitStep_keyword(ctx *Step_keywordContext) interface{}

	// Visit a parse tree produced by FeatureParser#examples_keyword.
	VisitExamples_keyword(ctx *Examples_keywordContext) interface{}
}
