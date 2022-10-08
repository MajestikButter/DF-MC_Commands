package commands

import (
	"errors"
	"time"

	"github.com/MajestikButter/DF-MC_Commands/commands/commands"
	"github.com/MajestikButter/DF-MC_Commands/commands/console"
	"github.com/MajestikButter/DF-MC_Commands/commands/shared"
	"github.com/MajestikButter/DF-MC_Permissions/permissions"
	"github.com/df-mc/dragonfly/server"
)

func Load(server *server.Server, permSystem *permissions.PermissionSystem, withoutCommands []string) {
	shared.PermSystem = permSystem
	shared.Server = server

	commands.LoadCommands(withoutCommands)
}

func StartConsole() {
	console.StartConsole()
}

func LoadFunctions(dir string) error {
	shared.FunctionsDir = dir
	return commands.LoadFunctions()
}

func StartTickFunctions() error {
	if _, ok := shared.Functions["tick.json"]; !ok {
		return errors.New("unable to start tick functions, no tick.json has been created in the function directory")
	}

	go func() {
		for {
			time.Sleep(time.Second / 20)

			console.ExecuteCommand("function tick.json", false)
		}
	}()
	return nil
}
