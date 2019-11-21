// Code generated from Feature.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Feature

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseFeatureVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFeatureVisitor) VisitFeature(ctx *FeatureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitFeatureHeader(ctx *FeatureHeaderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitFeatureBody(ctx *FeatureBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitBackground(ctx *BackgroundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitBlockBody(ctx *BlockBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitScenario(ctx *ScenarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitTags(ctx *TagsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitAnyText(ctx *AnyTextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitGiven(ctx *GivenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitWhen(ctx *WhenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitOr(ctx *OrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitAnd(ctx *AndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitThen(ctx *ThenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitStep(ctx *StepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitStepContent(ctx *StepContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitStepText(ctx *StepTextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitRow(ctx *RowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitCell(ctx *CellContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitContentNoQuotes(ctx *ContentNoQuotesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitContentNoPipes(ctx *ContentNoPipesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFeatureVisitor) VisitContent(ctx *ContentContext) interface{} {
	return v.VisitChildren(ctx)
}
