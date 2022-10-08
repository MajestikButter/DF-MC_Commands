package commands

import (
	"strings"

	"github.com/MajestikButter/DF-MC_Commands/commands/cmdtypes"
	"github.com/MajestikButter/DF-MC_Commands/commands/utils"
	"github.com/go-gl/mathgl/mgl64"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/npc"
)

// TODO: Make command source for npc to allow for bypassing required permissions

type NPCAction struct {
	Command    string
	IgnorePerm bool
}
type NPCStorage struct {
	Texture npc.Texture
	Model   npc.Model
	Entity  *player.Player
	Action  *NPCAction
}

var npcs = map[string]*NPCStorage{}

type NPCId string

func (NPCId) Type() string {
	return "NPCId"
}

// Options returns a number for every plot the player has.
func (NPCId) Options(source cmd.Source) []string {
	keys := []string{}
	for k := range npcs {
		keys = append(keys, k)
	}

	return keys
}

type NPCCreate struct {
	Create   cmd.SubCommand       `cmd:"create"`
	Id       string               `cmd:"NPCId"`
	Name     string               `cmd:"name"`
	SpawnPos mgl64.Vec3           `cmd:"spawnPos"`
	Texture  string               `cmd:"texturePath"`
	Geometry cmd.Optional[string] `cmd:"geometryPath"`
}

func (t NPCCreate) Run(source cmd.Source, output *cmd.Output) {
	id := t.Id
	_, exists := npcs[id]
	if exists {
		output.Errorf("An NPC already exists with the identifier %s", id)
		return
	}

	texture, tErr := npc.ParseTexture(t.Texture)
	if tErr != nil {
		output.Error("Failed to parse texture")
		return
	}

	geometry := npc.DefaultModel
	geoPath, geoUsed := t.Geometry.Load()
	if geoUsed {
		var gErr error
		geometry, gErr = npc.ParseModel(geoPath)
		if gErr != nil {
			output.Error("Failed to parse geometry")
			return
		}
	}

	skin, sErr := npc.Skin(texture, geometry)
	if sErr != nil {
		output.Error("Failed to create npc skin")
		return
	}

	npcs[id] = &NPCStorage{
		texture,
		geometry,
		npc.Create(
			npc.Settings{
				Name:       t.Name,
				Scale:      1,
				Skin:       skin,
				Position:   t.SpawnPos,
				Immobile:   true,
				Vulnerable: false,
			},
			source.World(),
			func(p *player.Player) {
				s := npcs[id]
				a := s.Action
				cmdStr := a.Command
				if !strings.HasPrefix(cmdStr, "/") {
					cmdStr = "/" + cmdStr
				}

				ExecuteCommand(s.Entity, cmdStr, p, &cmd.Output{}, a.IgnorePerm)
			},
		),
		&NPCAction{
			"give @s stick",
			true,
		},
	}
}

func (t NPCCreate) Allow(source cmd.Source) bool {
	return utils.CommandPermission(source, "minecraft.chat.command.npc.create")
}

type NPCDelete struct {
	Delete cmd.SubCommand `cmd:"delete"`
	Id     string         `cmd:"NPCId"`
}

func (t NPCDelete) Run(source cmd.Source, output *cmd.Output) {
	id := string(t.Id)
	sNPC, exists := npcs[id]
	if !exists {
		output.Errorf("No NPC with the identifier %s exists", id)
		return
	}
	err := sNPC.Entity.Close()
	if err != nil {
		output.Errorf("Failed to remove NPC with the identifier %s", id)
		return
	}
	delete(npcs, id)
	output.Printf("Removed NPC with the identifier %s successfully", id)
}

func (t NPCDelete) Allow(source cmd.Source) bool {
	return utils.CommandPermission(source, "minecraft.chat.command.npc.delete")
}

type NPCEditTexture struct {
	Edit    cmd.SubCommand `cmd:"edit"`
	Id      NPCId          `cmd:"NPCId"`
	Texture cmd.SubCommand `cmd:"texture"`
	Path    string         `cmd:"texturePath"`
}

func (t NPCEditTexture) Run(source cmd.Source, output *cmd.Output) {
	id := string(t.Id)
	sNPC, exists := npcs[id]
	if !exists {
		output.Errorf("No NPC with the identifier %s exists", id)
		return
	}

	texture, tErr := npc.ParseTexture(t.Path)
	if tErr != nil {
		output.Error("Failed to parse texture")
		return
	}

	skin, sErr := npc.Skin(texture, sNPC.Model)
	if sErr != nil {
		output.Error("Failed to create npc skin")
		return
	}

	sNPC.Entity.SetSkin(skin)
	sNPC.Texture = texture
}

func (t NPCEditTexture) Allow(source cmd.Source) bool {
	return utils.CommandPermission(source, "minecraft.chat.command.npc.edit.texture")
}

type NPCEditGeometry struct {
	Edit     cmd.SubCommand `cmd:"edit"`
	Id       NPCId          `cmd:"NPCId"`
	Geometry cmd.SubCommand `cmd:"geometry"`
	Path     string         `cmd:"geometryPath"`
}

func (t NPCEditGeometry) Run(source cmd.Source, output *cmd.Output) {
	id := string(t.Id)
	sNPC, exists := npcs[id]
	if !exists {
		output.Errorf("No NPC with the identifier %s exists", id)
		return
	}

	geometry, tErr := npc.ParseModel(t.Path)
	if tErr != nil {
		output.Error("Failed to parse geometry")
		return
	}

	skin, sErr := npc.Skin(sNPC.Texture, geometry)
	if sErr != nil {
		output.Error("Failed to create npc skin")
		return
	}

	sNPC.Entity.SetSkin(skin)
	sNPC.Model = geometry
}

func (t NPCEditGeometry) Allow(source cmd.Source) bool {
	return utils.CommandPermission(source, "minecraft.chat.command.npc.edit.geometry")
}

type NPCEditName struct {
	Edit    cmd.SubCommand `cmd:"edit"`
	Id      NPCId          `cmd:"NPCId"`
	Name    cmd.SubCommand `cmd:"name"`
	NewName string         `cmd:"newName"`
}

func (t NPCEditName) Run(source cmd.Source, output *cmd.Output) {
	id := string(t.Id)
	sNPC, exists := npcs[id]
	if !exists {
		output.Errorf("No NPC with the identifier %s exists", id)
		return
	}

	sNPC.Entity.SetNameTag(t.NewName)
}

func (t NPCEditName) Allow(source cmd.Source) bool {
	return utils.CommandPermission(source, "minecraft.chat.command.npc.edit.name")
}

type NPCEditAction struct {
	Edit       cmd.SubCommand   `cmd:"edit"`
	Id         NPCId            `cmd:"NPCId"`
	Action     cmd.SubCommand   `cmd:"action"`
	IgnorePerm bool             `cmd:"ignorePermission"`
	Command    cmdtypes.Command `cmd:"command"`
}

func (t NPCEditAction) Run(source cmd.Source, output *cmd.Output) {
	id := string(t.Id)
	sNPC, exists := npcs[id]
	if !exists {
		output.Errorf("No NPC with the identifier %s exists", id)
		return
	}

	action := sNPC.Action
	action.Command = string(t.Command)
	action.IgnorePerm = t.IgnorePerm
}

func (t NPCEditAction) Allow(source cmd.Source) bool {
	return utils.CommandPermission(source, "minecraft.chat.command.npc.edit.action")
}

type NPCEditPosition struct {
	Edit     cmd.SubCommand `cmd:"edit"`
	Id       NPCId          `cmd:"NPCId"`
	Pos      cmd.SubCommand `cmd:"position"`
	Position mgl64.Vec3     `cmd:"position"`
}

func (t NPCEditPosition) Run(source cmd.Source, output *cmd.Output) {
	id := string(t.Id)
	sNPC, exists := npcs[id]
	if !exists {
		output.Errorf("No NPC with the identifier %s exists", id)
		return
	}

	sNPC.Entity.Teleport(t.Position)
}

func (t NPCEditPosition) Allow(source cmd.Source) bool {
	return utils.CommandPermission(source, "minecraft.chat.command.npc.edit.position")
}

type NPCEditRotation struct {
	Edit     cmd.SubCommand `cmd:"edit"`
	Id       NPCId          `cmd:"NPCId"`
	Rot      cmd.SubCommand `cmd:"rotation"`
	Rotation mgl64.Vec2     `cmd:"location"`
}

func (t NPCEditRotation) Run(source cmd.Source, output *cmd.Output) {
	id := string(t.Id)
	sNPC, exists := npcs[id]
	if !exists {
		output.Errorf("No NPC with the identifier %s exists", id)
		return
	}

	rot := t.Rotation
	sNPC.Entity.Move(mgl64.Vec3{}, t.Rotation.Y()-rot.Y(), t.Rotation.X()-rot.X())
}

func (t NPCEditRotation) Allow(source cmd.Source) bool {
	return utils.CommandPermission(source, "minecraft.chat.command.npc.edit.position")
}
