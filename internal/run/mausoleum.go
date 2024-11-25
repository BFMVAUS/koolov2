package run

import (
	"github.com/BFMVAUS/d2go/pkg/data"
	"github.com/BFMVAUS/d2go/pkg/data/area"
	"github.com/BFMVAUS/koolov2/internal/action"
	"github.com/BFMVAUS/koolov2/internal/config"
	"github.com/BFMVAUS/koolov2/internal/context"
)

type Mausoleum struct {
	ctx *context.Status
}

func NewMausoleum() *Mausoleum {
	return &Mausoleum{
		ctx: context.Get(),
	}
}

func (a Mausoleum) Name() string {
	return string(config.MausoleumRun)
}

func (a Mausoleum) Run() error {

	// Define a defaut filter
	monsterFilter := data.MonsterAnyFilter()

	// Update filter if we selected to clear only elites
	if a.ctx.CharacterCfg.Game.Mausoleum.FocusOnElitePacks {
		monsterFilter = data.MonsterEliteFilter()
	}

	// Use the waypoint
	err := action.WayPoint(area.ColdPlains)
	if err != nil {
		return err
	}

	// Move to the BurialGrounds
	if err = action.MoveToArea(area.BurialGrounds); err != nil {
		return err
	}

	// Move to the Mausoleum
	if err = action.MoveToArea(area.Mausoleum); err != nil {
		return err
	}

	// Open a TP If we're the leader
	action.OpenTPIfLeader()

	// Clear the area
	return action.ClearCurrentLevel(a.ctx.CharacterCfg.Game.Mausoleum.OpenChests, monsterFilter)
}
