package step

import (
	"errors"
	"fmt"
	"time"

	"github.com/BFMVAUS/d2go/pkg/data"
	"github.com/BFMVAUS/d2go/pkg/data/area"
	"github.com/BFMVAUS/koolov2/internal/context"
	"github.com/BFMVAUS/koolov2/internal/game"
	"github.com/BFMVAUS/koolov2/internal/utils"
)

func InteractEntrance(area area.ID) error {
	maxInteractionAttempts := 5
	interactionAttempts := 0
	waitingForInteraction := false
	currentMouseCoords := data.Position{}
	lastRun := time.Time{}

	ctx := context.Get()
	ctx.SetLastStep("InteractEntrance")

	for {
		// Pause the execution if the priority is not the same as the execution priority
		ctx.PauseIfNotPriority()

		// Give some extra time to render the UI
		if ctx.Data.AreaData.Area == area && time.Since(lastRun) > time.Millisecond*500 && ctx.Data.AreaData.IsInside(ctx.Data.PlayerUnit.Position) {
			// We've successfully entered the new area
			return nil
		}

		if interactionAttempts > maxInteractionAttempts {
			return fmt.Errorf("area %s [%d] could not be interacted", area.Area().Name, area)
		}

		if waitingForInteraction && time.Since(lastRun) < time.Millisecond*500 {
			continue
		}

		lastRun = time.Now()
		for _, l := range ctx.Data.AdjacentLevels {
			if l.Area == area {
				distance := ctx.PathFinder.DistanceFromMe(l.Position)
				if distance > 10 {
					return errors.New("entrance too far away")
				}

				if l.IsEntrance {
					lx, ly := ctx.PathFinder.GameCoordsToScreenCords(l.Position.X-2, l.Position.Y-2)
					if ctx.Data.HoverData.UnitType == 5 || ctx.Data.HoverData.UnitType == 2 && ctx.Data.HoverData.IsHovered {
						ctx.HID.Click(game.LeftButton, currentMouseCoords.X, currentMouseCoords.Y)
						waitingForInteraction = true
					}

					x, y := utils.Spiral(interactionAttempts)
					currentMouseCoords = data.Position{X: lx + x, Y: ly + y}
					ctx.HID.MovePointer(lx+x, ly+y)
					interactionAttempts++
					continue
				}

				return fmt.Errorf("area %s [%d] is not an entrance", area.Area().Name, area)
			}
		}
	}
}
