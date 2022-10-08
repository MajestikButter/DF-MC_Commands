package commands

import (
	"github.com/MajestikButter/DF-MC_Commands/commands/shared"
	"github.com/MajestikButter/DF-MC_Commands/commands/utils"

	"github.com/df-mc/dragonfly/server/cmd"
)

type Reload struct{}

func (t Reload) Run(source cmd.Source, output *cmd.Output) {
	shared.PermSystem.Load()

	LoadFunctions()

	output.Printf("Reloaded files.")
}

func (t Reload) Allow(source cmd.Source) bool {
	return utils.CommandPermission(source, "minecraft.chat.command.reload")
}
