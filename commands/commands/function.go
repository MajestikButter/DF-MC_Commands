package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/MajestikButter/DF-MC_Commands/commands/shared"
	"github.com/MajestikButter/DF-MC_Commands/commands/utils"

	"github.com/df-mc/dragonfly/server/cmd"
)

func loadFuncDir(base, subDir string, r *map[string][]string) error {
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
			f := path.Join(subDir, n[:len(n)-11])
			(*r)[f] = strings.Split(string(c), "\n")
		}
	}
	return nil
}

type JSON struct {
	Values []string `json:"values"`
}

func parseJSON(file string, res map[string][]string) ([]string, error) {
	contents, err := os.ReadFile(file)
	if err == nil {
		return nil, err
	}

	fileStruct := JSON{}
	err = json.Unmarshal(contents, &fileStruct)
	if err != nil {
		return nil, err
	}

	for _, v := range fileStruct.Values {
		if _, ok := res[v]; !ok {
			return nil, fmt.Errorf("error parsing tick.json: %s is not a valid function", v)
		}
	}
	return fileStruct.Values, nil
}

func LoadFunctions() error {
	dir := shared.FunctionsDir
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
	err := loadFuncDir(p, "", &res)
	if err != nil {
		return err
	}

	invalidCommands := ""
	for fPath, contents := range res {
		for i, cmd := range contents {
			command, _ := utils.FindCommand(cmd)
			if command == nil {
				invalidCommands += fmt.Sprintf("  %s:%v\n", fPath, i)
			}
		}
	}
	if invalidCommands != "" {
		return fmt.Errorf("invalid commands found:\n%s", invalidCommands)
	}

	tickContents, err := parseJSON(path.Join(dir, "tick.json"), res)
	if err == nil {
		res["tick.json"] = tickContents
	}

	loadContents, err := parseJSON(path.Join(dir, "load.json"), res)
	if err == nil {
		res["load.json"] = loadContents
	}

	shared.Functions = res
	return nil
}

type Function struct {
	Function string `cmd:"function"`
}

func (f Function) Run(source cmd.Source, output *cmd.Output) {
	cmds, ok := shared.Functions[f.Function]
	if !ok {
		output.Errorf("Unable to find function %s", f.Function)
		return
	}
	for _, cmd := range cmds {
		if strings.TrimSpace(cmd) == "" {
			continue
		}

		command, commandName := utils.FindCommand(cmd)
		args := strings.TrimPrefix(strings.TrimPrefix(cmd, commandName), " ")
		command.Execute(args, source)
	}
}

func (f Function) Allow(source cmd.Source) bool {
	return utils.CommandPermission(source, "minecraft.chat.command.function")
}
