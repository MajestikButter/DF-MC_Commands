package commands

import (
	"fmt"

	"github.com/MajestikButter/DF-MC_Commands/commands/cmdtypes"
	"github.com/MajestikButter/DF-MC_Commands/commands/utils"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player/chat"
)

type Say struct {
	Message cmdtypes.Message `cmd:"message"`
}

func (t Say) Run(source cmd.Source, output *cmd.Output) {
	msg := fmt.Sprintf("[%s] %s", source.Name(), string(t.Message))
	chat.Global.WriteString(msg)
}

func (t Say) Allow(source cmd.Source) bool {
	return utils.CommandPermission(source, "minecraft.chat.command.say")
}
