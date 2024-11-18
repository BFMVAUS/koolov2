package step

import (
	"time"

	"github.com/BFMVAUS/d2go/pkg/data/skill"
	"github.com/BFMVAUS/koolov2/internal/context"
)

func SwapToMainWeapon() error {
	return swapWeapon(false)
}

func SwapToCTA() error {
	return swapWeapon(true)
}

func swapWeapon(toCTA bool) error {
	lastRun := time.Time{}

	ctx := context.Get()
	ctx.SetLastStep("SwapToCTA")

	for {
		// Pause the execution if the priority is not the same as the execution priority
		ctx.PauseIfNotPriority()

		if time.Since(lastRun) < time.Millisecond*500 {
			continue
		}

		_, found := ctx.Data.PlayerUnit.Skills[skill.BattleOrders]
		if (toCTA && found) || (!toCTA && !found) {
			return nil
		}

		ctx.HID.PressKeyBinding(ctx.Data.KeyBindings.SwapWeapons)

		lastRun = time.Now()
	}
}
