package commands

import (
	"github.com/MajestikButter/DF-MC_Commands/commands/console"
	system "github.com/MajestikButter/DF-MC_Commands/commands/shared"
	"github.com/MajestikButter/DF-MC_Commands/commands/utils"

	"github.com/df-mc/dragonfly/server/cmd"
)

type Function struct {
	Function string `cmd:"function"`
}

func (f Function) Run(source cmd.Source, output *cmd.Output) {
	cmds, ok := system.Functions[f.Function]
	if !ok {
		output.Errorf("Unable to find function %s", f.Function)
		return
	}
	console.ExecuteCommands(cmds, output)
}

func (f Function) Allow(source cmd.Source) bool {
	return utils.CommandPermission(source, "minecraft.chat.command.function")
}
