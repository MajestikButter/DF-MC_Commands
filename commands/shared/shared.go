package shared

import (
	"github.com/MajestikButter/DF-MC_Permissions/permissions"
	"github.com/df-mc/dragonfly/server"
)

var PermSystem *permissions.PermissionSystem
var Server *server.Server
var Functions map[string][]string
