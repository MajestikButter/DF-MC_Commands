package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/MajestikButter/DF-MC_Commands/commands/utils"
)

var source = &Console{}
var silencedSource = &Console{true}

func ExecuteCommand(cmdStr string, output bool) {
	command, commandName := utils.FindCommand(cmdStr)
	if command == nil {
		if output {
			fmt.Printf("Unknown command '%v'\n", commandName)
		}
		return
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
