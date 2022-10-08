package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/MajestikButter/DF-MC_Commands/commands/utils"
	"github.com/df-mc/dragonfly/server/cmd"
)

var source = &Console{}
var silencedSource = &Console{true}

func ExecuteCommands(cmds []string, output *cmd.Output) {
	count := 0
	for _, cmd := range cmds {
		if strings.TrimSpace(cmd) == "" {
			continue
		}
		ExecuteCommand(cmd, false)
		count++
	}
	if output != nil {
		output.Printf("Ran %v commands", count)
	}
}

func ExecuteCommand(cmdStr string, output bool) {
	command, commandName := utils.FindCommand(cmdStr)
	if output {
		output := &cmd.Output{}
		output.Errorf("Unknown command '%v'", commandName)
		for _, e := range output.Errors() {
			fmt.Println(e)
		}
	}

	args := strings.TrimPrefix(strings.TrimPrefix(cmdStr, commandName), " ")
	if output {
		command.Execute(args, source)
	} else {
		command.Execute(args, silencedSource)
	}
}
func StartConsole() {
	go func() {
		time.Sleep(time.Millisecond * 500)
		fmt.Println("Type help for commands.")
		scanner := bufio.NewScanner(os.Stdin)
		for {
			if scanner.Scan() {
				commandString := scanner.Text()
				if len(commandString) == 0 {
					continue
				}
				ExecuteCommand(commandString, true)
			}
		}
	}()
}
