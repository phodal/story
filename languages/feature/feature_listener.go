// Code generated from Feature.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Feature

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FeatureListener is a complete listener for a parse tree produced by FeatureParser.
type FeatureListener interface {
	antlr.ParseTreeListener

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterFeature_elements is called when entering the feature_elements production.
	EnterFeature_elements(c *Feature_elementsContext)

	// EnterScenario is called when entering the scenario production.
	EnterScenario(c *ScenarioContext)

	// EnterScenario_outline is called when entering the scenario_outline production.
	EnterScenario_outline(c *Scenario_outlineContext)

	// EnterSteps is called when entering the steps production.
	EnterSteps(c *StepsContext)

	// EnterStep is called when entering the step production.
	EnterStep(c *StepContext)

	// EnterExamples_sections is called when entering the examples_sections production.
	EnterExamples_sections(c *Examples_sectionsContext)

	// EnterExamples is called when entering the examples production.
	EnterExamples(c *ExamplesContext)

	// EnterMultiline_arg is called when entering the multiline_arg production.
	EnterMultiline_arg(c *Multiline_argContext)

	// EnterTable is called when entering the table production.
	EnterTable(c *TableContext)

	// EnterTable_row is called when entering the table_row production.
	EnterTable_row(c *Table_rowContext)

	// EnterCell is called when entering the cell production.
	EnterCell(c *CellContext)

	// EnterTags is called when entering the tags production.
	EnterTags(c *TagsContext)

	// EnterTag is called when entering the tag production.
	EnterTag(c *TagContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterComment_line is called when entering the comment_line production.
	EnterComment_line(c *Comment_lineContext)

	// EnterLine_to_eol is called when entering the line_to_eol production.
	EnterLine_to_eol(c *Line_to_eolContext)

	// EnterFeature_keyword is called when entering the feature_keyword production.
	EnterFeature_keyword(c *Feature_keywordContext)

	// EnterScenario_keyword is called when entering the scenario_keyword production.
	EnterScenario_keyword(c *Scenario_keywordContext)

	// EnterScenario_outline_keyword is called when entering the scenario_outline_keyword production.
	EnterScenario_outline_keyword(c *Scenario_outline_keywordContext)

	// EnterStep_keyword is called when entering the step_keyword production.
	EnterStep_keyword(c *Step_keywordContext)

	// EnterExamples_keyword is called when entering the examples_keyword production.
	EnterExamples_keyword(c *Examples_keywordContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitFeature_elements is called when exiting the feature_elements production.
	ExitFeature_elements(c *Feature_elementsContext)

	// ExitScenario is called when exiting the scenario production.
	ExitScenario(c *ScenarioContext)

	// ExitScenario_outline is called when exiting the scenario_outline production.
	ExitScenario_outline(c *Scenario_outlineContext)

	// ExitSteps is called when exiting the steps production.
	ExitSteps(c *StepsContext)

	// ExitStep is called when exiting the step production.
	ExitStep(c *StepContext)

	// ExitExamples_sections is called when exiting the examples_sections production.
	ExitExamples_sections(c *Examples_sectionsContext)

	// ExitExamples is called when exiting the examples production.
	ExitExamples(c *ExamplesContext)

	// ExitMultiline_arg is called when exiting the multiline_arg production.
	ExitMultiline_arg(c *Multiline_argContext)

	// ExitTable is called when exiting the table production.
	ExitTable(c *TableContext)

	// ExitTable_row is called when exiting the table_row production.
	ExitTable_row(c *Table_rowContext)

	// ExitCell is called when exiting the cell production.
	ExitCell(c *CellContext)

	// ExitTags is called when exiting the tags production.
	ExitTags(c *TagsContext)

	// ExitTag is called when exiting the tag production.
	ExitTag(c *TagContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitComment_line is called when exiting the comment_line production.
	ExitComment_line(c *Comment_lineContext)

	// ExitLine_to_eol is called when exiting the line_to_eol production.
	ExitLine_to_eol(c *Line_to_eolContext)

	// ExitFeature_keyword is called when exiting the feature_keyword production.
	ExitFeature_keyword(c *Feature_keywordContext)

	// ExitScenario_keyword is called when exiting the scenario_keyword production.
	ExitScenario_keyword(c *Scenario_keywordContext)

	// ExitScenario_outline_keyword is called when exiting the scenario_outline_keyword production.
	ExitScenario_outline_keyword(c *Scenario_outline_keywordContext)

	// ExitStep_keyword is called when exiting the step_keyword production.
	ExitStep_keyword(c *Step_keywordContext)

	// ExitExamples_keyword is called when exiting the examples_keyword production.
	ExitExamples_keyword(c *Examples_keywordContext)
}
