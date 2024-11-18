package run

import (
	"github.com/BFMVAUS/d2go/pkg/data"
	"github.com/BFMVAUS/d2go/pkg/data/area"
	"github.com/BFMVAUS/koolov2/internal/action"
	"github.com/BFMVAUS/koolov2/internal/config"
	"github.com/BFMVAUS/koolov2/internal/context"
)

type LowerKurast struct {
	ctx *context.Status
}

func NewLowerKurast() *LowerKurast {
	return &LowerKurast{
		ctx: context.Get(),
	}
}

func (a LowerKurast) Name() string {
	return string(config.LowerKurastRun)
}

func (a LowerKurast) Run() error {

	// Use Waypoint to Lower Kurast
	err := action.WayPoint(area.LowerKurast)
	if err != nil {
		return err
	}

	// Clear Lower Kurast
	return action.ClearCurrentLevel(true, data.MonsterAnyFilter())

}
