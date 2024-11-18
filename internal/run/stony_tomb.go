package run

import (
	"github.com/BFMVAUS/d2go/pkg/data"
	"github.com/BFMVAUS/d2go/pkg/data/area"
	"github.com/BFMVAUS/koolov2/internal/action"
	"github.com/BFMVAUS/koolov2/internal/config"
	"github.com/BFMVAUS/koolov2/internal/context"
)

type StonyTomb struct {
	ctx *context.Status
}

func NewStonyTomb() *StonyTomb {
	return &StonyTomb{
		ctx: context.Get(),
	}
}

func (s StonyTomb) Name() string {
	return string(config.StonyTombRun)
}

func (s StonyTomb) Run() error {

	// Setup default filter
	monsterFilter := data.MonsterAnyFilter()

	// Update filter if we selected to clear only elites
	if s.ctx.CharacterCfg.Game.StonyTomb.FocusOnElitePacks {
		monsterFilter = data.MonsterEliteFilter()
	}

	// Use the waypoint
	err := action.WayPoint(area.DryHills)
	if err != nil {
		return err
	}

	// Move to the correct area
	if err = action.MoveToArea(area.RockyWaste); err != nil {
		return err
	}

	// Move to the correct area
	if err = action.MoveToArea(area.StonyTombLevel1); err != nil {
		return err
	}

	// Open a TP If we're the leader
	action.OpenTPIfLeader()

	// Clear the area
	if err = action.ClearCurrentLevel(s.ctx.CharacterCfg.Game.StonyTomb.OpenChests, monsterFilter); err != nil {
		return err
	}

	// Move to lvl2
	if err = action.MoveToArea(area.StonyTombLevel2); err != nil {
		return err
	}

	// Clear the area
	return action.ClearCurrentLevel(s.ctx.CharacterCfg.Game.StonyTomb.OpenChests, monsterFilter)
}
