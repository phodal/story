package parser

import (
	. "../languages/feature"
	"fmt"
)

func NewFeatureAppListener() *FeatureAppListener {
	return &FeatureAppListener{}
}

type FeatureAppListener struct {
	BaseFeatureListener
}

func (s *FeatureAppListener) EnterScenario(ctx *ScenarioContext) {
	fmt.Println(ctx.GetText())
}

func (s *FeatureAppListener) EnterAnd(ctx *AndContext) {
	fmt.Println(ctx.GetText())
}
