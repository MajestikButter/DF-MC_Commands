package cmdtypes

import (
	"reflect"
	"strings"

	"github.com/df-mc/dragonfly/server/cmd"
)

type Message string

func (Message) Parse(line *cmd.Line, v reflect.Value) error {
	v.SetString(strings.Join(line.Leftover(), " "))
	return nil
}

func (Message) Type() string {
	return "message"
}
