package utils

import (
	"github.com/MajestikButter/DF-MC_Commands/cmdtypes"
	system "github.com/MajestikButter/DF-MC_Commands/shared"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

func StringToGameMode(arg string) world.GameMode {
	return cmdtypes.GameModeMap[arg]
}

func GameModeToName(gm world.GameMode) string {
	return map[world.GameMode]string{
		world.GameModeSurvival:  "survival",
		world.GameModeCreative:  "creative",
		world.GameModeAdventure: "adventure",
		world.GameModeSpectator: "spectator",
	}[gm]
}

func CommandPermission(source cmd.Source, perm string) bool {
	s, so := source.(cmdtypes.SudoSrc)
	if so {
		if s.IgnorePerm {
			return true
		} else {
			return system.PermSystem.GetCommandPermission(s.Player, perm)
		}
	}
	return system.PermSystem.GetCommandPermission(source, perm)
}

func ToPlayer(source interface{}) (*player.Player, bool) {
	plr, p := source.(*player.Player)
	if p {
		return plr, true
	}
	sudo, s := source.(cmdtypes.SudoSrc)
	if s {
		return sudo.Player, true
	}
	return nil, false
}
