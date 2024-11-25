package run

import (
	"github.com/BFMVAUS/d2go/pkg/data"
	"github.com/BFMVAUS/d2go/pkg/data/area"
	"github.com/BFMVAUS/d2go/pkg/data/npc"
	"github.com/BFMVAUS/koolov2/internal/action"
	"github.com/BFMVAUS/koolov2/internal/config"
	"github.com/BFMVAUS/koolov2/internal/context"
	"github.com/BFMVAUS/koolov2/internal/game"
)

type Threshsocket struct {
	ctx *context.Status
}

func NewThreshsocket() *Threshsocket {
	return &Threshsocket{
		ctx: context.Get(),
	}
}

func (t Threshsocket) Name() string {
	return string(config.ThreshsocketRun)
}

func (t Threshsocket) Run() error {

	// Use waypoint to crystalinepassage
	err := action.WayPoint(area.CrystallinePassage)
	if err != nil {
		return err
	}

	// Move to ArreatPlateau
	if err = action.MoveToArea(area.ArreatPlateau); err != nil {
		return err
	}

	// Kill Threshsocket
	return t.ctx.Char.KillMonsterSequence(func(d game.Data) (data.UnitID, bool) {
		if m, found := d.Monsters.FindOne(npc.BloodBringer, data.MonsterTypeSuperUnique); found {
			return m.UnitID, true
		}

		return 0, false
	}, nil)
}
