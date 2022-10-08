package cmdtypes

import (
	"strings"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
)

type Item string

func (Item) Type() string {
	return "Item"
}

// Options returns a number for every plot the player has.
func (Item) Options(source cmd.Source) []string {
	items := map[string]bool{}
	for _, v := range world.Items() {
		id, _ := v.EncodeItem()
		items[strings.Replace(id, "minecraft:", "", 1)] = true
	}

	keys := []string{}
	for k := range items {
		keys = append(keys, k)
	}

	return keys
}
