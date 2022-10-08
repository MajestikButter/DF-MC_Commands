package cmdtypes

import (
	"strings"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
)

type Entity string

func (Entity) Type() string {
	return "Entity"
}

// Options returns a number for every plot the player has.
func (Entity) Options(source cmd.Source) []string {
	entities := map[string]bool{}
	for _, v := range world.Entities() {
		id := v.EncodeEntity()
		entities[strings.Replace(id, "minecraft:", "", 1)] = true
	}

	keys := []string{}
	for k := range entities {
		keys = append(keys, k)
	}

	return keys
}
