package commands

import (
	"golang.org/x/exp/slices"

	"github.com/df-mc/dragonfly/server/cmd"
)

var commands = []cmd.Command{
	cmd.New("gamemode", "Changes the player to a specific game mode.", []string{"gm", "gamemode"}, GameMode{}),
	cmd.New("stop", "Stop the server.", []string{"stop"}, Stop{}),
	cmd.New("transferserver", "Transfer players to another server.", []string{"transferserver", "transfer", "tserv"}, TransferServer{}),
	cmd.New("reload", "Reloads files.", []string{"reload"}, Reload{}),
	cmd.New("clear", "Remove items from a player's inventory", []string{"clear"}, Clear{}),
	cmd.New("give", "Give items to a player", []string{"give"}, Give{}),
	cmd.New("npc", "Manage NPCs", []string{"npc"}, NPCCreate{}, NPCEditGeometry{}, NPCEditName{}, NPCEditTexture{}, NPCEditAction{}),
	cmd.New("sudo", "Run a command or send a message as a player", []string{"sudo"}, Sudo{}),
	cmd.New("function", "Run a set of commands from an mcfunction file", []string{"function"}, Function{}),
}

func LoadCommands(withOut []string) {
	for _, c := range commands {
		if slices.Contains(withOut, c.Name()) {
			continue
		}

		cmd.Register(c)
	}
}
