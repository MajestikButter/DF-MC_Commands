package commands

import (
	"github.com/MajestikButter/DF-MC_Commands/commands/cmdtypes"
	"github.com/MajestikButter/DF-MC_Commands/commands/utils"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type Message struct {
	Target  []cmd.Target     `cmd:"target"`
	Message cmdtypes.Message `cmd:"message"`
}

func (t Message) Run(source cmd.Source, output *cmd.Output) {
	targets := t.Target
	players := make([]*player.Player, 0)
	for _, target := range targets {
		plr, ok := utils.ToPlayer(target)
		if !ok {
			output.Error("Invalid selector")
			return
		}
		players = append(players, plr)
	}

	for _, plr := range players {
		plr.Message(t.Message)
	}
}

func (t Message) Allow(source cmd.Source) bool {
	return utils.CommandPermission(source, "minecraft.chat.command.message")
}
