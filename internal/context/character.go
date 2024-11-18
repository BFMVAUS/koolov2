package context

import (
	"github.com/BFMVAUS/d2go/pkg/data"
	"github.com/BFMVAUS/d2go/pkg/data/skill"
	"github.com/BFMVAUS/d2go/pkg/data/stat"
	"github.com/BFMVAUS/koolov2/internal/game"
)

type Character interface {
	CheckKeyBindings() []skill.ID
	BuffSkills() []skill.ID
	PreCTABuffSkills() []skill.ID
	KillCountess() error
	KillAndariel() error
	KillSummoner() error
	KillDuriel() error
	KillMephisto() error
	KillPindle() error
	KillNihlathak() error
	KillCouncil() error
	KillDiablo() error
	KillIzual() error
	KillBaal() error
	KillMonsterSequence(
		monsterSelector func(d game.Data) (data.UnitID, bool),
		skipOnImmunities []stat.Resist,
	) error
}

type LevelingCharacter interface {
	Character
	// StatPoints Stats will be assigned in the order they are returned by this function.
	StatPoints() map[stat.ID]int
	SkillPoints() []skill.ID
	SkillsToBind() (skill.ID, []skill.ID)
	ShouldResetSkills() bool
	KillAncients() error
}
