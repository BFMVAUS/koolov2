package run

import (
	"github.com/BFMVAUS/d2go/pkg/data"
	"github.com/BFMVAUS/d2go/pkg/data/area"
	"github.com/BFMVAUS/d2go/pkg/data/npc"
	"github.com/BFMVAUS/koolov2/internal/action"
	"github.com/BFMVAUS/koolov2/internal/config"
	"github.com/BFMVAUS/koolov2/internal/context"
)

type Summoner struct {
	ctx *context.Status
}

func NewSummoner() *Summoner {
	return &Summoner{
		ctx: context.Get(),
	}
}

func (s Summoner) Name() string {
	return string(config.SummonerRun)
}

func (s Summoner) Run() error {

	// Use the waypoint
	err := action.WayPoint(area.ArcaneSanctuary)
	if err != nil {
		return err
	}

	// Move to boss position
	if err = action.MoveTo(func() (data.Position, bool) {
		m, found := s.ctx.Data.NPCs.FindOne(npc.Summoner)
		return m.Positions[0], found
	}); err != nil {
		return err
	}

	// Kill Summoner
	return s.ctx.Char.KillSummoner()
}
