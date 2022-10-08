package commands

import (
	"fmt"
	"strings"

	"github.com/MajestikButter/DF-MC_Commands/cmdtypes"
	"github.com/MajestikButter/DF-MC_Commands/utils"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type Sudo struct {
	Target []cmd.Target  `cmd:"target"`
	Sudo   cmdtypes.Sudo `cmd:"sudo"`
}

func ExecuteCommand(source cmd.Source, commandLine string, p *player.Player, o *cmd.Output, ignorePerm bool) error {
	args := strings.Split(commandLine, " ")
	command, ok := cmd.ByAlias(args[0][1:])
	if !ok {
		//lint:ignore ST1005 Error string is capitalised because it is shown to the player.
		return fmt.Errorf("Unknown command: %v", args[0])
	}
	sudo := cmdtypes.SudoSrc{Player: p, OrgSource: source, IgnorePerm: ignorePerm}
	command.Execute(strings.Join(args[1:], " "), sudo)
	return nil
}

func (t Sudo) Run(source cmd.Source, output *cmd.Output) {
	sudo := string(t.Sudo)
	chat := !strings.HasPrefix(sudo, "/")
	var players = make([]*player.Player, 0)
	for _, target := range t.Target {
		plr, ok := utils.ToPlayer(target)
		if !ok {
			output.Error("Invalid selector")
			return
		}
		players = append(players, plr)
	}

	for _, plr := range players {
		if chat {
			plr.Chat(sudo)
		} else {
			err := ExecuteCommand(source, sudo, plr, output, true)
			if err != nil {
				output.Error(err.Error())
				continue
			}
		}
		output.Printf("Sudo-ed %s successfully", plr.Name())
	}
}

func (t Sudo) Allow(source cmd.Source) bool {
	return utils.CommandPermission(source, "minecraft.chat.command.sudo")
}
