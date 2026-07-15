package assets

import "embed"

// Files contains every Skill and workspace template required at runtime.
// The all: prefix includes dotfiles such as .claude and .gitkeep.
//
//go:embed all:assets
var Files embed.FS
