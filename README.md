# **DF-MC_Commands**

## Installation

```
go get github.com/MajestikButter/DF-MC_Commands@latest
```

## Usage

### Setting up

```go
import (
  "github.com/MajestikButter/DF-MC_Commands/commands"
)

func main() {
  // Loads commands and sets up Server and PermissionSystem within the module
  commands.Load(/* server.Server */, /* permissions.PermissionSystem */, /* Commands to not load: */[]string{"npc"})
}
```
