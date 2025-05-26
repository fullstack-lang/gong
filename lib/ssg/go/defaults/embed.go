package defaults

import "embed"

//go:embed all:layouts
var LayoutsDir embed.FS

//go:embed layouts/_default
var Defaults embed.FS

//go:embed static
var StaticDir embed.FS
