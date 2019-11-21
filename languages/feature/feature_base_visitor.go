// Code generated from Feature.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Feature

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseFeatureVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFeatureVisitor) VisitCompilationUnit(ctx *CompilationUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitFeature_elements(ctx *Feature_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitScenario(ctx *ScenarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitScenario_outline(ctx *Scenario_outlineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitSteps(ctx *StepsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitStep(ctx *StepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitExamples_sections(ctx *Examples_sectionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitExamples(ctx *ExamplesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitMultiline_arg(ctx *Multiline_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitTable(ctx *TableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitTable_row(ctx *Table_rowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitCell(ctx *CellContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitTags(ctx *TagsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitTag(ctx *TagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitComment_line(ctx *Comment_lineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitLine_to_eol(ctx *Line_to_eolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitFeature_keyword(ctx *Feature_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitScenario_keyword(ctx *Scenario_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitScenario_outline_keyword(ctx *Scenario_outline_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitStep_keyword(ctx *Step_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitExamples_keyword(ctx *Examples_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}
