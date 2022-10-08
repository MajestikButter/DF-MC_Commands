package commands

import (
	"math"
	"strings"

	"github.com/MajestikButter/DF-MC_Commands/cmdtypes"
	"github.com/MajestikButter/DF-MC_Commands/utils"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type Give struct {
	Player   []cmd.Target        `cmd:"target"`
	ItemName cmdtypes.Item       `cmd:"itemName"`
	Amount   cmd.Optional[int16] `cmd:"amount"`
	Data     cmd.Optional[int16] `cmd:"data"`
}

func giveItem(player *player.Player, itemType world.Item, data int16, amount int) int {
	inv := player.Inventory()

	stack := item.NewStack(itemType, amount)

	added, err := inv.AddItem(stack)
	if err != nil && !player.GameMode().CreativeInventory() {
		left := amount - added

		max := stack.MaxCount()
		iterations := math.Ceil(float64(left) / float64(max))
		for i := 0.0; i < iterations; i++ {
			if left > max {
				left -= max
				newStack := item.NewStack(itemType, max)
				added += player.Drop(newStack)
			} else {
				newStack := item.NewStack(itemType, left)
				added += player.Drop(newStack)
			}
		}
	}
	return added
}

func (t Give) Run(source cmd.Source, output *cmd.Output) {
	targets := t.Player
	itemName := t.ItemName

	data, dataUsed := t.Data.Load()
	if !dataUsed {
		data = 0
	}

	name := string(itemName)
	if !strings.Contains(name, ":") {
		name = "minecraft:" + name
	}

	itemType, isItem := world.ItemByName(name, data)

	if !isItem || name == "air" {
		output.Errorf("Unknown item identifier provided %s", name)
		return
	}

	amount16, amountUsed := t.Amount.Load()
	amount := 0
	if !amountUsed {
		amount = 1
	} else {
		amount = int(math.Max(1, float64(amount16)))
	}

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
		given := giveItem(player, itemType, data, amount)
		output.Printf("Gave %s x%d to %s", name, given, player.Name())
	}
}

func (t Give) Allow(source cmd.Source) bool {
	return utils.CommandPermission(source, "minecraft.chat.command.give")
}
