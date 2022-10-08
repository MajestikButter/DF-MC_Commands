package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/df-mc/dragonfly/server/cmd"
)

func StartConsole() {
	go func() {
		time.Sleep(time.Millisecond * 500)
		source := &Console{}
		fmt.Println("Type help for commands.")
		scanner := bufio.NewScanner(os.Stdin)
		for {
			if scanner.Scan() {
				commandString := scanner.Text()
				if len(commandString) == 0 {
					continue
				}
				commandName := strings.Split(commandString, " ")[0]
				command, ok := cmd.ByAlias(commandName)

				if !ok {
					output := &cmd.Output{}
					output.Errorf("Unknown command '%v'", commandName)
					for _, e := range output.Errors() {
						fmt.Println(e)
					}
					continue
				}

				command.Execute(strings.TrimPrefix(strings.TrimPrefix(commandString, commandName), " "), source)
			}
		}
	}()
}