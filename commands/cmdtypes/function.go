package cmdtypes

import (
	"github.com/MajestikButter/DF-MC_Commands/commands/shared"
	"github.com/df-mc/dragonfly/server/cmd"
)

type Function string

func (Function) Type() string {
	return "Function"
}

func (Function) Options(source cmd.Source) []string {
	keys := []string{}
	for k := range shared.Functions {
		keys = append(keys, k)
	}

	return keys
}
