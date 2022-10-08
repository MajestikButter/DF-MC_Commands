package utils

import (
	"strings"

	"github.com/MajestikButter/DF-MC_Commands/commands/cmdtypes"
	system "github.com/MajestikButter/DF-MC_Commands/commands/shared"
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

func FindCommand(cmdStr string) (*cmd.Command, string) {
	commandName := strings.Split(cmdStr, " ")[0]
	command, _ := cmd.ByAlias(commandName)

	return &command, commandName
}
