package main

import (
	"github.com/MajestikButter/DF-MC_Commands/commands"
	"github.com/MajestikButter/DF-MC_Commands/console"
	"github.com/MajestikButter/DF-MC_Commands/shared"
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
