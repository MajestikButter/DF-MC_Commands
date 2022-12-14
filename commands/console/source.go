package console

import (
	"fmt"

	system "github.com/MajestikButter/DF-MC_Commands/commands/shared"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Console struct {
	silence bool
}

func (c *Console) SendCommandOutput(output *cmd.Output) {
	if c.silence {
		return
	}
	for _, m := range output.Messages() {
		fmt.Println(m)
	}

	for _, e := range output.Errors() {
		fmt.Println(e.Error())
	}
}

func (*Console) Name() string {
	return "Console"
}

func (*Console) Position() mgl64.Vec3 {
	return mgl64.Vec3{}
}

func (*Console) World() *world.World {
	return system.Server.World()
}
