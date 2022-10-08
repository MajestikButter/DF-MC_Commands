package commands

import (
	"fmt"

	"github.com/MajestikButter/DF-MC_Commands/commands/utils"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type TransferServer struct {
	Target  []cmd.Target      `cmd:"target"`
	Address string            `cmd:"address"`
	Port    cmd.Optional[int] `cmd:"port"`
}

func (t TransferServer) Run(source cmd.Source, output *cmd.Output) {
	targets := t.Target

	var players = make([]*player.Player, 0)
	for _, target := range targets {
		plr, ok := utils.ToPlayer(target)
		if !ok {
			output.Error("Invalid selector")
			return
		}
		players = append(players, plr)
	}

	port, used := t.Port.Load()
	if !used {
		port = 19132
	}

	address := fmt.Sprintf("%s:%d", t.Address, port)

	for _, player := range players {
		err := player.Transfer(address)
		if err != nil {
			output.Errorf("Error occurred while transferring %s: %s", player.Name(), err.Error())
		} else {
			output.Printf("Transfering %v to %s", player.Name(), address)
		}
	}
}

func (t TransferServer) Allow(source cmd.Source) bool {
	return utils.CommandPermission(source, "minecraft.chat.command.transferserver")
}
