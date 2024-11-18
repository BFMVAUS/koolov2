package step

import (
	"time"

	"github.com/BFMVAUS/d2go/pkg/data/object"
	"github.com/BFMVAUS/d2go/pkg/data/skill"
	"github.com/BFMVAUS/koolov2/internal/context"
	"github.com/BFMVAUS/koolov2/internal/game"
	"github.com/BFMVAUS/koolov2/internal/utils"
)

func OpenPortal() error {
	ctx := context.Get()
	ctx.SetLastStep("OpenPortal")

	lastRun := time.Time{}
	for {
		// Pause the execution if the priority is not the same as the execution priority
		ctx.PauseIfNotPriority()

		_, found := ctx.Data.Objects.FindOne(object.TownPortal)
		if found {
			return nil
		}

		// Give some time to portal to popup before retrying...
		if time.Since(lastRun) < time.Second*2 {
			continue
		}

		ctx.HID.PressKeyBinding(ctx.Data.KeyBindings.MustKBForSkill(skill.TomeOfTownPortal))
		utils.Sleep(250)
		ctx.HID.Click(game.RightButton, 300, 300)
		lastRun = time.Now()
	}
}
