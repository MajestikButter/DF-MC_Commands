package commands

import (
	"errors"
	"os"
	"path"
	"strings"
	"time"

	"github.com/MajestikButter/DF-MC_Commands/commands/commands"
	"github.com/MajestikButter/DF-MC_Commands/commands/console"
	"github.com/MajestikButter/DF-MC_Commands/commands/shared"
	"github.com/MajestikButter/DF-MC_Permissions/permissions"
	"github.com/df-mc/dragonfly/server"
)

func Load(server *server.Server, permSystem *permissions.PermissionSystem, withoutCommands []string) {
	shared.PermSystem = permSystem
	shared.Server = server

	commands.LoadCommands(withoutCommands)
}

func StartConsole() {
	console.StartConsole()
}

func loadFuncDir(base, subDir string, r map[string][]string) error {
	files, err := os.ReadDir(path.Join(base, subDir))
	if err != nil {
		return err
	}

	for _, e := range files {
		if e.IsDir() {
			sd := path.Join(subDir, e.Name())
			err := loadFuncDir(base, sd, r)
			if err != nil {
				return err
			}
		} else if strings.HasSuffix(e.Name(), ".mcfunction") {
			n := e.Name()
			c, err := os.ReadFile(path.Join(base, subDir, n))
			if err != nil {
				return err
			}
			r[path.Join(subDir, n[:len(n)-11])] = strings.Split(string(c), "\n")
		}
	}
	return nil
}

func LoadFunctions(dir string) error {
	p := ""
	if path.IsAbs(dir) {
		p = dir
	} else {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		p = path.Join(cwd, dir)
	}

	res := map[string][]string{}
	err := loadFuncDir(p, "", res)
	if err != nil {
		return err
	}

	shared.Functions = res
	return nil
}

func StartTickFunctions() error {
	if _, ok := shared.Functions["tick.json"]; !ok {
		return errors.New("unable to start tick functions, no tick.json has been created in the function directory")
	}

	go func() {
		time.Sleep(time.Second / 20)

		console.ExecuteCommands(shared.Functions["tick.json"], nil)
	}()
	return nil
}
