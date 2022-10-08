package commands

import (
	system "github.com/MajestikButter/DF-MC_Commands/commands/shared"
	"github.com/MajestikButter/DF-MC_Commands/commands/utils"

	"github.com/df-mc/dragonfly/server/cmd"
)

type Stop struct{}

func (t Stop) Run(source cmd.Source, output *cmd.Output) {
	output.Printf("Stopping server.")
	if err := system.Server.Close(); err != nil {
		output.Error("Error shutting down server: %v", err)
	}
}

func (t Stop) Allow(source cmd.Source) bool {
	return utils.CommandPermission(source, "minecraft.chat.command.stop")
}
