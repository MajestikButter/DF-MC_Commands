package cmdtypes

import (
	"reflect"
	"strings"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Sudo string

// Parse takes an arbitrary amount of arguments from the command Line passed and parses it, so that it can
// store it to value v. If the arguments cannot be parsed from the Line, an error should be returned.
func (Sudo) Parse(line *cmd.Line, v reflect.Value) error {
	v.SetString(strings.Join(line.Leftover(), " "))
	return nil
}

func (Sudo) Type() string {
	return "sudo"
}

type SudoSrc struct {
	*player.Player

	OrgSource  cmd.Source
	IgnorePerm bool
}

func (s SudoSrc) SendCommandOutput(o *cmd.Output) {
	s.OrgSource.SendCommandOutput(o)
}

func (SudoSrc) Name() string {
	return "Sudo"
}

func (s SudoSrc) Position() mgl64.Vec3 {
	return s.OrgSource.Position()
}

func (s SudoSrc) World() *world.World {
	return s.OrgSource.World()
}
