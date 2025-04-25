package defaults

import "embed"

//go:embed layouts
var LayoutsDir embed.FS

//go:embed static
var StaticDir embed.FS
