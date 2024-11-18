package run

import (
	"github.com/BFMVAUS/d2go/pkg/data"
	"github.com/BFMVAUS/d2go/pkg/data/area"
	"github.com/BFMVAUS/koolov2/internal/action"
	"github.com/BFMVAUS/koolov2/internal/config"
	"github.com/BFMVAUS/koolov2/internal/context"
)

type ArachnidLair struct {
	ctx *context.Status
}

func NewArachnidLair() *ArachnidLair {
	return &ArachnidLair{
		ctx: context.Get(),
	}
}

func (a ArachnidLair) Name() string {
	return string(config.ArachnidLairRun)
}

func (a ArachnidLair) Run() error {
	err := action.WayPoint(area.SpiderForest)
	if err != nil {
		return err
	}

	err = action.MoveToArea(area.SpiderCave)
	if err != nil {
		return err
	}

	action.OpenTPIfLeader()

	// Clear ArachnidLair
	return action.ClearCurrentLevel(true, data.MonsterAnyFilter())
}
