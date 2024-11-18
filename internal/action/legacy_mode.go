package action

import (
	"github.com/BFMVAUS/koolov2/internal/context"
	"github.com/BFMVAUS/koolov2/internal/game"
	"github.com/BFMVAUS/koolov2/internal/ui"
	"github.com/BFMVAUS/koolov2/internal/utils"
)

func SwitchToLegacyMode() {
	ctx := context.Get()
	ctx.SetLastAction("SwitchToLegacyMode")

	if ctx.CharacterCfg.ClassicMode && !ctx.Data.LegacyGraphics {
		ctx.Logger.Debug("Switching to legacy mode...")
		ctx.HID.PressKey(ctx.Data.KeyBindings.LegacyToggle.Key1[0])
		utils.Sleep(500) // Add small delay to allow the game to switch

		// Close the mini panel if option is enabled
		if ctx.CharacterCfg.CloseMiniPanel {
			utils.Sleep(100)
			ctx.HID.Click(game.LeftButton, ui.CloseMiniPanelClassicX, ui.CloseMiniPanelClassicY)
			utils.Sleep(100)
		}
	}
}
