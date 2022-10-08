package commands

import (
	"github.com/MajestikButter/DF-MC_Commands/commands/utils"
	"github.com/go-gl/mathgl/mgl64"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type TeleportPosition struct {
	Target   []cmd.Target `cmd:"target"`
	Position mgl64.Vec3   `cmd:"position"`
}

func (t TeleportPosition) Run(source cmd.Source, output *cmd.Output) {
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
		plr.Teleport(t.Position)
	}
}

func (t TeleportPosition) Allow(source cmd.Source) bool {
	return utils.CommandPermission(source, "minecraft.chat.command.teleport.position")
}

type TeleportPlayer struct {
	Who    []cmd.Target `cmd:"who"`
	Target []cmd.Target `cmd:"target"`
}

func (t TeleportPlayer) Run(source cmd.Source, output *cmd.Output) {
	whos := t.Who
	players := make([]*player.Player, 0)
	for _, who := range whos {
		plr, ok := utils.ToPlayer(who)
		if !ok {
			output.Error("Invalid selector")
			return
		}
		players = append(players, plr)
	}

	if len(t.Target) > 1 {
		return
	}
	target := t.Target[0].(*player.Player)

	for _, plr := range players {
		plr.Teleport(target.Position())
	}
}

func (t TeleportPlayer) Allow(source cmd.Source) bool {
	return utils.CommandPermission(source, "minecraft.chat.command.teleport.position")
}
