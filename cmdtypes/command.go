package cmdtypes

import (
	"reflect"
	"strings"

	"github.com/df-mc/dragonfly/server/cmd"
)

type Command string

// Parse takes an arbitrary amount of arguments from the command Line passed and parses it, so that it can
// store it to value v. If the arguments cannot be parsed from the Line, an error should be returned.
func (Command) Parse(line *cmd.Line, v reflect.Value) error {
	v.SetString(strings.Join(line.Leftover(), " "))
	return nil
}

func (Command) Type() string {
	return "command"
}
