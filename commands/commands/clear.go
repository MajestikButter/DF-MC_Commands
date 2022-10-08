package commands

import (
	"math"
	"strings"

	"github.com/MajestikButter/DF-MC_Commands/cmdtypes"
	"github.com/MajestikButter/DF-MC_Commands/utils"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type Clear struct {
	Player   cmd.Optional[[]cmd.Target]  `cmd:"target"`
	ItemName cmd.Optional[cmdtypes.Item] `cmd:"itemName"`
	Data     cmd.Optional[int16]         `cmd:"data"`
	MaxCount cmd.Optional[int]           `cmd:"maxCount"`
}

func clearItem(player *player.Player, itemName cmdtypes.Item, data int16, maxCount int) int {
	inv := player.Inventory()

	name := string(itemName)
	name = strings.Replace(name, "minecraft:", "", 1)

	count := maxCount + 0
	cleared := 0
	for i, item := range inv.Slots() {
		if item.Count() <= 0 {
			continue
		}
		id, aux := item.Item().EncodeItem()
		id = strings.Replace(id, "minecraft:", "", 1)

		if itemName == "" || (name == id && (data < 0 || data == aux)) {
			itemCount := item.Count()
			if maxCount == 0 {
				cleared += itemCount
				continue
			}
			if itemCount > count && maxCount != -1 {
				inv.SetItem(i, item.Grow(-count))
				cleared += count
				return cleared
			} else {
				count -= itemCount
				cleared += itemCount
				inv.RemoveItem(item)
			}
		}
	}
	return cleared
}

func (t Clear) Run(source cmd.Source, output *cmd.Output) {
	targets, hasTarget := t.Player.Load()

	itemName, nameUsed := t.ItemName.Load()
	if !nameUsed {
		itemName = ""
	}

	data, dataUsed := t.Data.Load()
	if !dataUsed {
		data = -1
	}

	maxCount, countUsed := t.MaxCount.Load()
	if !countUsed {
		maxCount = -1
	} else {
		maxCount = int(math.Max(0, float64(maxCount)))
	}

	if hasTarget {
		var players = make([]*player.Player, 0)
		for _, target := range targets {
			plr, ok := utils.ToPlayer(target)
			if !ok {
				output.Error("Invalid selector")
				return
			}
			players = append(players, plr)
		}

		for _, player := range players {
			res := clearItem(player, itemName, data, maxCount)
			if res > 0 {
				output.Printf("Cleared %d items from %s", res, player.Name())
			} else {
				output.Errorf("Unable to find items to clear from %s", player.Name())
			}
		}
	} else {
		plr, ok := utils.ToPlayer(source)
		if !ok {
			output.Error("Unable to clear items from non-player")
			return
		}
		res := clearItem(plr, itemName, data, maxCount)
		if res > 0 {
			output.Printf("Cleared %d items", res)
		} else {
			output.Error("Unable to find items to clear")
		}
	}
}

func (t Clear) Allow(source cmd.Source) bool {
	return utils.CommandPermission(source, "minecraft.chat.command.clear")
}
