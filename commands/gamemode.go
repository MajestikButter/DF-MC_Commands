package commands

import (
	"github.com/MajestikButter/DF-MC_Commands/cmdtypes"
	"github.com/MajestikButter/DF-MC_Commands/utils"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type GameMode struct {
	GameMode cmdtypes.Gamemode          `cmd:"gamemode"`
	Target   cmd.Optional[[]cmd.Target] `cmd:"target"`
}

func (t GameMode) Run(source cmd.Source, output *cmd.Output) {
	mode := utils.StringToGameMode(string(t.GameMode))
	modeString := utils.GameModeToName(mode)

	targets, used := t.Target.Load()

	if p, ok := utils.ToPlayer(source); ok && !used {
		p.SetGameMode(mode)
		output.Printf("Set own game mode to %s.", modeString)
		return
	}

	if len(targets) > 0 {
		var players = make([]*player.Player, 0)
		for _, target := range targets {
			plr, ok := utils.ToPlayer(target)
			if !ok {
				output.Error("Invalid selector")
				return
			}
			players = append(players, plr)
		}
		for _, p := range players {
			p.SetGameMode(mode)
			output.Printf("Set %s's gamemode to %s.", p.Name(), modeString)
			p.Messagef("Your game mode has been changed to %s.", modeString)
		}
	} else {
		output.Error("Usage: /gamemode <GameMode: mode> [Target: target]")
	}
}

func (t GameMode) Allow(source cmd.Source) bool {
	return utils.CommandPermission(source, "minecraft.chat.command.gamemode")
}
