package cmdtypes

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
)

type Gamemode string

func (Gamemode) Type() string {
	return "mode"
}

func (Gamemode) Options(source cmd.Source) []string {
	keys := make([]string, 0, len(GameModeMap))
	for k := range GameModeMap {
		keys = append(keys, k)
	}
	return keys
}

var GameModeMap = map[string]world.GameMode{
	"0": world.GameModeSurvival, "s": world.GameModeSurvival, "survival": world.GameModeSurvival,
	"1": world.GameModeCreative, "c": world.GameModeCreative, "creative": world.GameModeCreative,
	"2": world.GameModeAdventure, "a": world.GameModeAdventure, "adventure": world.GameModeAdventure,
	"3": world.GameModeSpectator, "sp": world.GameModeSpectator, "spectator": world.GameModeSpectator,
}
