package step

import (
	"github.com/BFMVAUS/d2go/pkg/data/skill"
	"github.com/BFMVAUS/koolov2/internal/context"
)

func SetSkill(id skill.ID) {
	ctx := context.Get()
	ctx.SetLastStep("SetSkill")

	if kb, found := ctx.Data.KeyBindings.KeyBindingForSkill(id); found {
		if ctx.Data.PlayerUnit.RightSkill != id {
			ctx.HID.PressKeyBinding(kb)
		}
	}
}
