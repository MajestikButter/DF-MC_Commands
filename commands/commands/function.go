package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/MajestikButter/DF-MC_Commands/commands/cmdtypes"
	"github.com/MajestikButter/DF-MC_Commands/commands/console"
	"github.com/MajestikButter/DF-MC_Commands/commands/shared"
	"github.com/MajestikButter/DF-MC_Commands/commands/utils"
	"github.com/go-gl/mathgl/mgl64"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
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

func parseJSON(file string, res map[string][]string) ([]string, bool, error) {
	contents, err := os.ReadFile(file)
	if err != nil {
		return nil, false, err
	}

	fileStruct := JSON{}
	err = json.Unmarshal(contents, &fileStruct)
	if err != nil {
		return nil, true, err
	}

	vals := fileStruct.Values
	for i, v := range vals {
		if _, ok := res[v]; !ok {
			return nil, true, fmt.Errorf("error parsing tick.json: %s is not a valid function", v)
		}
		vals[i] = "function " + v
	}
	return vals, true, nil
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

	tickContents, exists, err := parseJSON(path.Join(p, "tick.json"), res)
	if err == nil {
		res["tick.json"] = tickContents
	} else if exists {
		return err
	}

	loadContents, exists, err := parseJSON(path.Join(p, "load.json"), res)
	if err == nil {
		res["load.json"] = loadContents
	} else if exists {
		return err
	}

	if _, ok := res["load.json"]; ok {
		console.ExecuteCommand("function load.json", false)
	}

	shared.Functions = res
	return nil
}

type FunctionSource struct {
	source cmd.Source
}

func (c *FunctionSource) SendCommandOutput(output *cmd.Output) {
}

func (c *FunctionSource) Name() string {
	return c.source.Name()
}

func (c *FunctionSource) Position() mgl64.Vec3 {
	return c.source.Position()
}

func (c *FunctionSource) World() *world.World {
	return c.source.World()
}

type Function struct {
	Function cmdtypes.Function `cmd:"function"`
}

func (f Function) Run(source cmd.Source, output *cmd.Output) {
	cmds, ok := shared.Functions[string(f.Function)]
	if !ok {
		output.Errorf("Unable to find function %s", f.Function)
		return
	}
	count := 0
	for _, cmd := range cmds {
		if strings.TrimSpace(cmd) == "" {
			continue
		}

		command, commandName := utils.FindCommand(cmd)
		args := strings.TrimPrefix(strings.TrimPrefix(cmd, commandName), " ")
		command.Execute(args, &FunctionSource{source})

		count++
	}
	output.Printf("Executed %v commands from %s", count, f.Function)
}

func (f Function) Allow(source cmd.Source) bool {
	return utils.CommandPermission(source, "minecraft.chat.command.function")
}
