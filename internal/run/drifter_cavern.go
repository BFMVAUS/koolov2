package run

import (
	"github.com/BFMVAUS/d2go/pkg/data"
	"github.com/BFMVAUS/d2go/pkg/data/area"
	"github.com/BFMVAUS/koolov2/internal/action"
	"github.com/BFMVAUS/koolov2/internal/config"
	"github.com/BFMVAUS/koolov2/internal/context"
)

type DrifterCavern struct {
	ctx *context.Status
}

func NewDriverCavern() *DrifterCavern {
	return &DrifterCavern{
		ctx: context.Get(),
	}
}

func (s DrifterCavern) Name() string {
	return string(config.DrifterCavernRun)
}

func (s DrifterCavern) Run() error {
	// Define a default monster filter
	monsterFilter := data.MonsterAnyFilter()

	// Update filter if we selected to clear only elites
	if s.ctx.CharacterCfg.Game.DrifterCavern.FocusOnElitePacks {
		monsterFilter = data.MonsterEliteFilter()
	}

	// Use the waypoint
	err := action.WayPoint(area.GlacialTrail)
	if err != nil {
		return err
	}

	// Move to the correct area
	if err = action.MoveToArea(area.DrifterCavern); err != nil {
		return err
	}

	// Clear the area
	return action.ClearCurrentLevel(s.ctx.CharacterCfg.Game.DrifterCavern.OpenChests, monsterFilter)
}
