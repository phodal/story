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
